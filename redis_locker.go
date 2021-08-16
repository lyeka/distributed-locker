package distributed_locker

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const EquivalentDeletionLuaScript string = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

type RedisLocker struct {
	key         string
	unlockValue string
	conn        *redis.Conn
}

func NewRedisLocker(key string, conn *redis.Conn) RedisLocker {
	return RedisLocker{
		key:         key,
		unlockValue: RandStr(10),
		conn:        conn,
	}
}

func (r RedisLocker) Lock(ctx context.Context, ex time.Duration) (isSet bool, err error) {
	res := r.conn.SetNX(ctx, r.key, r.unlockValue, ex)
	return res.Val(), res.Err()
}

func (r RedisLocker) UnLock(ctx context.Context) error {
	return r.conn.Eval(ctx, EquivalentDeletionLuaScript, []string{r.key}, []string{r.unlockValue}).Err()
}
