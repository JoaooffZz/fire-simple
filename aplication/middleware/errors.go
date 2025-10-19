package middleware

import "fmt"

type NotFoundEnv struct {
	Field string
}

func (n *NotFoundEnv) Error() string {
	return fmt.Sprintf("ENV: %s not found", n.Field)
}

type InvalidFormaterEnv struct {
	Field string
}

func (i *InvalidFormaterEnv) Error() string {
	return fmt.Sprintf("ENV: %s invalid format, only numbers and letters are allowed", i.Field)
}

type NotAuthClient struct {
	Msg string
}

func (n *NotAuthClient) Error() string {
	return fmt.Sprintf("CLIENT NOT AUTH CONNECTION. Message: %s", n.Msg)
}
