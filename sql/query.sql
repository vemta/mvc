-- name: FindItem :one
SELECT * FROM VMT_Items 
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
WHERE ID = ?;

-- name: CreateItem :exec
INSERT INTO VMT_Items (ID, Title, Description, IsGood, CreatedAt) VALUES (?,?,?,?,?);

-- name: FindOrder :many
SELECT
VMT_Orders.ID OrderID,
VMT_Orders.Price OrderPrice,
VMT_Orders.PaymentMethod OrderPaymentMethod,
VMT_Orders.Status OrderStatus,
VMT_Orders.DiscountRaw OrderDiscountRaw,
VMT_Orders.DiscountPercentual OrderDiscountPercentual,
VMT_Users.Email CustomerEmail,
VMT_Users.FullName CustomerFullName,
VMT_Users.Birthdate CustomerBirthdate,
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
INNER JOIN VMT_Users on VMT_User.Email = VMT_Orders.Customer 
INNER JOIN VMT_OrderDetails ON VMT_OrderDetails.OrderID = VMT_Orders.ID 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
WHERE VMT_Orders.ID = ?;

-- name: FindItemCostHistory :many
SELECT * FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Cost';

-- name: FindItemPriceHistory :many
SELECT * FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Price';