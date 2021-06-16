package repo

import (
	"context"

	"github.com/ozoncp/ocp-prize-api/internal/prize"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"

	"github.com/rs/zerolog/log"
)

const (
	tableName = "prizes"
)

// IRepo for prize
type IRepo interface {
	AddPrizes(ctx context.Context, prizes []prize.Prize) (uint64, error)
	RemovePrize(ctx context.Context, prizeID uint64) (bool, error)
	DescribePrize(ctx context.Context, prizeID uint64) (*prize.Prize, error)
	ListPrizes(ctx context.Context, limit, offset uint64) ([]prize.Prize, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) IRepo {
	return &repo{db: db}
}

func (r *repo) AddPrizes(ctx context.Context, prizes []prize.Prize) (uint64, error) {
	log.Printf("Add Prizes to database")
	query := sq.Insert(tableName).
		Columns("link", "issueID").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, prize := range prizes {

		query = query.Values(prize.Link, prize.IssueID)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	return uint64(id), err
}

func (r *repo) RemovePrize(ctx context.Context, prizeID uint64) (bool, error) {

	log.Printf("Remove Prize from database")
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": prizeID}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	var found bool
	err := query.QueryRowContext(ctx).Scan(&found)
	if err != nil {
		log.Printf(err.Error())
	}
	return found, err
}

func (r *repo) DescribePrize(ctx context.Context, prizeID uint64) (*prize.Prize, error) {

	log.Printf("Describe prize")
	query := sq.Select("id", "link", "IssueID").
		From(tableName).
		Where(sq.Eq{"id": prizeID}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var prize prize.Prize
	err := query.QueryRowContext(ctx).Scan(&prize.ID, &prize.Link, &prize.IssueID)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return &prize, nil
}

func (r *repo) ListPrizes(ctx context.Context, limit, offset uint64) ([]prize.Prize, error) {

	log.Printf("List Prizes")
	query := sq.Select("id", "link", "IssueID").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var prizes []prize.Prize
	for rows.Next() {
		var prize prize.Prize
		if err := rows.Scan(&prize.ID, &prize.Link, &prize.IssueID); err != nil {
			log.Printf(err.Error())
			continue
		}
		prizes = append(prizes, prize)
	}
	return prizes, nil
}
