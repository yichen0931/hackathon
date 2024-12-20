package models

type Vendor struct {
	VendorID       string
	VendorName     string
	Address        string
	IsOpen         bool
	IsDiscountOpen bool
	DiscountStart  string
	DiscountEnd    string
	Password       string
	VendorImage    string
}

type Meal struct {
	MealID                    string
	VendorID                  string
	MealName                  string
	Description               string
	Price                     float64
	Availability              bool
	SustainabilityCreditScore int
	VendorImage               string
}

type Rider struct {
	RiderID      string
	RiderName    string
	VehiclePlate string
	Availability bool
}

type Customer struct {
	CustomerID                           string
	CustomerName                         string
	Address                              string
	AccumulatedSustainabilityCreditScore int
	Password                             string
}

type Discount struct {
	MealID        string  `json:"MealID"`
	DiscountPrice float64 `json:"DiscountPrice"`
	Quantity      int     `json:"Quantity"`
}

//const (
//	CART OrderStatus = iota
//	PENDING
//	ORDERRECEIVED
//	GROUPORDER
//	PREPARING
//	PICKED
//	DELIVERED
//)

type Orders struct {
	OrderID         string
	CustomerID      string
	RiderID         string
	OrderStatus     string
	OrderEnd        string
	Total           float64
	DeliveryAddress string
}

type OrderDetail struct {
	OrderID   string
	MealID    string
	MealQty   int
	MealPrice float64
}

type CustomerSessions struct {
	SessionID     string
	UserID        string
	SessionExpiry string
}

type VendorSessions struct {
	SessionID     string
	UserID        string
	SessionExpiry string
}
