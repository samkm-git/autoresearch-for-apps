package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// ConnectEvent represents the payload from Amazon Connect
type ConnectEvent struct {
	Details struct {
		ContactData struct {
			Attributes map[string]string `json:"Attributes"`
		} `json:"ContactData"`
	} `json:"Details"`
}

// Logic Agent Implementation: Nova Sonic 2 Processor
func HandleRequest(ctx context.Context, event ConnectEvent) (map[string]interface{}, error) {
	fmt.Printf("Processing call for contact attributes: %v\n", event.Details.ContactData.Attributes)

	// 1. Initialize Bedrock Client
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	client := bedrockruntime.NewFromConfig(cfg)

	// 2. Prepare Nova Sonic 2 Request
	// Note: We use the patient prompt (e.g. from attributes or speech-to-text)
	userPrompt := event.Details.ContactData.Attributes["PatientPrompt"]
	if userPrompt == "" {
		userPrompt = "Hello, how can I help you today?"
	}

	// This is a placeholder for the actual Nova Sonic JSON structure
	payload := map[string]interface{}{
		"prompt": userPrompt,
		"max_tokens": 200,
		"temperature": 0.5,
	}
	payloadBytes, _ := json.Marshal(payload)

	// 3. Invoke Amazon Nova Sonic 2
	modelID := "amazon.nova-sonic-v1:0" 
	output, err := client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     &modelID,
		ContentType: "application/json",
		Body:        payloadBytes,
	})

	if err != nil {
		log.Printf("Nova Sonic invocation failed: %v", err)
		return map[string]interface{}{
			"status": "VoiceProcessingError",
			"response": "I'm sorry, I'm having trouble understanding. Please say that again.",
		}, nil
	}

	// 4. Parse Response
	var responseData map[string]interface{}
	json.Unmarshal(output.Body, &responseData)
	
	// Simplify for Amazon Connect return attributes
	return map[string]interface{}{
		"status":   "Success",
		"nova_msg": responseData["completion"],
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
