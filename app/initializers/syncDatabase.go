package initializers

import (
	"github.com/kin-downey/sanpaiki-backend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}