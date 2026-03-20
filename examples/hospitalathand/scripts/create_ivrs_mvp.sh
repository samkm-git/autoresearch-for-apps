#!/bin/bash

# Amazon Connect MVP Creation Script
# Logic: Automates the transition from legacy JSON to a live Connect Instance

# --- Configuration ---
INSTANCE_ALIAS="hospital-mvp-$(date +%s)"
REGION="ap-southeast-1"  # Stick to Singapore as requested
FLOW_FILE="aws-connect-go/connect_flow.json"

echo "----------------------------------------------------"
echo "🚀 Starting IVRS Modernization MVP via AWS CLI"
echo "Target Region: $REGION"
echo "Instance Alias: $INSTANCE_ALIAS"
echo "----------------------------------------------------"

# 1. Create the Amazon Connect Instance
echo "Step 1: Creating Amazon Connect Instance..."
INSTANCE_RESPONSE=$(aws connect create-instance \
    --identity-management-type CONNECT_MANAGED \
    --instance-alias "$INSTANCE_ALIAS" \
    --inbound-calls-enabled \
    --outbound-calls-enabled \
    --region "$REGION" \
    --query '{Id:Id,Arn:Arn}' \
    --output text)

if [ $? -ne 0 ]; then
    echo "❌ Failed to create instance. This might be due to the AISPL account restriction in $REGION."
    exit 1
fi

# Split the output text (Format: ARN  ID)
read -r INSTANCE_ARN INSTANCE_ID <<< "$INSTANCE_RESPONSE"

echo "✅ Instance requested. ID: $INSTANCE_ID"

# 2. Wait for instance to become active
echo "Step 2: Waiting for instance to become ACTIVE (this can take 2-5 minutes)..."
while true; do
    STATUS=$(aws connect describe-instance --instance-id "$INSTANCE_ID" --region "$REGION" --query 'Instance.InstanceStatus' --output text)
    echo "Current Status: $STATUS..."
    if [ "$STATUS" == "ACTIVE" ]; then
        break
    fi
    sleep 30
done

echo "✅ Instance is ACTIVE."

# 3. Create the Contact Flow
echo "Step 3: Uploading converted Contact Flow from $FLOW_FILE..."
FLOW_CONTENT=$(cat "$FLOW_FILE")

FLOW_ID=$(aws connect create-contact-flow \
    --instance-id "$INSTANCE_ID" \
    --name "LegacyModernizedFlow" \
    --type CONTACT_FLOW \
    --description "Automated migration from legacy PHP logic" \
    --content "$FLOW_CONTENT" \
    --region "$REGION" \
    --query 'ContactFlowId' \
    --output text)

if [ $? -eq 0 ]; then
    echo "✅ Success! Contact Flow Created with ID: $FLOW_ID"
    echo "----------------------------------------------------"
    echo "MVP Summary:"
    echo "Instance: $INSTANCE_ALIAS"
    echo "Flow ID: $FLOW_ID"
    echo "Next: Attach a phone number in the AWS Console to test."
else
    echo "❌ Failed to upload Contact Flow. Please check if the JSON format in $FLOW_FILE is valid for Connect."
fi
