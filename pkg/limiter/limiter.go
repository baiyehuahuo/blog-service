package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// LimiterIface 限流器有多种实现方式，抽象出来以便接口的实现
type LimiterIface interface {
	Key(*gin.Context) string                        // 对应限流器键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool) // 获取令牌桶
	AddBuckets(...BucketRule) LimiterIface          // 新增多个令牌桶
}

type Limiter struct {
	limiterBucket map[string]*ratelimit.Bucket
}

type BucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多久放 N 个令牌
	Capacity     int64         // 令牌桶容量
	Quantum      int64         // 每次到达间隔时间后放的具体令牌数量
}
