package database

import (
	"fmt"
	"github.com/venomuz/alif-task/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewClient(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.PSQL.Host,
		cfg.PSQL.Port,
		cfg.PSQL.User,
		cfg.PSQL.Password,
		cfg.PSQL.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	return db, err
}
