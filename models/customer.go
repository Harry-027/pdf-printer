package models

import "github.com/Harry-027/pdf-printer/utils"

// Customer struct ...
type Customer struct {
	Company_Name    string
	Company_Contact string
	Order_ID        string
	Order_Date      string
	Customer_Name   string
	Product_Name    string
	Quantity        string
	Price           string
	Discount        string
	Total           string
}

var CustomerBucket = []byte(utils.CUSTOMER_TYPE)
