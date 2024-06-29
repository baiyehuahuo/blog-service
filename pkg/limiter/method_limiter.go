package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

var _ LimiterIface = (*MethodLimiter)(nil)

func NewMethodLimiter() *MethodLimiter {
	return &MethodLimiter{
		Limiter: &Limiter{
			limiterBucket: make(map[string]*ratelimit.Bucket),
		},
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBucket[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...BucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBucket[rule.Key]; ok {
			l.limiterBucket[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}
