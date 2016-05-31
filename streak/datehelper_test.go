package streak

import (
	"testing"
	"time"

	"github.com/simulatedsimian/assert"
)

func TestFromYearDay(t *testing.T) {
	assert := assert.Make(t)

	tm := time.Date(2001, 10, 10, 0, 0, 0, 0, time.Local)
	yd := tm.YearDay()
	assert(tm).Equal(fromYearDay(2001, yd))

	tm = time.Date(2001, 12, 31, 0, 0, 0, 0, time.Local)
	yd = tm.YearDay()
	assert(tm).Equal(fromYearDay(2001, yd))

	assert(isLeap(2016)).Equal(true)
}
