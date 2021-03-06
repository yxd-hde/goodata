package generated

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	//    "github.com/nu7hatch/gouuid"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Category struct {
	CategoryID   int32
	CategoryName string
	Description  string
	Picture      []byte
}

type CustomerDemographic struct {
	CustomerTypeID string
	CustomerDesc   string
}

type Customer struct {
	CustomerID   string
	CompanyName  string
	ContactName  string
	ContactTitle string
	Address      string
	City         string
	Region       string
	PostalCode   string
	Country      string
	Phone        string
	Fax          string
}

type Employee struct {
	EmployeeID      int32
	LastName        string
	FirstName       string
	Title           string
	TitleOfCourtesy string
	BirthDate       time.Time
	HireDate        time.Time
	Address         string
	City            string
	Region          string
	PostalCode      string
	Country         string
	HomePhone       string
	Extension       string
	Photo           []byte
	Notes           string
	ReportsTo       int32
	PhotoPath       string
}

type Order_Detail struct {
	OrderID   int32
	ProductID int32
	UnitPrice float32
	Quantity  int16
	Discount  float32
}

type Order struct {
	OrderID        int32
	CustomerID     string
	EmployeeID     int32
	OrderDate      time.Time
	RequiredDate   time.Time
	ShippedDate    time.Time
	ShipVia        int32
	Freight        float32
	ShipName       string
	ShipAddress    string
	ShipCity       string
	ShipRegion     string
	ShipPostalCode string
	ShipCountry    string
}

type Product struct {
	ProductID       int32
	ProductName     string
	SupplierID      int32
	CategoryID      int32
	QuantityPerUnit string
	UnitPrice       float32
	UnitsInStock    int16
	UnitsOnOrder    int16
	ReorderLevel    int16
	Discontinued    bool
}

type Region struct {
	RegionID          int32
	RegionDescription string
}

type Shipper struct {
	ShipperID   int32
	CompanyName string
	Phone       string
}

type Supplier struct {
	SupplierID   int32
	CompanyName  string
	ContactName  string
	ContactTitle string
	Address      string
	City         string
	Region       string
	PostalCode   string
	Country      string
	Phone        string
	Fax          string
	HomePage     string
}

type Territory struct {
	TerritoryID          string
	TerritoryDescription string
	RegionID             int32
}

type Alphabetical_list_of_product struct {
	ProductID       int32
	ProductName     string
	SupplierID      int32
	CategoryID      int32
	QuantityPerUnit string
	UnitPrice       float32
	UnitsInStock    int16
	UnitsOnOrder    int16
	ReorderLevel    int16
	Discontinued    bool
	CategoryName    string
}

type Category_Sales_for_1997 struct {
	CategoryName  string
	CategorySales float32
}

type Current_Product_List struct {
	ProductID   int32
	ProductName string
}

type Customer_and_Suppliers_by_City struct {
	City         string
	CompanyName  string
	ContactName  string
	Relationship string
}

type Invoice struct {
	ShipName       string
	ShipAddress    string
	ShipCity       string
	ShipRegion     string
	ShipPostalCode string
	ShipCountry    string
	CustomerID     string
	CustomerName   string
	Address        string
	City           string
	Region         string
	PostalCode     string
	Country        string
	Salesperson    string
	OrderID        int32
	OrderDate      time.Time
	RequiredDate   time.Time
	ShippedDate    time.Time
	ShipperName    string
	ProductID      int32
	ProductName    string
	UnitPrice      float32
	Quantity       int16
	Discount       float32
	ExtendedPrice  float32
	Freight        float32
}

type Order_Details_Extended struct {
	OrderID       int32
	ProductID     int32
	ProductName   string
	UnitPrice     float32
	Quantity      int16
	Discount      float32
	ExtendedPrice float32
}

type Order_Subtotal struct {
	OrderID  int32
	Subtotal float32
}

type Orders_Qry struct {
	OrderID        int32
	CustomerID     string
	EmployeeID     int32
	OrderDate      time.Time
	RequiredDate   time.Time
	ShippedDate    time.Time
	ShipVia        int32
	Freight        float32
	ShipName       string
	ShipAddress    string
	ShipCity       string
	ShipRegion     string
	ShipPostalCode string
	ShipCountry    string
	CompanyName    string
	Address        string
	City           string
	Region         string
	PostalCode     string
	Country        string
}

type Product_Sales_for_1997 struct {
	CategoryName string
	ProductName  string
	ProductSales float32
}

type Products_Above_Average_Price struct {
	ProductName string
	UnitPrice   float32
}

type Products_by_Category struct {
	CategoryName    string
	ProductName     string
	QuantityPerUnit string
	UnitsInStock    int16
	Discontinued    bool
}

type Sales_by_Category struct {
	CategoryID   int32
	CategoryName string
	ProductName  string
	ProductSales float32
}

type Sales_Totals_by_Amount struct {
	SaleAmount  float32
	OrderID     int32
	CompanyName string
	ShippedDate time.Time
}

type Summary_of_Sales_by_Quarter struct {
	ShippedDate time.Time
	OrderID     int32
	Subtotal    float32
}

type Summary_of_Sales_by_Year struct {
	ShippedDate time.Time
	OrderID     int32
	Subtotal    float32
}

func Customers() []byte {
	proxyUrl, err := url.Parse("http://icache:80")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	var set = runtime.FuncForPC(reflect.ValueOf(Customers).Pointer()).Name()
	fmt.Printf("set: %v\n", set)
	req, err := http.NewRequest("GET", "http://services.odata.org/V3/Northwind/Northwind.svc/Customers()", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	// req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json;odata=verbose")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return body
}
