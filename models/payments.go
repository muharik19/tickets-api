package models

type RequestPayment struct {
	UserID   int `json:"userId" validate:"required"`
	TicketID int `json:"ticketId" validate:"required"`
}

type Payment struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Hp          string  `json:"hp"`
	Address     string  `json:"address"`
	TicketID    int     `json:"ticketId"`
	Ticket      string  `json:"ticket"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type ResponsePayment struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *Payment `json:"data,omitempty"`
}

type MetaDataPayment struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListPayment struct {
	Payments *[]Payment       `json:"payments"`
	Meta     *MetaDataPayment `json:"meta"`
}

type ResponsePayments struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *ListPayment `json:"data,omitempty"`
}
