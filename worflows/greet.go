package worflows

import (
	"temporal/activities"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Greetings(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{

		ScheduleToStartTimeout: time.Hour,
		StartToCloseTimeout:    time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, activities.Greet, name).Get(ctx, &result)
	return result, err
}
