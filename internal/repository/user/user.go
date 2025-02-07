package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"CtrlAltDestiny/internal/entity"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Repository struct {
	db      *pgxpool.Pool
	tracer  trace.Tracer
	builder goqu.DialectWrapper
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:      db,
		tracer:  otel.Tracer("Repository"),
		builder: goqu.Dialect("postgres"),
	}
}

func (r *Repository) GetUser(ctx context.Context, id int) (entity.User, error) {
	ctx, span := r.tracer.Start(ctx, "Repository.GetUser()")
	defer span.End()

	sql, args, err := r.builder.Select().From("users").Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return entity.User{}, fmt.Errorf("build sql: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("query rows: %w", err)
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, err
		}
		return entity.User{}, fmt.Errorf("collect one row: %w", err)
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user entity.User) error {
	ctx, span := r.tracer.Start(ctx, "Repository.CreateUser()")
	defer span.End()

	now := time.Now()

	record := goqu.Record{
		"age":        user.Age,
		"name":       user.Name,
		"social":     user.Social,
		"created_at": now,
		"updated_at": now,
	}

	if user.ID != 0 {
		// To create a user with a specific ID.
		record["id"] = user.ID
	}

	sql, args, err := r.builder.Insert("users").Rows(record).ToSQL()
	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	if _, err := r.db.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, user entity.User) error {
	ctx, span := r.tracer.Start(ctx, "Repository.UpdateUser()")
	defer span.End()

	sql, args, err := r.builder.Update("users").
		Set(goqu.Record{
			"age":        user.Age,
			"name":       user.Name,
			"social":     user.Social,
			"updated_at": time.Now(),
		}).
		Where(goqu.C("id").Eq(user.ID)).ToSQL()
	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	res, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	if res.RowsAffected() == 0 {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	ctx, span := r.tracer.Start(ctx, "Repository.DeleteUser()")
	defer span.End()

	sql, args, err := r.builder.Delete("users").Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	if _, err = r.db.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	return nil
}
