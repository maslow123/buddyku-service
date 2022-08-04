package services

import (
	"database/sql"

	"github.com/maslow123/buddyku-users/pkg/utils"
)

type Server struct {
	Jwt utils.JwtWrapper
	DB  *sql.DB
}
