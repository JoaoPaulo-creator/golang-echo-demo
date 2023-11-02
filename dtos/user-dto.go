package dtos

type UserDTO struct {
	Name   string  `json:"name"`
	Age    uint8   `json:"age"`
	Salary float64 `json:"salary"`
}

type UserResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    uint8   `json:"age"`
	Salary float64 `json:"salary"`
}
