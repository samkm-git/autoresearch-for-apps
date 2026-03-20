package main

import (
	"testing"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestAwsConnectGoStack(t *testing.T) {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	stack := NewAwsConnectGoStack(app, "TestStack", &AwsConnectGoStackProps{
		awscdk.StackProps{Env: env()},
	})

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"AttributeDefinitions": []interface{}{
			map[string]interface{}{
				"AttributeName": "PatientID",
				"AttributeType": "S",
			},
		},
	})
}
