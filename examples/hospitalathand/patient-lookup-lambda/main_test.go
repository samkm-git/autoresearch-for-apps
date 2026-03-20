package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MockDynamo struct {
	GetItemFunc func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func (m *MockDynamo) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return m.GetItemFunc(ctx, params, optFns...)
}

func TestHandleRequest_Found(t *testing.T) {
	dbClient = &MockDynamo{
		GetItemFunc: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{
				Item: map[string]types.AttributeValue{
					"Name": &types.AttributeValueMemberS{Value: "John Doe"},
				},
			}, nil
		},
	}

	res, _ := HandleRequest(context.Background(), LookupEvent{PatientID: "123"})
	if res["status"] != "Found" || res["name"] != "John Doe" {
		t.Errorf("Expected John Doe, got %v", res)
	}
}

func TestHandleRequest_NotFound(t *testing.T) {
	dbClient = &MockDynamo{
		GetItemFunc: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{Item: nil}, nil
		},
	}

	res, _ := HandleRequest(context.Background(), LookupEvent{PatientID: "999"})
	if res["status"] != "NotFound" {
		t.Errorf("Expected NotFound, got %v", res["status"])
	}
}
