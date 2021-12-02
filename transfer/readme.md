# Tips for use transfer

## Example

```go
package main

import (
	"fmt"
	"testing"
)

type Time struct {
	SchoolYear string `json:"schoolYear"`
	Semester   string `json:"semester"`
	WeekNow    int    `json:"weekNow"`
	WeekDayNow int    `json:"weekDayNow"`
	TimeStamp  int    `json:"timeStamp"`
	Section    int    `json:"section"`
}

func main() {
	transfer.Init("appId", "appSecret",
		Endpoint("http://localhost:8080"),
	)

	data := new(Time)
	resp, body, errs := transfer.Get("gateway", "/time", "")
	code, status, err := req.EndStruct(&data)
	fmt.Println(code, status, err)
}

```