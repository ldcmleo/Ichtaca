package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/ldcmleo/Ichtaca/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user *model.User) error {
	query := `
		INSERT INTO users (name, last_name, email, common_name, finger_print, is_admin, revoked)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query,
		user.Name,
		user.LastName,
		user.Email,
		user.CommonName,
		user.FingerPrint,
		user.IsAdmin,
		user.Revoked,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetAll() ([]*model.User, error) {
	return nil, nil
}

func (r *UserRepo) GetByID(id int) (*model.User, error) {
	query := `
		SELECT id, name, last_name, email, common_name, finger_print, is_admin, revoked, created_at 
		FROM users 
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.CommonName,
		&user.FingerPrint,
		&user.IsAdmin,
		&user.Revoked,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Update(user *model.User) error {
	return nil
}

func (r *UserRepo) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
