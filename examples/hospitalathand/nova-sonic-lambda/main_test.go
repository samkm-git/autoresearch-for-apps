package main

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type MockBedrockClient struct {
	InvokeModelFunc func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error)
}

func (m *MockBedrockClient) InvokeModel(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
	return m.InvokeModelFunc(ctx, params, optFns...)
}

func TestHandleRequest_FallbackPrompt(t *testing.T) {
	bedrockClient = &MockBedrockClient{
		InvokeModelFunc: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			
			// Verify payload contains fallback prompt
			var payload map[string]interface{}
			json.Unmarshal(params.Body, &payload)
			
			prompt, ok := payload["prompt"].(string)
			if !ok || prompt != "Hello, how can I help you today?" {
				t.Errorf("Expected fallback prompt, got: %v", prompt)
			}

			outputJSON := `{"completion": "I am here to help."}`
			return &bedrockruntime.InvokeModelOutput{
				Body: []byte(outputJSON),
			}, nil
		},
	}

	event := ConnectEvent{} // Empty attributes
	res, err := HandleRequest(context.Background(), event)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if res["status"] != "Success" || res["nova_msg"] != "I am here to help." {
		t.Errorf("Unexpected response: %v", res)
	}
}

func TestHandleRequest_CustomPrompt(t *testing.T) {
	bedrockClient = &MockBedrockClient{
		InvokeModelFunc: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			
			var payload map[string]interface{}
			json.Unmarshal(params.Body, &payload)
			
			prompt, _ := payload["prompt"].(string)
			if prompt != "I have a headache" {
				t.Errorf("Expected custom prompt, got: %v", prompt)
			}

			outputJSON := `{"completion": "Please take some rest."}`
			return &bedrockruntime.InvokeModelOutput{
				Body: []byte(outputJSON),
			}, nil
		},
	}

	event := ConnectEvent{}
	event.Details.ContactData.Attributes = map[string]string{
		"PatientPrompt": "I have a headache",
	}
	res, err := HandleRequest(context.Background(), event)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if res["nova_msg"] != "Please take some rest." {
		t.Errorf("Unexpected completion: %v", res["nova_msg"])
	}
}

func TestHandleRequest_BedrockError(t *testing.T) {
	bedrockClient = &MockBedrockClient{
		InvokeModelFunc: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			return nil, errors.New("simulated bedrock failure")
		},
	}

	event := ConnectEvent{}
	res, err := HandleRequest(context.Background(), event)
	if err != nil {
		t.Fatalf("Expected no error returned to Lambda, got: %v", err)
	}

	if res["status"] != "VoiceProcessingError" {
		t.Errorf("Expected VoiceProcessingError status, got: %v", res["status"])
	}
}
