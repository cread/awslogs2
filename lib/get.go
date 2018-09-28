package awslogs2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func Get(ses *session.Session, group string, streamPrefix string) {
	logs := cloudwatchlogs.New(ses)

	input := cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: aws.String(group),
	}

	if streamPrefix != "" {
		input.LogStreamNamePrefix = aws.String(streamPrefix)
	}

	logs.FilterLogEventsPages(
		&input,
		func(page *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) bool {
			for _, event := range page.Events {
				fmt.Print(aws.StringValue(event.Message))
			}
			return true
		})
}