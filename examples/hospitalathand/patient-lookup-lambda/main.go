package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type LookupEvent struct {
	PatientID string `json:"PatientID"`
}

type DynamoGetter interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

var dbClient DynamoGetter

func initDB(ctx context.Context) error {
	if dbClient != nil {
		return nil
	}
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	dbClient = dynamodb.NewFromConfig(cfg)
	return nil
}

func HandleRequest(ctx context.Context, event LookupEvent) (map[string]interface{}, error) {
	if err := initDB(ctx); err != nil {
		return nil, err
	}

	result, err := dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("HospitalAtHand-Patients"),
		Key: map[string]types.AttributeValue{
			"PatientID": &types.AttributeValueMemberS{Value: event.PatientID},
		},
	})

	if err != nil {
		log.Printf("DynamoDB error: %v", err)
		return map[string]interface{}{"status": "Error", "message": "Lookup failed"}, nil
	}

	if result.Item == nil {
		return map[string]interface{}{"status": "NotFound", "message": "No patient found"}, nil
	}

	return map[string]interface{}{
		"status": "Found",
		"name":   result.Item["Name"].(*types.AttributeValueMemberS).Value,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
