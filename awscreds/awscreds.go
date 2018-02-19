package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	profile := flag.String("profile", "default", "AWS profile to be used to parse key/secret.")
	field := flag.String("field", "key", "Credentials field {key|secret|both} to extract.")
	flag.Parse()

	credentialsFile := os.Getenv("HOME") + "/.aws/credentials"
	fieldValue := ParseCredentials(credentialsFile, *profile, *field)
	fmt.Println(fieldValue)
}

func ParseCredentials(file, profile, field string) (fieldValue string) {
	creds := credentials.NewSharedCredentials(file, profile)
	// Retrieve the credentials value
	credValue, err := creds.Get()
	if err != nil {
		panic(err)
	}

	if field == "key" {
		return credValue.AccessKeyID
	} else if field == "secret" {
		return credValue.SecretAccessKey
	} else if field == "both" {
		return credValue.AccessKeyID + "\n" + credValue.SecretAccessKey
	}
	return ""
}
