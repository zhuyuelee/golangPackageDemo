package models

import (
	"database/sql"
)

type model struct {
	ID        int64
	UpdatedAt sql.NullTime
	CreatedAt sql.NullTime
}
