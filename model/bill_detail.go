package model

type Bill_detail struct {
	Id           string
	Bill         Bill
	Product      Product
	ProductPrice int64
	Qty          int
}
