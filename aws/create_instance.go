/**
 * Create EC2 instance from existing AMI.

 * Don't hard-code your credentials!
 * Export the following environment variables instead:
 *
 * export AWS_ACCESS_KEY_ID='AKID'
 * export AWS_SECRET_ACCESS_KEY='SECRET'
 *
 * This example loads credentials from ~/.aws/credentials:
 * [default]
 * aws_access_key_id = ...
 * aws_secret_access_key = ...
 */
package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Set your region for future requests.
	defaults.DefaultConfig.Region = aws.String("us-west-2")

	svc := ec2.New(nil)
	var image = "ami-xxxxxxxx" // replace with your image ID
	params := &ec2.RunInstancesInput{
		ImageId:      aws.String(image),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	runResult, err := svc.RunInstances(params)
	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *runResult.Instances[0].InstanceId)

	time.Sleep(10 * time.Second)
	// Add tags to the instance
	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("MyInstanceName"),
			},
		},
	})
	if err != nil {
		var id = *runResult.Instances[0].InstanceId
		log.Println("Could not create tags for instance %s:  %s", id, err)
	}

	log.Println("Successfully tagged instance")
}
