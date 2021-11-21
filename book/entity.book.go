package book

import "time"

type Book struct { 
	Id uint  
	Title string
	Description string
	Price int
	Rating int
	CreatedAt time.Time
	UpdatedAt time.Time
}