package locker

import "time"

var (
	defaultAcquireTimeOut = 2 * time.Second
	defaultLockTimeOut    = 5 * time.Second
)

type LockOption struct {
	acquireTimeOut time.Duration
	lockTimeOut    time.Duration
}

type LockOptions []LockOption

func (l LockOptions) acquireTimeOut() time.Duration {
	for _, v := range l {
		if v.acquireTimeOut != 0 {
			return v.acquireTimeOut
		}
	}
	return defaultAcquireTimeOut
}

func (l LockOptions) lockTimeOut() time.Duration {
	for _, v := range l {
		if v.lockTimeOut != 0 {
			return v.lockTimeOut
		}
	}
	return defaultLockTimeOut
}
