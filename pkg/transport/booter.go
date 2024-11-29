package transport

import "context"

type Booter interface {
	Initialize(ctx context.Context) error
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
