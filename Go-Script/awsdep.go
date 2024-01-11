package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getStartupScript(instanceNumber int) string {
	// Create a Bash script to be executed on the launched EC2 instance

	return fmt.Sprintf(`#!/bin/bash
cd /home/ec2-user/

# Create a file with a greeting message
echo "Hello from EC2 instance %d!" > greeting.txt

# Display the content of the greeting file
cat greeting.txt
`, instanceNumber)
}

func main() {
	// Set up a new AWS session
	allI := []int{0, 19502, 40121, 62075, 85662, 111311, 139681, 171884, 210081, 259862}
	allJ := []int{1, 173755, 364939, 328234, 221109, 220588, 379311, 196624, 363867, 334399}
	inter := 7221501078

	// Create an AWS session
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			"input_key",
			"input_key",
			"", // Optional session token
		),
	}))

	// Create an EC2 client with the AWS session.
	svc := ec2.New(awsSession)

	// Iterate over instances to launch
	for i := 1; i < 10; i++ {
		// Define the input parameters
		startupScript := getStartupScript(i)
		encodedUserData := base64.StdEncoding.EncodeToString([]byte(startupScript))

		input := &ec2.RunInstancesInput{
			ImageId:          aws.String("image_id"),           // Replace with your desired AMI ID
			InstanceType:     aws.String("instance_type"),      // Replace with your desired instance type
			MinCount:         aws.Int64(1),
			MaxCount:         aws.Int64(1),
			KeyName:          aws.String("scrape2"),            // Replace with your key pair name
			SecurityGroupIds: aws.StringSlice([]string{"sg-0e85f46af0201ccaf"}),
			UserData:         aws.String(encodedUserData),
		}

		// Launch the instance
		result, err := svc.RunInstances(input)
		if err != nil {
			fmt.Printf("Error launching instance %d: %s\n", i, err)
			continue
		}

		// Print the results
		fmt.Printf("Successfully launched instance %d: %s\n", i, *result.Instances[0].InstanceId)
	}
}
