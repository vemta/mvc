// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type VmtItemvaluationlogValuationtype string

const (
	VmtItemvaluationlogValuationtypePrice VmtItemvaluationlogValuationtype = "Price"
	VmtItemvaluationlogValuationtypeCost  VmtItemvaluationlogValuationtype = "Cost"
)

func (e *VmtItemvaluationlogValuationtype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VmtItemvaluationlogValuationtype(s)
	case string:
		*e = VmtItemvaluationlogValuationtype(s)
	default:
		return fmt.Errorf("unsupported scan type for VmtItemvaluationlogValuationtype: %T", src)
	}
	return nil
}

type NullVmtItemvaluationlogValuationtype struct {
	VmtItemvaluationlogValuationtype VmtItemvaluationlogValuationtype `json:"vmt_itemvaluationlog_valuationtype"`
	Valid                            bool                             `json:"valid"` // Valid is true if VmtItemvaluationlogValuationtype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVmtItemvaluationlogValuationtype) Scan(value interface{}) error {
	if value == nil {
		ns.VmtItemvaluationlogValuationtype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VmtItemvaluationlogValuationtype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVmtItemvaluationlogValuationtype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VmtItemvaluationlogValuationtype), nil
}

type VmtItem struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Isgood      bool      `json:"isgood"`
	Createdat   time.Time `json:"createdat"`
}

type VmtItemsvaluation struct {
	Itemid             string    `json:"itemid"`
	Lastprice          float64   `json:"lastprice"`
	Lastcost           float64   `json:"lastcost"`
	Discountraw        float64   `json:"discountraw"`
	Discountpercentual float64   `json:"discountpercentual"`
	Updatedat          time.Time `json:"updatedat"`
}

type VmtItemvaluationlog struct {
	Item               string                           `json:"item"`
	Price              float64                          `json:"price"`
	Valuationtype      VmtItemvaluationlogValuationtype `json:"valuationtype"`
	Valorizatedat      time.Time                        `json:"valorizatedat"`
	Discountraw        float64                          `json:"discountraw"`
	Discountpercentual float64                          `json:"discountpercentual"`
}

type VmtOrder struct {
	ID                 string  `json:"id"`
	Customer           string  `json:"customer"`
	Price              float64 `json:"price"`
	Paymentmethod      int32   `json:"paymentmethod"`
	Status             int32   `json:"status"`
	Discountraw        float64 `json:"discountraw"`
	Discountpercentual float64 `json:"discountpercentual"`
}

type VmtOrderdetail struct {
	Orderid  string `json:"orderid"`
	Item     string `json:"item"`
	Quantity int32  `json:"quantity"`
}

type VmtUser struct {
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Birthdate time.Time `json:"birthdate"`
}