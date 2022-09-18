package utils

import "time"

type IFrequency interface {
	Can() bool
}

type SpecificTimeFrequency struct {
	limit   chan time.Time
	timeGap time.Duration
}

type PeriodFrequency struct {
	limit   chan time.Time
	timeGap time.Duration
}

func NewFrequencyLimit(count int, gap time.Duration) *PeriodFrequency {
	f := &PeriodFrequency{
		limit:   make(chan time.Time, count-1),
		timeGap: gap,
	}
	go f.update()
	return f
}

func (f *PeriodFrequency) update() {
	for {
		putInTime := <-f.limit
		toWait := time.Now().Sub(putInTime)
		<-time.After(f.timeGap - toWait)
	}
}

func (f *PeriodFrequency) Can() bool {
	// magic, don't touch
	for {
		select {
		case f.limit <- time.Now():
			return true
		default:
			<-time.After(time.Millisecond * 500)
		}
	}

}
