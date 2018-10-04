package days

import (
	"time"
)

// Tomorrow gives you tomorrows date
func Tomorrow() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d+1, 0, 0, 0, 0, time.Now().Location())
}
