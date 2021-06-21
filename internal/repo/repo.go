package repo

import (
	"context"

	"github.com/ozoncp/ocp-prize-api/internal/prize"

	"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"

	"github.com/rs/zerolog/log"
)

const (
	tableName = "prizes"
)

// IRepo for prize
type IRepo interface {
	AddPrizes(ctx context.Context, prizes []prize.Prize) ([]uint64, error)
	UpdatePrize(ctx context.Context, prize prize.Prize) error
	RemovePrize(ctx context.Context, prizeID uint64) (bool, error)
	DescribePrize(ctx context.Context, prizeID uint64) (*prize.Prize, error)
	ListPrizes(ctx context.Context, limit, offset uint64) ([]prize.Prize, error)
}

type repo struct {
	db *sqlx.DB
}

// NewRepo creates new repo with setted database
func NewRepo(db *sqlx.DB) IRepo {
	return &repo{db: db}
}

// AddPrizes to database
func (r *repo) AddPrizes(ctx context.Context, prizes []prize.Prize) ([]uint64, error) {
	log.Printf("Add Prizes to database")
	query := sq.Insert(tableName).
		Columns("link", "issueID").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, prize := range prizes {

		query = query.Values(prize.Link, prize.IssueID)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	ids := make([]uint64, 0)
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// UpdatePrize in database
func (r *repo) UpdatePrize(ctx context.Context, prize prize.Prize) error {
	log.Printf("Add Prizes to database")
	query := sq.Update(tableName).
		Set("link", prize.Link).
		Set("issueID", prize.IssueID).
		Where(squirrel.Eq{"id": prize.ID}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

// RemovePrize from database
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

// DescribePrize gets prize from database by ID
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

// ListPrizes gets list of prises from database
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
