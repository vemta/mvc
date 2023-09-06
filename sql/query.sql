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
VMT_Customers.Email CustomerEmail,
VMT_Customers.FullName CustomerFullName,
VMT_Customers.Birthdate CustomerBirthdate,
VMT_Items.ID ItemID,
VMT_Items.Category ItemCategory,
VMT_Items.Title ItemTitle,
VMT_Items.Description ItemDescription,
VMT_Items.IsGood ItemIsGood,
VMT_Items.CreatedAt ItemCreatedAt,
VMT_ItemCategories.ID ItemCategoryId,
VMT_ItemCategories.Name ItemCategoryName,
VMT_ItemsValuation.LastPrice ItemPrice,
VMT_ItemsValuation.LastCost ItemCost,
VMT_ItemsValuation.UpdatedAt ValuationUpdatedAt,
VMT_OrderDetails.Quantity DetailQuantity
FROM VMT_Orders 
INNER JOIN VMT_Customers on VMT_Customers.Email = VMT_Orders.Customer 
INNER JOIN VMT_OrderDetails ON VMT_OrderDetails.OrderID = VMT_Orders.ID 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item
INNER JOIN VMT_ItemCategories ON VMT_ItemCategories.ID = VMTItems.Category
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
VMT_Items.ID ItemID,
VMT_Items.Title ItemTitle,
VMT_Items.Description ItemDescription,
VMT_Items.IsGood ItemIsGood,
VMT_Items.CreatedAt ItemCreatedAt,
VMT_Items.Category ItemCategory,
VMT_ItemCategories.ID ItemCategoryId,
VMT_ItemCategories.Name ItemCategoryName,
VMT_ItemsValuation.LastPrice ItemPrice,
VMT_ItemsValuation.LastCost ItemCost,
VMT_ItemsValuation.UpdatedAt ValuationUpdatedAt,
VMT_OrderDetails.Quantity DetailQuantity
FROM VMT_OrderDetails
INNER JOIN VMT_Customers ON VMT_Customers.Email = VMT_Orders.Customer 
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_OrderDetails.Item
INNER JOIN VMT_ItemCategories ON VMT_ItemCategories.ID = VMTItems.Category
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_Items.ID
INNER JOIN VMT_Orders ON VMT_Orders.ID = VMT_OrderDetails.OrderID
WHERE VMT_Customers.Email = ? ORDER BY VMT_Orders.ID;

-- name: FindCustomer :one
SELECT * FROM VMT_Customers WHERE Email = ?;

-- name: CreateCustomer :exec
INSERT INTO VMT_Customers (Email, FullName, Birthdate) VALUES (?,?,?);

-- name: UpdateOrderStatus :exec
UPDATE VMT_Orders SET Status = ? WHERE ID = ?;

-- name: UpdateItemValorization :exec
UPDATE VMT_ItemsValuation SET LastPrice = ?, LastCost = ?, UpdatedAt = ? WHERE ItemID = ?;

-- name: FindValidOrderDiscountRules :many
SELECT * FROM VMT_OrderDiscountRules
WHERE ValidFrom <= ? AND (ValidUntil >= ? OR ValidUntil IS NULL);

-- name: FindItemDiscountRule :many
SELECT * FROM VMT_ItemsOfDiscountRule
INNER JOIN VMT_ItemDiscountRules ON VMT_ItemDiscountRules.ID = VMT_ItemsOfDiscountRule.DiscountRule
WHERE DiscountRule = ?;

-- name: FindOrderDiscountRule :one
SELECT * FROM VMT_OrderDiscountRules WHERE ID = ?;

-- name: FindAvailableDiscountRulesForItem :many
SELECT * FROM VMT_ItemsOfDiscountRule 
INNER JOIN VMT_ItemDiscountRules ON VMT_ItemDiscountRules.ID = VMT_ItemsOfDiscountRule.DiscountRule
WHERE VMT_ItemsOfDiscountRule.Item = ? OR VMT_ItemDiscountRules.AllItems = 1;

-- name: CreateItemDiscountRule :exec
INSERT INTO VMT_ItemDiscountRules 
(ID, Name, DiscountRaw, DiscountPercentual, ApplyFirst, ValidFrom, ValidUntil, AboveValue, BellowValue, AllItems)
VALUES (?,?,?,?,?,?,?,?,?,?);

-- name: CreateOrderDiscountRule :exec
INSERT INTO VMT_OrderDiscountRules 
(ID, Name, DiscountRaw, DiscountPercentual, ApplyFirst, ValidFrom, ValidUntil, AboveValue, BellowValue)
VALUES (?,?,?,?,?,?,?,?,?);