package schooltimev1

import (
	"time"

	"github.com/hduhelp/api_open_sdk/types"
)

type SectionReader interface {
	StartTime(date *types.Date, section int32) time.Time
	EndTime(date *types.Date, section int32) time.Time
}
