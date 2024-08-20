package openapi

import (
	"fmt"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}
