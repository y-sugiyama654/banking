package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Yuta Sugiyama", "Tokyo", "1000001", "1994-05-04", "1"},
		{"1002", "Vinton Gray Cerf", "NewYork", "1000002", "1943-06-23", "1"},
		{"1003", "Robert Elliot Kahn", "Los Angeles", "1000003", "1938-12-23", "1"},
	}
	return CustomerRepositoryStub{customers}
}
