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
VMT_OrderDetails.Quantity DetailQuantity,
VMT_OrderDetails.DiscountPercentual ItemDiscountPercentual,
VMT_OrderDetails.DiscountRaw ItemDiscountRaw
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
VMT_Orders.DiscountRaw DiscountRaw,
VMT_Orders.DiscountPercentual DiscountPercentual,
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
VMT_OrderDetails.Quantity DetailQuantity,
VMT_OrderDetails.DiscountPercentual ItemDiscountPercentual,
VMT_OrderDetails.DiscountRaw ItemDiscountRaw
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

-- name: FindValidDiscountRulesForItem :many
SELECT * FROM VMT_DiscountRuleItems
INNER JOIN VMT_DiscountRules ON VMT_DiscountRules.ID = VMT_DiscountRuleItems.DiscountRule
WHERE VMT_DiscountRuleItems.Item = ?
AND VMT_DiscountRules.ValidFrom >= ? AND VMT_DiscountRules.ValidUntil <= ?
AND VMT_DiscountRules.AboveValue >= ? AND VMT_DiscountRules.BellowValue <= ?
ORDER BY VMT_DiscountRules.ID;

-- name: FindValidDiscountRulesForOrder :many
SELECT * FROM VMT_DiscountRules
WHERE VMT_DiscountRules.AboveValue >= ? AND BellowValue <= ? 
AND VMT_DiscountRules.ValidFrom <= ? AND VMT_DiscountRules.ValidUntil >= ?;

-- name: CreateDiscountRule :exec
INSERT INTO VMT_DiscountRules 
(ID, Name, DiscountRaw, DiscountPercentual, ApplyFirst, AboveValue, BellowValue, ValidFrom, ValidUntil, Type)
VALUES (?,?,?,?,?,?,?,?,?,?);

-- name: CreateItemForDiscountRule :exec
INSERT INTO VMT_DiscountRuleItems
(DiscountRule, Item)
VALUES (?,?);

-- name: FindActiveDiscountRules :many
SELECT * FROM VMT_DiscountRules WHERE ValidFrom <= ? AND ValidUntil >= ?;

-- name: FindValidItemsForDiscountRuleDetailed :many 
SELECT 
VMT_Items.ID ItemID,
VMT_Items.Title ItemTitle,
VMT_Items.Description ItemDescription,
VMT_Items.IsGood ItemIsGood,
VMT_Items.CreatedAt ItemCreatedAt,
VMT_Items.Category ItemCategory,
VMT_ItemsValuation.LastPrice LastPrice,
VMT_ItemsValuation.LastCost LastCost,
VMT_ItemsValuation.UpdatedAt PriceUpdatedAt
FROM VMT_DiscountRuleItems
INNER JOIN VMT_Items ON VMT_Items.ID = VMT_DiscountRuleItems.Item
INNER JOIN VMT_ItemsValuation ON VMT_ItemsValuation.ItemID = VMT_DiscountRuleItems.Item
WHERE DiscountRule = ?;

-- name: FindValidItemsForDiscountRule :many 
SELECT Item FROM VMT_DiscountRuleItems WHERE DiscountRule = ?;