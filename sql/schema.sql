CREATE TABLE VMT_Orders (
    ID                 VARCHAR(64) PRIMARY KEY NOT NULL,
	Customer           VARCHAR(64) NOT NULL,
	Price              DECIMAL NOT NULL,
	PaymentMethod      INT NOT NULL,
    Status             INT NOT NULL,
	DiscountRaw        DECIMAL
	DiscountPercentual DECIMAL
);

CREATE TABLE VMT_OrderDetails (
    Order VARCHAR(64) NOT NULL,
    Item VARCHAR(64) NOT NULL,
    Quantity INT NOT NULL,

    PRIMARY KEY (Order, Item)
);

CREATE TABBLE VMT_Items (
    ID          VARCHAR(64) PRIMARY KEY NOT NULL,
	Title       VARCHAR(512) NOT NULL,
	Description VARCHAR(1024),
	IsGood      BIT NOT NULL,
	CreatedAt   DATETIME NOT NULL
);

