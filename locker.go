package distributed_locker

import (
	"context"
	"time"
)

type Locker interface {
	Lock(ctx context.Context, ex time.Duration) (bool, error)
	UnLock(ctx context.Context) error
}
