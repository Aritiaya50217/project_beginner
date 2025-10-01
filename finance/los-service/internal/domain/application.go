package domain

type Application struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"`
	Status     string `json:"status"` // pending, approved, rejected
}

type ApplicationRepository interface {
	Save(app *Application) error
	FindByID(id int64) (*Application, error)
}
