package streak

import (
	"sort"
	"time"
)

const (
	Day time.Duration = 24 * time.Hour
)

func fromYearDay(year, yearDay int) time.Time {
	return time.Date(year, 1, yearDay, 0, 0, 0, 0, time.Local)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

type Streak struct {
	firstDay time.Time
	lastDay  time.Time
	commits  int
}

func (s Streak) Length() int {
	return int(s.lastDay.Sub(s.firstDay)/Day + 1)
}

type StreakRecorder struct {
	commitDays map[int][]int
}

func (sr *StreakRecorder) AddCommit(t time.Time) {
	if sr.commitDays == nil {
		sr.commitDays = make(map[int][]int)
	}
	t = t.Local()

	days := sr.commitDays[t.Year()]
	if days == nil {
		if isLeap(t.Year()) {
			days = make([]int, 366)
		} else {
			days = make([]int, 365)
		}
		sr.commitDays[t.Year()] = days
	}
	days[t.YearDay()-1]++
}

func (sr *StreakRecorder) GetStreaks() (streaks []Streak) {
	if sr.commitDays == nil {
		return
	}

	years := []int{}
	for k := range sr.commitDays {
		years = append(years, k)
	}
	sort.Ints(years)

	var currentStreak *Streak
	for _, year := range years {
		for day, commits := range sr.commitDays[year] {
			if commits > 0 {
				if currentStreak == nil {
					currentStreak = &Streak{fromYearDay(year, day+1), fromYearDay(year, day+1), 0}
				}
				currentStreak.lastDay = fromYearDay(year, day+1)
				currentStreak.commits += commits
			} else {
				if currentStreak != nil {
					if currentStreak.Length() > 1 {
						streaks = append(streaks, *currentStreak)
					}
					currentStreak = nil
				}
			}
		}
	}

	return
}
