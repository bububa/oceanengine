package core

import (
	"time"
)

type RateLimiter interface {
	Take() time.Duration
}
