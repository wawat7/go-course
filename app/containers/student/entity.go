package student

import "time"

type Student struct {
	Id string
	Name string
	PhoneNumber string
	Email string
	Gender string
	IsActive bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FormatRequest struct {
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}

type FormatResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
