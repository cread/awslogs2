package awslogs2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func Streams(ses *session.Session, group string, streamPrefix string) {
	logs := cloudwatchlogs.New(ses)

	input := cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(group),
	}

	if streamPrefix != "" {
		input.LogStreamNamePrefix = aws.String(streamPrefix)
	}

	logs.DescribeLogStreamsPages(
		&input,
		func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			for _, stream := range page.LogStreams {
				fmt.Println(aws.StringValue(stream.LogStreamName))
			}
			return true
		})

}