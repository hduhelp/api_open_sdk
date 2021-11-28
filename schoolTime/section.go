package schoolTime

import (
	"github.com/hduhelp/api_open_sdk/types"
	"time"
)

type SectionReader interface {
	StartTime(date *types.Date, section int32) time.Time
	EndTime(date *types.Date, section int32) time.Time
}
