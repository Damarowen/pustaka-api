package DB

import (
	"gorm.io/gorm"
)

type DbConn struct {
	pustaka_api *gorm.DB
}
