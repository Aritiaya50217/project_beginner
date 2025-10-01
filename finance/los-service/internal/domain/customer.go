package domain

type Customer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerRepository interface {
	Save(customer *Customer) error
	FindByID(id int64) (*Customer, error)
}
