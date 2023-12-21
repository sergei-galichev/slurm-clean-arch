package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MIGRATIONS_DIR", "./services/contact/internal/repository/storage/postgres/migrations")
}

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options Options
}

type Options struct {
	DefaultLimit  uint64
	DefaultOffset uint64
}

func New(db *pgxpool.Pool, options Options) (*Repository, error) {
	if err := migrations(db); err != nil {
		return nil, err
	}

	var r = &Repository{
		db:     db,
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	r.SetOptions(options)

	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}

	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool) (err error) {
	ctx := context.Background()
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			err = closeErr
			return
		}
	}()

	dir := viper.GetString("MIGRATIONS_DIR")
	goose.SetTableName("contact_version")
	if err = goose.RunContext(ctx, "up", db, dir); err != nil {
		return fmt.Errorf("goose %s error : %w", "up", err)
	}
	return nil
}
