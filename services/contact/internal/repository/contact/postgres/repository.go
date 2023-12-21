package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db     *pgxpool.Pool
	genSQL squirrel.StatementBuilderType

	Options Options
}

type Options struct {
	DefaultLimit  uint64
	DefaultOffset uint64
}

func New(db *pgxpool.Pool, options Options) (*Repository, error) {
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

	if r.Options != options {
		r.Options = options
	}
}
