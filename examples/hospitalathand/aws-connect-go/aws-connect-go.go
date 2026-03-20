package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect"
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
	// Read the generated Connect Flow
	flowFile, err := ioutil.ReadFile("connect_flow.json")
	if err != nil {
		fmt.Printf("Could not read connect_flow.json: %v\n", err)
	}
	flowJSON := string(flowFile)

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

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
	 Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  jsii.String("ap-southeast-1"),
	}
}
