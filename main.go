package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cread/awslogs2/lib"
	"os"
)

func listBuckets() {
	// Load ~/.aws/config
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	fmt.Printf("Region: %s\n", *ses.Config.Region)

	svc := s3.New(ses)
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name),
			aws.TimeValue(b.CreationDate),
		)
	}
}

func Usage(cmd string) {
	fmt.Printf("usage: %s [options] <command> <subcommand> [<subcommand> ...] [parameters]\n", cmd)
	os.Exit(2)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		Usage(args[0])
	}

	// Load ~/.aws/config
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	prefix := ""

	switch args[1] {
	case "help":
		// TODO: Wire up some help...
	case "groups":
		// groups [prefix]
		if len(args) == 3 {
			prefix = args[2]
		}
		awslogs2.Groups(ses, prefix)
	case "streams":
		// streams <group> [prefix]
		if len(args) == 4 {
			prefix = args[3]
		}
		awslogs2.Streams(ses, args[2], prefix)
	case "get":
		// get <group> [prefix]
		if len(args) == 4 {
			prefix = args[3]
		}
		awslogs2.Get(ses, args[2], prefix)
	}
}
