package models

import "fmt"

type SessionItem struct {

	Id int
	ItemId int
	SessionId int

}

func (si SessionItem) String() string {
	return fmt.Sprintf("SessionItem<%d %d %d>", si.Id, si.ItemId, si.SessionId)
}