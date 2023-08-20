-- name: FindItem :one
SELECT * FROM VMT_Items WHERE ID = ?;

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
CAST(VMT_Items.IsGood AS TINYINT) ItemIsGood,
VMT_Items.CreatedAt ItemCreatedAt,

VMT_OrderDetails.Quantity DetailQuantity

FROM VMT_Orders 

INNER JOIN VMT_Users on VMT_User.Email = VMT_Orders.Customer 
INNER JOIN VMT_OrderDetails ON VMT_OrderDetails.OrderID = VMT_Orders.ID 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item

WHERE VMT_Orders.ID = ?;