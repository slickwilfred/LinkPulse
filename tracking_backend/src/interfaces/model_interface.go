package interfaces

import (
	"tracking_backend/src/database"
)

type Model interface {
	InitializeConnection(db *database.DB)
}
