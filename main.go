package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

func getTag(tags []types.Tag, key string) string {
	for _, tag := range tags {
		if *tag.Key == key {
			return *tag.Value
		}
	}
	return ""
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	result, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance id: %s, name: %s, state: %s\n", *instance.InstanceId, getTag(instance.Tags,
				"Name"), instance.State.Name)
		}
	}

}
