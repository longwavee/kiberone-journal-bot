package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/longwavee/kiberone-journal-bot/internal/config"
	"github.com/longwavee/kiberone-journal-bot/internal/model"
)

type Storage struct {
	db *sql.DB
}

func New(config config.Storage) (*Storage, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func (s *Storage) Worker(id int64) *model.Worker {
	var u model.Worker

	query := "SELECT first_name, last_name, username, tutor_work, assis_work, outwork FROM workers WHERE id = $1"
	s.db.QueryRow(query, id).Scan(u.FirstName, u.LastName, u.Username, u.TutorWork, u.AssisWork, u.Outwork)

	return &u
}
