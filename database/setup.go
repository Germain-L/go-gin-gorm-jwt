package database

import (
	"fmt"
	"monolith/models"
	"time"

	"gorm.io/gorm"
)

const maxRetries = 5
const retryInterval = 5 * time.Second

func ConnectToDB(dialector gorm.Dialector, cfg *gorm.Config) (*gorm.DB, error) {
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(dialector, cfg)
		if err == nil {
			db.AutoMigrate(&models.User{}, &models.RefreshToken{}, &models.File{})
			return db, nil
		}

		fmt.Printf("Failed to connect to database (attempt #%d), retrying in %v\n", i+1, retryInterval)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
}
