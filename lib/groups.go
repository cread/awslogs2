package awslogs2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func Groups(ses *session.Session, groupPrefix string) {
	logs := cloudwatchlogs.New(ses)

	input := cloudwatchlogs.DescribeLogGroupsInput{}

	if groupPrefix != "" {
		input.LogGroupNamePrefix = aws.String(groupPrefix)
	}

	logs.DescribeLogGroupsPages(
		&input,
		func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
			for _, group := range page.LogGroups {
				fmt.Println(aws.StringValue(group.LogGroupName))
			}
			return true
		})
}