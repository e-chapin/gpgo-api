package models

import (
	"fmt"
	"time"
)

type PracticeSession struct {

	Id int
	SessionDate time.Time
	SessionNotes string
	Completed bool

}

func (ps PracticeSession) String() string {
	return fmt.Sprintf("PracticeSession<%d, %S, %v>", ps.Id, ps.SessionDate, ps.Completed)
}