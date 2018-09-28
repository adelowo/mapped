package utils

import "time"

type (
	Timestamp int64
)

func Now() Timestamp {
	return Timestamp(time.Now().Unix())
}

func (t Timestamp) Since(other Timestamp) time.Duration {
	return t.AsTime().Sub(other.AsTime())
}

func (t Timestamp) AsTime() time.Time {
	return time.Unix(int64(t), 0)
}
