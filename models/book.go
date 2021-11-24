package models

import "time"

// type Book struct {
// 	Id uint
// 	Title string
// 	Description string
// 	Price int
// 	Rating int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type Book struct {
	ID          uint   `gorm:"primary_key:auto_increment" json:"ID SAYA"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Price       int    `gorm:"type:int" json:"price"`
	Rating      int  `gorm:"type:int" json:"rating"`
	UserID      uint64 `gorm:"not null" json:"-"`
	//* relasi foreign key userId ke tabel user
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

