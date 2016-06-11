package streak

import (
	"fmt"
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
	assert(isLeap(2015)).Equal(false)
}

func TestStreak(t *testing.T) {
	assert := assert.Make(t)

	s := Streak{fromYearDay(2001, 1), fromYearDay(2001, 1), 0}
	assert(s.Length()).Equal(1)
	s = Streak{fromYearDay(2001, 1), fromYearDay(2001, 2), 0}
	assert(s.Length()).Equal(2)
	s = Streak{fromYearDay(2001, 1), fromYearDay(2002, 1), 0}
	assert(s.Length()).Equal(366)
}

func TestStreakRec(t *testing.T) {
	assert := assert.Make(t)

	sr := StreakRecorder{}

	assert(len(sr.GetStreaks())).Equal(0)
	sr.AddCommit(fromYearDay(2016, 1))
	assert(len(sr.GetStreaks())).Equal(0)
	sr.AddCommit(fromYearDay(2016, 2))
	assert(len(sr.GetStreaks())).Equal(1)
	assert(sr.GetStreaks()[0].Length()).Equal(2)
	sr.AddCommit(fromYearDay(2016, 3))
	assert(len(sr.GetStreaks())).Equal(1)
	assert(sr.GetStreaks()[0].Length()).Equal(3)
	sr.AddCommit(fromYearDay(2015, 365))
	assert(len(sr.GetStreaks())).Equal(1)
	assert(sr.GetStreaks()[0].Length()).Equal(4)

	fmt.Println(sr.GetStreaks()[0])
}
