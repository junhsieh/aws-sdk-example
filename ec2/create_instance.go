/*
   Copyright 2010-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/
/*
 * Reference: https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/ec2
 */
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func main() {
	region := "us-west-2"
	zone := region + "a"

	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
	// Specify the details of the instance that you want to create.
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An CentOS 7 AMI ID for t2.nano instances in the us-west-2 region
		ImageId:      aws.String("ami-d2c924b2"),
		InstanceType: aws.String("t2.nano"),
		KeyName:      aws.String("mykey"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		Placement: &ec2.Placement{
			AvailabilityZone: aws.String(zone),
		},
	})

	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFirstInstance"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	log.Println("Successfully tagged instance")
}
