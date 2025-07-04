package common

import "time"

type Timestamped interface {
	Timestamp() time.Time
}
