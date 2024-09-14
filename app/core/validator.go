package core

import "context"

type Validator interface {
	Valid(ctx context.Context) (problems map[string]string)
}
