package distributed_locker

import (
	"context"
	"time"
)

// Locker 分布式锁实现接口
type Locker interface {
	// Lock 加锁
	// param
	// ex: 锁过期时间
	// return
	// isSet: 是否成功加锁
	Lock(ctx context.Context, ex time.Duration) (isSet bool, err error)

	// UnLock 释放锁
	UnLock(ctx context.Context) error
}
