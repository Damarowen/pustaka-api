package dto

type Bookinput struct {
	Title string ` json:"title" binding:"required" `
	Price int    ` json:"price" binding:"required" `
	//* dari depan sub_title
	//Subtitle string `json:"sub_title"`
}
