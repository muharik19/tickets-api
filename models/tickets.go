package models

type RequestTicket struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description"`
}

type Ticket struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	DeletedAt   string  `json:"deletedAt,omitempty"`
}

type ResponseTicket struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    *Ticket `json:"data,omitempty"`
}

type MetaDataTicket struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListTicket struct {
	Tickets *[]Ticket       `json:"tickets"`
	Meta    *MetaDataTicket `json:"meta"`
}

type ResponseTickets struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *ListTicket `json:"data,omitempty"`
}
