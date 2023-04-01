package main

import (
	"context"
	"log"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	credentials "github.com/aws/aws-sdk-go-v2/credentials"
	stscredsv2 "github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	stsv2 "github.com/aws/aws-sdk-go-v2/service/sts"
)

func AssumeRoleV2(ak string, sak string, role string, region string) awsv2.Credentials {

	ctx := context.Background()
	assumecnf, _ := config.LoadDefaultConfig(
		ctx, config.WithRegion(region),
		config.WithCredentialsProvider(awsv2.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				ak,
				sak, "",
			)),
		),
	)

	stsclient := stsv2.NewFromConfig(assumecnf)

	cnf, _ := config.LoadDefaultConfig(
		ctx, config.WithRegion(region),
		config.WithCredentialsProvider(awsv2.NewCredentialsCache(
			stscredsv2.NewAssumeRoleProvider(
				stsclient,
				string(role),
				func(o *stscredsv2.AssumeRoleOptions) {
					o.RoleARN = role
				},
			)),
		),
	)
	sts_creds, err := cnf.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return sts_creds
}
