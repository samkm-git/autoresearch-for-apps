package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AwsConnectGoStackProps struct {
	awscdk.StackProps
}

func NewAwsConnectGoStack(scope constructs.Construct, id string, props *AwsConnectGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create Amazon DynamoDB Table (Replacing Connect due to AISPL restriction)
	awsdynamodb.NewTable(stack, jsii.String("PatientRecords"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("PatientID"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("HospitalAtHand-Patients"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAwsConnectGoStack(app, "AwsConnectGoStack", &AwsConnectGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	account := os.Getenv("CDK_DEFAULT_ACCOUNT")
	if account == "" {
		account = "123456789012"
	}
	return &awscdk.Environment{
	 Account: jsii.String(account),
	 Region:  jsii.String("ap-southeast-1"),
	}
}
