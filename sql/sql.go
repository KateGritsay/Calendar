package sql

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx"
	"time"
)

type EventSQL struct {
	UUID        int64
	Title       string
	Description string
	CreatedAt   time.Time
	UserId      uint64 `db:"user_id"`
}

type MyCalendar struct{
	db *sqlx.DB
	err error
}

func connector() (*sqlx.DB, error) {
	dns := "postgres://myuser:mypass@localhost:5432/mydb?sslmode=verify­full"   //Create POSTGRES!!!
	db, err := sqlx.Open("pgx", dns)

	if err != nil {
		return db, err
	}

	err = db.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}

func (m *MyCalendar) CreateEvent(ctx context.Context, e EventSQL) (int64, error) {
	var uuid int64

	query := `INSERT INTO event(user_id, title, createdat, description)
			VALUES ($1, $2, $3, $4) RETURNING uuid`

	err := m.db.QueryRowContext(ctx, query, e.UserId, e.Title,e.CreatedAt, e.Description).Scan(&uuid)

	return uuid, err
}

func (m *MyCalendar) UpdateEvent(ctx context.Context, event EventSQL) (EventSQL, error) {
	var updated EventSQL

	query := `UPDATE event 
		SET (user_id, title, description, start, "end", notice_time) = 
			(:userId, :title, :description, :start, :end, :noticeTime)
		WHERE uuid = :uuid
		RETURNING *;`

	rows, err := m.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"uuid":        event.UUID,
		"userId":      event.UserId,
		"title":       event.Title,
		"createdat":	event.CreatedAt,
		"description": event.Description,

	})

	defer rows.Close()

	if err != nil {
		return updated, err
	}


	rows.Next()
	err = rows.StructScan(&updated)

	return updated, err
}

func (m *MyCalendar) GetEvent(ctx context.Context, uuid int64) (EventSQL, error) {
	var event EventSQL
	query := `SELECT * FROM event WHERE uuid = :uuid;`
	rows, err := m.db.NamedQueryContext(ctx, query, map[string]interface{}{"uuid": uuid})

	defer rows.Close()

	if err != nil {
		return event, err
	}

	rows.Next()
	err = rows.StructScan(&event)

	return event, err
}
