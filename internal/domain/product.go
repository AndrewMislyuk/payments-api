package domain

type Product struct {
	Id string `json:"id" binding:"required"`
}
