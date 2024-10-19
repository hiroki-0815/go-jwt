package initializers

import "github.com/hiroki-0815/go-jwt.git/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
