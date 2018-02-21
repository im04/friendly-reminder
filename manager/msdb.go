package manager

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func Open() (* sql.DB, error) {
	return sql.Open("mysql","root:m123698745@(127.0.0.1:3306)/friendly_reminder?parseTime=true")
}