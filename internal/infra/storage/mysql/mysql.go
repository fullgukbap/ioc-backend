package mysql

import (
	"fmt"
	"ioc-backend/internal/infra/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	*gorm.DB
}

func NewMysql(config *config.Config) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Mysql{db}, nil
}
