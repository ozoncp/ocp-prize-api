package repo_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/repo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repo", func() {
	var (
		testRepo repo.IRepo
		ctx      context.Context
		mock     sqlmock.Sqlmock
		db       *sql.DB
		sqlxDB   *sqlx.DB
		err      error
	)
	BeforeEach(func() {
		ctx = context.Background()
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")
		testRepo = repo.NewRepo(sqlxDB)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("test repo", func() {

		It("Test create prize", func() {
			prizeToAdd := prize.NewPrize(1, 2, "www")
			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1)
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizeToAdd.Link, prizeToAdd.IssueID).WillReturnRows(rows)

			ids, err := testRepo.AddPrizes(ctx, []prize.Prize{prizeToAdd})
			Expect(err).Should(BeNil())
			Expect(ids[0]).Should(BeEquivalentTo(1))
		})

		It("Test incorrect create prize", func() {

			prizeToAdd := prize.NewPrize(1, 2, "www")
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizeToAdd.Link, prizeToAdd.IssueID).WillReturnError(errors.New("can't insert prize"))

			_, err := testRepo.AddPrizes(ctx, []prize.Prize{prizeToAdd})
			Expect(err).ShouldNot(BeNil())
		})

		It("Test multi create prize", func() {
			prizesToAdd := []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1).AddRow(2).AddRow(3).AddRow(4).AddRow(5)
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizesToAdd[0].Link, prizesToAdd[0].IssueID,
					prizesToAdd[1].Link, prizesToAdd[1].IssueID,
					prizesToAdd[2].Link, prizesToAdd[2].IssueID,
					prizesToAdd[3].Link, prizesToAdd[3].IssueID,
					prizesToAdd[4].Link, prizesToAdd[4].IssueID).WillReturnRows(rows)

			ids, err := testRepo.AddPrizes(ctx, prizesToAdd)
			Expect(err).Should(BeNil())
			Expect(ids[0]).Should(BeEquivalentTo(1))
		})

		It("Test incorrect multi create prize", func() {
			prizesToAdd := []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizesToAdd[0].Link, prizesToAdd[0].IssueID,
					prizesToAdd[1].Link, prizesToAdd[1].IssueID,
					prizesToAdd[2].Link, prizesToAdd[2].IssueID,
					prizesToAdd[3].Link, prizesToAdd[3].IssueID,
					prizesToAdd[4].Link, prizesToAdd[4].IssueID).
				WillReturnError(errors.New("can't multi add prizes"))

			_, err := testRepo.AddPrizes(ctx, prizesToAdd)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test update prize", func() {

			prizeToUpdate := prize.NewPrize(1, 2, "www")
			mock.ExpectExec("UPDATE prizes SET").
				WithArgs(prizeToUpdate.Link, prizeToUpdate.IssueID, prizeToUpdate.ID).WillReturnResult(sqlmock.NewResult(1, 1))

			err := testRepo.UpdatePrize(ctx, prizeToUpdate)
			Expect(err).Should(BeNil())
		})

		It("Test update prize", func() {
			prizeToUpdate := prize.NewPrize(1, 2, "www")
			mock.ExpectExec("UPDATE prizes SET").
				WithArgs(prizeToUpdate.Link, prizeToUpdate.IssueID, prizeToUpdate.ID).
				WillReturnError(errors.New("can't update prize"))

			err := testRepo.UpdatePrize(ctx, prizeToUpdate)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test remove prize", func() {
			var prizeID uint64 = 1
			rows := sqlmock.NewRows([]string{"found"}).
				AddRow(true)
			mock.ExpectQuery("DELETE FROM prizes").
				WithArgs(prizeID).WillReturnRows(rows)

			res, err := testRepo.RemovePrize(ctx, prizeID)
			Expect(err).Should(BeNil())
			Expect(res).Should(BeEquivalentTo(true))
		})

		It("Test incorrect remove prize", func() {
			var prizeID uint64 = 1
			mock.ExpectQuery("DELETE FROM prizes").
				WithArgs(prizeID).WillReturnError(errors.New("can't remove prize"))

			res, err := testRepo.RemovePrize(ctx, prizeID)
			Expect(err).ShouldNot(BeNil())
			Expect(res).Should(BeEquivalentTo(false))
		})

		It("Test describe prize", func() {
			var prizeID uint64 = 1
			rows := sqlmock.NewRows([]string{"id", "link", "issueId"}).
				AddRow(prizeID, "www", 2)
			mock.ExpectQuery("SELECT (.+) FROM prizes WHERE").
				WithArgs(prizeID).
				WillReturnRows(rows)

			resPrize, err := testRepo.DescribePrize(ctx, prizeID)

			Expect(err).Should(BeNil())
			Expect(resPrize.ID).Should(BeEquivalentTo(1))
			Expect(resPrize.IssueID).Should(BeEquivalentTo(2))
			Expect(resPrize.Link).Should(BeEquivalentTo("www"))
		})

		It("Test incorrect describe prize", func() {

			var prizeID uint64 = 1
			mock.ExpectQuery("SELECT (.+) FROM prizes WHERE").
				WithArgs(prizeID).
				WillReturnError(errors.New("can't remove prize"))

			_, err := testRepo.DescribePrize(ctx, prizeID)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test list prize", func() {

			var limit uint64 = 3
			var offset uint64 = 0
			rows := sqlmock.NewRows([]string{"id", "link", "IssueID"}).
				AddRow(1, "www", 2).AddRow(2, "www", 3)
			mock.ExpectQuery("SELECT id, link, IssueID FROM prizes LIMIT 3 OFFSET 0").
				WillReturnRows(rows)

			prizes, err := testRepo.ListPrizes(ctx, limit, offset)
			Expect(err).Should(BeNil())
			Expect(prizes[0].ID).Should(BeEquivalentTo(1))
			Expect(prizes[1].ID).Should(BeEquivalentTo(2))
			Expect(prizes[0].IssueID).Should(BeEquivalentTo(2))
			Expect(prizes[1].IssueID).Should(BeEquivalentTo(3))
			Expect(prizes[0].Link).Should(BeEquivalentTo("www"))
			Expect(prizes[0].Link).Should(BeEquivalentTo("www"))
		})

		It("Test incorrect list prize", func() {

			var limit uint64 = 3
			var offset uint64 = 0
			mock.ExpectQuery("SELECT id, link, IssueID FROM prizes LIMIT 3 OFFSET 0").
				WillReturnError(errors.New("can't list prize"))

			_, err := testRepo.ListPrizes(ctx, limit, offset)
			Expect(err).ShouldNot(BeNil())
		})

	})
})
