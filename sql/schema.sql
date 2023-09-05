CREATE TABLE VMT_SystemOptions (
	BaseCurrency VARCHAR(3) NOT NULL
);

CREATE TABLE VMT_ItemCategories (
	ID INT PRIMARY KEY NOT NULL,
	Name VARCHAR(64) NOT NULL
);

CREATE TABLE VMT_CustomerCart (
	Customer VARCHAR(64) NOT NULL,
	Item VARCHAR(64) NOT NULL,
	Quantity INT NOT NULL
);

CREATE TABLE VMT_Customers (
	Email VARCHAR(64) PRIMARY KEY NOT NULL,
	FullName VARCHAR(128) NOT NULL,
	Birthdate DATE NOT NULL
);

CREATE TABLE VMT_Orders (
    ID                 VARCHAR(64) PRIMARY KEY NOT NULL,
	Customer           VARCHAR(64) NOT NULL,
	Price              DECIMAL(10,2) NOT NULL,
	PaymentMethod      INT NOT NULL,
    Status             INT NOT NULL
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
	IsGood      TINYINT(1) NOT NULL,
	CreatedAt   DATETIME NOT NULL,
	Category INT NOT NULL
);

CREATE TABLE VMT_ItemsValuation (
	ItemID VARCHAR(64) PRIMARY KEY NOT NULL,
	LastPrice DECIMAL(10,2) NOT NULL,
	LastCost DECIMAL(10,2) NOT NULL,
	UpdatedAt DATETIME NOT NULL
);

CREATE TABLE VMT_ItemValuationLog (
	Item VARCHAR(64) NOT NULL,
	Price DECIMAL(10,2) NOT NULL,
	ValuationType ENUM("Price", "Cost") NOT NULL,
	ValorizatedAt DATETIME NOT NULL
);

CREATE TABLE VMT_OrderDiscountRules (
	ID VARCHAR(64) PRIMARY KEY NOT NULL,
	Name VARCHAR(128) NOT NULL,
	DiscountRaw DECIMAL(10,2) NOT NULL,
	DiscountPercentual DECIMAL(10,2) NOT NULL,
	ApplyFirst ENUM('RAW', 'PERCENTUAL') NOT NULL,
	ValidFrom DATETIME NOT NULL,
	ValidUntil DATETIME,
	AboveValue DECIMAL(10,2) NOT NULL,
	BellowValue DECIMAL(10,2) NOT NULL
);

CREATE TABLE VMT_ItemDiscountRules (
	ID VARCHAR(64) PRIMARY KEY NOT NULL,
	Name VARCHAR(128) NOT NULL,
	DiscountRaw DECIMAL(10,2) NOT NULL,
	DiscountPercentual DECIMAL(10,2) NOT NULL,
	ApplyFirst ENUM('RAW', 'PERCENTUAL') NOT NULL,
	ValidFrom DATETIME NOT NULL,
	ValidUntil DATETIME,
	AboveValue DECIMAL(10,2) NOT NULL,
	BellowValue DECIMAL(10,2) NOT NULL
);

CREATE TABLE VMT_ItemsOfDiscountRule (
	DiscountRule VARCHAR(64) NOT NULL,
	Item VARCHAR(64) NOT NULL,

	PRIMARY KEY (DiscountRule, Item)
);