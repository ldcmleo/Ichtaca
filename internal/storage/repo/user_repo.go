package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

func (r *UserRepo) GetAll() ([]model.User, error) {
	query := `
		SELECT id, name, last_name, email, common_name, finger_print, is_admin, revoked, created_at 
		FROM users`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
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

		users = append(users, user)
	}

	return users, nil
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
	setCluses := []string{}
	args := []any{}
	index := 1

	if user.ID == "" {
		return fmt.Errorf("user ID is required")
	}

	if user.Name != "" {
		setCluses = append(setCluses, fmt.Sprintf("name = $%d", index))
		args = append(args, user.Name)
		index++
	}

	if user.LastName != "" {
		setCluses = append(setCluses, fmt.Sprintf("last_name = $%d", index))
		args = append(args, user.LastName)
		index++
	}

	if user.Email != "" {
		setCluses = append(setCluses, fmt.Sprintf("email = $%d", index))
		args = append(args, user.Email)
		index++
	}

	if user.CommonName != "" {
		setCluses = append(setCluses, fmt.Sprintf("common_name = $%d", index))
		args = append(args, user.CommonName)
		index++
	}

	if user.FingerPrint != "" {
		setCluses = append(setCluses, fmt.Sprintf("finger_print = $%d", index))
		args = append(args, user.FingerPrint)
		index++
	}

	if user.IsAdmin {
		setCluses = append(setCluses, fmt.Sprintf("is_admin = $%d", index))
		args = append(args, user.IsAdmin)
		index++
	}

	if user.Revoked {
		setCluses = append(setCluses, fmt.Sprintf("revoked = $%d", index))
		args = append(args, user.Revoked)
		index++
	}

	if len(setCluses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, user.ID)
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d",
		strings.Join(setCluses, ", "),
		index,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.db.ExecContext(ctx, query, args...)
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
