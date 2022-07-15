package bar

import (
	"errors"
	"fmt"
	"math"
)

var progress int = 0

type Progress struct {
	currentPercent float64
	total          int64
	isFinished     bool
}

func NewProgress(limit *int64) Progress {
	return Progress{total: *limit}
}

func (p *Progress) Step(n int64) error {
	if n > p.total {
		return errors.New("total more than step")
	}

	p.currentPercent = math.Floor(float64(n) * 100 / float64(p.total))
	p.progressPrint(false)

	if p.currentPercent == 100 {
		p.Finish()
	}

	return nil
}

func (p *Progress) progressPrint(done bool) {
	if done {
		fmt.Println("\rCopying Done!")
	} else {
		fmt.Printf("\r%.0f%%", p.currentPercent)
	}
}

func (p *Progress) Finish() error {
	if p.isFinished {
		return errors.New("progress was already finished")
	}
	p.currentPercent = 100
	p.isFinished = true

	p.progressPrint(true)

	return nil
}
