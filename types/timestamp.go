package types

import (
	"encoding/json"
	"time"
)

func (x *Timestamp) MarshalJSON() ([]byte, error) {
	t := time.Unix(x.Seconds, int64(x.Nanos)).UTC()
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Unix())
}

// NewTimestamp constructs a new Timestamp from the provided time.Time.
func NewTimestamp(t time.Time) *Timestamp {
	return &Timestamp{Seconds: int64(t.Unix()), Nanos: int32(t.Nanosecond())}
}

// AsTime converts x to a time.Time.
func (x *Timestamp) AsTime() time.Time {
	return time.Unix(x.GetSeconds(), int64(x.GetNanos()))
}
