package domain

type Bicycle struct {
	ID        string
	Name      string
	Type      string
	Price     float64
	Available bool
}

type Order struct {
	ID         string
	UserID     string
	BicycleID  string
	TotalPrice float64
	Status     string
}

type Rental struct {
	ID        string
	UserID    string
	BicycleID string
	StartTime int64
	EndTime   int64
	Status    string
}
