package activities

import (
	"context"
	"fmt"
)

func Greet(ctx context.Context, user string) (string, error) {

	fmt.Print("sendGreeting activity called")
	res := fmt.Sprintf("hello %s", user)
	return res, nil
}
