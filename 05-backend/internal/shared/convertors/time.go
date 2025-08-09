package convertors

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampToTime converts google's proto timestamp to native time.Time
func TimestampToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime().Truncate(time.Microsecond)
}

// TimeToTimestamp converts the native time.Time to google's proto timestamp
func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t.Truncate(time.Microsecond))
}
