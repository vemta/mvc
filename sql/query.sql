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
WHERE VMT_Orders.ID = ?;

-- name: FindItemCostHistory :many
SELECT * FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Cost';

-- name: FindItemPriceHistory :many
SELECT * FROM VMT_ItemValuationLog WHERE Item = ? AND Type = 'Price';

-- name: FindCustomerOrders :many
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
VMT_Orders.DiscountPercentual DiscountPercentual,
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
FROM VMT_OrderDetails
INNER JOIN VMT_Customers ON VMT_Customers.Email = VMT_Orders.Customer 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
INNER JOIN VMT_Orders ON VMT_Orders.ID = VMT_OrderDetails.OrderID
WHERE VMT_Customers.Email = ? ORDER BY VMT_Orders.ID;

-- name: FindCustomer :one
SELECT * FROM VMT_Customers WHERE Email = ?;

-- name: CreateCustomer :exec
INSERT INTO VMT_Customers (Email, FullName, Birthdate) VALUES (?,?,?);