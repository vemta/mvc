CREATE TABLE VMT_Orders (
    ID                 VARCHAR(64) PRIMARY KEY NOT NULL,
	Customer           VARCHAR(64) NOT NULL,
	Price              DECIMAL(10,2) NOT NULL,
	PaymentMethod      INT NOT NULL,
    Status             INT NOT NULL,
	DiscountRaw        DECIMAL(10,2) NOT NULL,
	DiscountPercentual DECIMAL(10,2) NOT NULL
);

CREATE TABLE VMT_OrderDetails (
    OrderID VARCHAR(64) NOT NULL,
    Item VARCHAR(64) NOT NULL,
    Quantity INT NOT NULL,

    PRIMARY KEY (OrderID, Item)
);

CREATE TABLE VMT_Items (
    ID          VARCHAR(64) PRIMARY KEY NOT NULL,
	Title       VARCHAR(512) NOT NULL,
	Description VARCHAR(1024) NOT NULL,
	IsGood      BIT NOT NULL,
	CreatedAt   DATETIME NOT NULL
);

