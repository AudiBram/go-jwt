package initializer

import "jwt-gin/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
