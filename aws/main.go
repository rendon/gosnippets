package main

// Import the AWS SDK for Go
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

/**
 * Don't hard-code your credentials!
 * Export the following environment variables instead:
 *
 * export AWS_ACCESS_KEY_ID='AKID'
 * export AWS_SECRET_ACCESS_KEY='SECRET'
 */

func main() {
	// Set your region for future requests.
	defaults.DefaultConfig.Region = aws.String("us-west-2")

	svc := ec2.New(nil)

	var minCount int64 = 1
	var maxCount int64 = 1
	params := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-989b7bab"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     &minCount,
		MaxCount:     &maxCount,
	}

	runResult, err := svc.RunInstances(params)
	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *runResult.Instances[0].InstanceId)

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
		log.Println("Could not create tags for instance", *runResult.Instances[0].InstanceId, err)
	}

	log.Println("Successfully tagged instance")
}
