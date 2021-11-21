package book

type BookRequest struct {
	Title       string ` json:"title" binding:"required" `
	Price       int    ` json:"price" binding:"required" `
	Description string ` json:"description" binding:"required" `
	Rating      int    ` json:"rating" binding:"required" `
}

type BookResponse struct {
	Id          uint    ` json:"ID"`
	Title       string ` json:"TITLE" `
	Price       int    ` json:"PRICE"`
	Description string ` json:"DESC"`
	Rating      int    ` json:"RATING" `
}
