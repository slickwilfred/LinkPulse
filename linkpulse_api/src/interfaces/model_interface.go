package interfaces

import (
	"linkpulse_api/src/database"
)
type Model interface {
	InitializeConnection(db *database.DB)
}