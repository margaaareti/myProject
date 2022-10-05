package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"myProject/models"
	"myProject/server/repository"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user models.User) (uint16, error) {

	var id uint16
	var mailStatus int
	var userStatus int

	isExistGet := fmt.Sprintf(`SELECT COUNT(id), (SELECT COUNT (id) FROM %[1]s WHERE username = $2) FROM %[1]s WHERE email = $1`, repository.UserTable)
	isExistRow := r.db.QueryRow(ctx, isExistGet, user.Email, user.Username)
	if err := isExistRow.Scan(&mailStatus, &userStatus); err != nil {
		return 0, err
	}

	if userStatus != 0 {
		return 0, errors.New(userAlrExist)
	} else if mailStatus != 0 {
		return 0, errors.New(emailAlrExist)
	} else {
		query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, username,password, email) values($1,$2,$3,$4,$5,$6) RETURNING id", repository.UserTable)
		row := r.db.QueryRow(ctx, query, user.Name, user.Surname, user.Patronymic, user.Username, user.Password, user.Email)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}
	return id, nil
}
