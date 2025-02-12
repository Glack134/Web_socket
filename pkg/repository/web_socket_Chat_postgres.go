package repository

import (
	"fmt"

	"github.com/Glack134/web_socket"
	"github.com/jmoiron/sqlx"
)

type web_socketChatPostgres struct {
	db *sqlx.DB
}

func NewWedsocketChatPostgres(db *sqlx.DB) *web_socketChatPostgres {
	return &web_socketChatPostgres{db: db}
}

func (r *web_socketChatPostgres) Create(userId int, list web_socket.WebsocketChat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createChatQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", web_socketChatTable)
	row := tx.QueryRow(createChatQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1 $2)", usersChatsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
