package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	r := os.Getenv("AWS_REGION")
	if r == "" {
		r = "us-east-1"
	}

	var role = flag.String("role", "", "The role ARN we will be assuming")
	var sess = flag.String("session", "assumerole", "Session name to include in role assumption [optional]")
	var externalid = flag.String("externalid", "", "External ID specified to assume role [optional]")

	flag.Parse()

	var awsexternalid = aws.String(*externalid)

	if *externalid == "" {
		awsexternalid = nil
	}

	svc := sts.New(session.New(), &aws.Config{Region: aws.String(r)})

	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(*role), // Required
		RoleSessionName: aws.String(*sess), // Required
		ExternalId:      awsexternalid,
	}

	// Call the DescribeInstances Operation
	resp, err := svc.AssumeRole(params)
	if err != nil {
		panic(err)
	}

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Printf("export AWS_ACCESS_KEY=\"%s\"\n", *resp.Credentials.AccessKeyId)
	fmt.Printf("export AWS_SECRET_KEY=\"%s\"\n", *resp.Credentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=\"%s\"\n", *resp.Credentials.SessionToken)
}
