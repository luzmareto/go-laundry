package model

import "time"

type Bill struct {
	Id         string
	BillDate   time.Time
	FinishDate string
	Employee   Employee
	Customer   Customer
	Items      []Bill_detail
}
