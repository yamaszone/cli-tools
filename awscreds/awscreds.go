package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	profile := flag.String("profile", "default", "AWS profile to be used to parse key/secret.")
	field := flag.String("field", "", "Credentials field {key|secret} to extract. Prints key followed by secret if none specified.")
	flag.Parse()

	creds := credentials.NewSharedCredentials(os.Getenv("HOME")+"/.aws/credentials", *profile)
	// Retrieve the credentials value
	credValue, err := creds.Get()
	if err != nil {
		panic(err)
	}

	if *field == "key" {
		fmt.Println(credValue.AccessKeyID)
	} else if *field == "secret" {
		fmt.Println(credValue.SecretAccessKey)
	} else {
		fmt.Println(credValue.AccessKeyID)
		fmt.Println(credValue.SecretAccessKey)
	}
}
