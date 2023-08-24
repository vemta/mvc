// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createItem = `-- name: CreateItem :exec
INSERT INTO VMT_Items (ID, Title, Description, IsGood, CreatedAt) VALUES (?,?,?,?,?)
`

type CreateItemParams struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Isgood      bool      `json:"isgood"`
	Createdat   time.Time `json:"createdat"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) error {
	_, err := q.db.ExecContext(ctx, createItem,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Isgood,
		arg.Createdat,
	)
	return err
}

const findCustomerOrders = `-- name: FindCustomerOrders :many
SELECT
VMT_Orders.ID OrderId,
VMT_Customers.Email CustomerEmail,
VMT_Customers.FullName CustomerFullName,
VMT_Customers.Birthdate CustomerBirthdate,
VMT_Orders.Customer Customer,
VMT_Orders.Price OrderPrice,
VMT_Orders.PaymentMethod PaymentMethod,
VMT_Orders.Status OrderStatus,
VMT_Orders.DiscountRaw DiscountRaw,
VMT_Orders.DiscountPercentual DiscountPercentual
FROM VMT_Orders
INNER JOIN VMT_Customers ON VMT_Customers.Email = VMT_Orders.Customer 
WHERE VMT_Customers.Email = ? ORDER BY VMT_Orders.ID
`

type FindCustomerOrdersRow struct {
	Orderid            string    `json:"orderid"`
	Customeremail      string    `json:"customeremail"`
	Customerfullname   string    `json:"customerfullname"`
	Customerbirthdate  time.Time `json:"customerbirthdate"`
	Customer           string    `json:"customer"`
	Orderprice         float64   `json:"orderprice"`
	Paymentmethod      int32     `json:"paymentmethod"`
	Orderstatus        int32     `json:"orderstatus"`
	Discountraw        float64   `json:"discountraw"`
	Discountpercentual float64   `json:"discountpercentual"`
}

func (q *Queries) FindCustomerOrders(ctx context.Context, email string) ([]FindCustomerOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, findCustomerOrders, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindCustomerOrdersRow
	for rows.Next() {
		var i FindCustomerOrdersRow
		if err := rows.Scan(
			&i.Orderid,
			&i.Customeremail,
			&i.Customerfullname,
			&i.Customerbirthdate,
			&i.Customer,
			&i.Orderprice,
			&i.Paymentmethod,
			&i.Orderstatus,
			&i.Discountraw,
			&i.Discountpercentual,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findItem = `-- name: FindItem :one
SELECT id, title, description, isgood, createdat, itemid, lastprice, lastcost, discountraw, discountpercentual, updatedat FROM VMT_Items 
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
WHERE ID = ?
`

type FindItemRow struct {
	ID                 string    `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	Isgood             bool      `json:"isgood"`
	Createdat          time.Time `json:"createdat"`
	Itemid             string    `json:"itemid"`
	Lastprice          float64   `json:"lastprice"`
	Lastcost           float64   `json:"lastcost"`
	Discountraw        float64   `json:"discountraw"`
	Discountpercentual float64   `json:"discountpercentual"`
	Updatedat          time.Time `json:"updatedat"`
}

func (q *Queries) FindItem(ctx context.Context, id string) (FindItemRow, error) {
	row := q.db.QueryRowContext(ctx, findItem, id)
	var i FindItemRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Isgood,
		&i.Createdat,
		&i.Itemid,
		&i.Lastprice,
		&i.Lastcost,
		&i.Discountraw,
		&i.Discountpercentual,
		&i.Updatedat,
	)
	return i, err
}

const findItemCostHistory = `-- name: FindItemCostHistory :many
SELECT item, price, valuationtype, valorizatedat, discountraw, discountpercentual FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Cost'
`

func (q *Queries) FindItemCostHistory(ctx context.Context, item string) ([]VmtItemvaluationlog, error) {
	rows, err := q.db.QueryContext(ctx, findItemCostHistory, item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VmtItemvaluationlog
	for rows.Next() {
		var i VmtItemvaluationlog
		if err := rows.Scan(
			&i.Item,
			&i.Price,
			&i.Valuationtype,
			&i.Valorizatedat,
			&i.Discountraw,
			&i.Discountpercentual,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findItemPriceHistory = `-- name: FindItemPriceHistory :many
SELECT item, price, valuationtype, valorizatedat, discountraw, discountpercentual FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Price'
`

func (q *Queries) FindItemPriceHistory(ctx context.Context, item string) ([]VmtItemvaluationlog, error) {
	rows, err := q.db.QueryContext(ctx, findItemPriceHistory, item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VmtItemvaluationlog
	for rows.Next() {
		var i VmtItemvaluationlog
		if err := rows.Scan(
			&i.Item,
			&i.Price,
			&i.Valuationtype,
			&i.Valorizatedat,
			&i.Discountraw,
			&i.Discountpercentual,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOrder = `-- name: FindOrder :many
SELECT
VMT_Orders.ID OrderID,
VMT_Orders.Price OrderPrice,
VMT_Orders.PaymentMethod OrderPaymentMethod,
VMT_Orders.Status OrderStatus,
VMT_Orders.DiscountRaw OrderDiscountRaw,
VMT_Orders.DiscountPercentual OrderDiscountPercentual,
VMT_Customers.Email CustomerEmail,
VMT_Customers.FullName CustomerFullName,
VMT_Customers.Birthdate CustomerBirthdate,
VMT_Items.ID ItemID,
VMT_Items.Title ItemTitle,
VMT_Items.Description ItemDescription,
VMT_Items.IsGood ItemIsGood,
VMT_Items.CreatedAt ItemCreatedAt,
VMT_ItemsValuation.DiscountRaw ItemDiscountRaw,
VMT_ItemsValuation.DiscountPercentual ItemDiscountPercentual,
VMT_ItemsValuation.LastPrice ItemPrice,
VMT_ItemsValuation.LastCost ItemCost,
VMT_OrderDetails.Quantity DetailQuantity
FROM VMT_Orders 
INNER JOIN VMT_Customers on VMT_Customers.Email = VMT_Orders.Customer 
INNER JOIN VMT_OrderDetails ON VMT_OrderDetails.OrderID = VMT_Orders.ID 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
WHERE VMT_Orders.ID = ?
`

type FindOrderRow struct {
	Orderid                 string    `json:"orderid"`
	Orderprice              float64   `json:"orderprice"`
	Orderpaymentmethod      int32     `json:"orderpaymentmethod"`
	Orderstatus             int32     `json:"orderstatus"`
	Orderdiscountraw        float64   `json:"orderdiscountraw"`
	Orderdiscountpercentual float64   `json:"orderdiscountpercentual"`
	Customeremail           string    `json:"customeremail"`
	Customerfullname        string    `json:"customerfullname"`
	Customerbirthdate       time.Time `json:"customerbirthdate"`
	Itemid                  string    `json:"itemid"`
	Itemtitle               string    `json:"itemtitle"`
	Itemdescription         string    `json:"itemdescription"`
	Itemisgood              bool      `json:"itemisgood"`
	Itemcreatedat           time.Time `json:"itemcreatedat"`
	Itemdiscountraw         float64   `json:"itemdiscountraw"`
	Itemdiscountpercentual  float64   `json:"itemdiscountpercentual"`
	Itemprice               float64   `json:"itemprice"`
	Itemcost                float64   `json:"itemcost"`
	Detailquantity          int32     `json:"detailquantity"`
}

func (q *Queries) FindOrder(ctx context.Context, id string) ([]FindOrderRow, error) {
	rows, err := q.db.QueryContext(ctx, findOrder, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindOrderRow
	for rows.Next() {
		var i FindOrderRow
		if err := rows.Scan(
			&i.Orderid,
			&i.Orderprice,
			&i.Orderpaymentmethod,
			&i.Orderstatus,
			&i.Orderdiscountraw,
			&i.Orderdiscountpercentual,
			&i.Customeremail,
			&i.Customerfullname,
			&i.Customerbirthdate,
			&i.Itemid,
			&i.Itemtitle,
			&i.Itemdescription,
			&i.Itemisgood,
			&i.Itemcreatedat,
			&i.Itemdiscountraw,
			&i.Itemdiscountpercentual,
			&i.Itemprice,
			&i.Itemcost,
			&i.Detailquantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
