package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlHandler struct {
	Db *gorm.DB
}

func NewMysqlHandler(c *config) (*MysqlHandler, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.user,
		c.password,
		c.port,
		c.database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &MysqlHandler{
		Db: db,
	}, nil
}
