package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ldcmleo/Ichtaca/internal/model"
)

type PermissionRepo struct {
	db *sql.DB
}

func NewPermissionrepo(db *sql.DB) *PermissionRepo {
	return &PermissionRepo{
		db: db,
	}
}

func (r *PermissionRepo) Create(p *model.Permission) error {
	query := `
		INSERT INTO permissions (user_id, scope)
		VALUES ($1, $2)
		RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.db.QueryRowContext(ctx, query,
		p.UserID,
		p.Scope,
	).Scan(
		&p.ID,
		&p.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PermissionRepo) GetAll() ([]model.Permission, error) {
	query := "SELECT id, user_id, scope, created_at FROM permissions"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var ps []model.Permission
	for res.Next() {
		var p model.Permission
		err := res.Scan(
			&p.ID,
			&p.UserID,
			&p.Scope,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (r *PermissionRepo) GetByID(id int) (*model.Permission, error) {
	query := "SELECT id, user_id, scope, created_at FROM permissions WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, id)
	var p model.Permission
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.Scope,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PermissionRepo) Update(p *model.Permission) error {
	setClauses := []string{}
	args := []any{}
	index := 1

	if p.ID == 0 {
		return errors.New("permission ID is required")
	}

	if p.UserID != 0 {
		setClauses = append(setClauses, fmt.Sprintf("user_id = $%d", index))
		args = append(args, p.UserID)
		index++
	}

	if p.Scope != "" {
		setClauses = append(setClauses, fmt.Sprintf("scope = $%d", index))
		args = append(args, p.Scope)
		index++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, p.ID)
	query := fmt.Sprintf("UPDATE permissions SET %s WHERE id = $%d",
		strings.Join(setClauses, ", "),
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

func (r *PermissionRepo) Delete(id int) error {
	query := "DELETE FROM permissions WHERE id = $1"
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
