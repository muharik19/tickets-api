package models

type RequestUser struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Hp      string `json:"hp" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Hp        string `json:"hp"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

type ResponseUser struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *User  `json:"data,omitempty"`
}

type MetaDataUser struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListUser struct {
	Users *[]User       `json:"users"`
	Meta  *MetaDataUser `json:"meta"`
}

type ResponseUsers struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *ListUser `json:"data,omitempty"`
}
