package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect"
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

	// Create Amazon Connect Instance
	instance := awsconnect.NewCfnInstance(stack, jsii.String("HospitalAtHandConnectInstance"), &awsconnect.CfnInstanceProps{
		Attributes: &awsconnect.CfnInstance_AttributesProperty{
			InboundCalls:  jsii.Bool(true),
			OutboundCalls: jsii.Bool(true),
		},
		IdentityManagementType: jsii.String("CONNECT_MANAGED"),
		InstanceAlias:          jsii.String("hospital-at-hand-ivrs-modern"),
	})

	// Create Amazon Connect Contact Flow
	awsconnect.NewCfnContactFlow(stack, jsii.String("HospitalAtHandLegacyFlow"), &awsconnect.CfnContactFlowProps{
		Content:     jsii.String(flowJSON),
		InstanceArn: instance.AttrArn(),
		Name:        jsii.String("HospitalAtHandLegacyFlowConverted"),
		Type:        jsii.String("CONTACT_FLOW"),
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
