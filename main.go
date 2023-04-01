package main

import "fmt"

const (
	AccessKeyId     = "XXXXXXXXXXXXXXXX"
	SecretAccessKey = "1111111111111111222222223333334444444444"
	Region          = "us-east-1"
	RoleARN         = "arn:aws:iam::012345678901:role/my-role-to-assume"
)

func main() {

	// sdk-v1
	sessv1, err := AssumedRoleV1(RoleARN, Region)
	if err != nil {
		fmt.Println("Error getting session :", err)
		return
	}
	creds, err := sessv1.Config.Credentials.Get()
	if err != nil {
		fmt.Println("Error Getting Credentials:", err)
	}
	fmt.Println(creds.SessionToken)

	// sdk-v2
	sts := AssumeRoleV2(AccessKeyId, SecretAccessKey, RoleARN, Region)
	fmt.Println("v2 session token: ", sts.SessionToken)

}
