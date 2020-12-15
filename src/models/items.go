package models

import "fmt"

type PracticeItem struct {
	tableName   struct{} `pg:"practice_item"` // todo don't think I need this line anymore
	Id          int
	Title        string
	Description string
	Url         string
	Active bool
}

func (p PracticeItem) String() string {
	return fmt.Sprintf("Item<%d %s %v>", p.Id, p.Title, p.Active)
}