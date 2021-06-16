package api_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-prize-api/internal/api"

	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
)

var _ = Describe("Api", func() {

	var (
		ctx     context.Context
		testApi desc.OcpPrizeApiServer
		mock    sqlmock.Sqlmock
		db      *sql.DB
		sqlxDB  *sqlx.DB
		err     error
	)
	BeforeEach(func() {
		ctx = context.Background()
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("test api functions", func() {

		It("Test create prize", func() {
			request := &desc.CreatePrizeV1Request{
				Link:    "www",
				IssueId: 1,
			}
			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1)
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(request.Link, request.IssueId).WillReturnRows(rows)
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.CreatePrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.PrizeId).Should(BeEquivalentTo(1))
		})

		It("Test incorrect create prize", func() {
			request := &desc.CreatePrizeV1Request{
				Link:    "www",
				IssueId: 1,
			}
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(request.Link, request.IssueId).WillReturnError(errors.New("can't insert prize"))
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.CreatePrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test remove prize", func() {
			request := &desc.RemovePrizeV1Request{
				PrizeId: 1,
			}
			rows := sqlmock.NewRows([]string{"found"}).
				AddRow(true)
			mock.ExpectQuery("DELETE FROM prizes").
				WithArgs(request.PrizeId).WillReturnRows(rows)
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.RemovePrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Found).Should(BeEquivalentTo(true))
		})

		It("Test incorrect remove prize", func() {
			request := &desc.RemovePrizeV1Request{
				PrizeId: 1,
			}
			mock.ExpectQuery("DELETE FROM prizes").
				WithArgs(request.PrizeId).WillReturnError(errors.New("can't remove prize"))
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.RemovePrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test describe prize", func() {
			request := &desc.DescribePrizeV1Request{
				PrizeId: 1,
			}
			rows := sqlmock.NewRows([]string{"id", "link", "issueId"}).
				AddRow(request.PrizeId, "www", 2)
			mock.ExpectQuery("SELECT (.+) FROM prizes WHERE").
				WithArgs(request.PrizeId).
				WillReturnRows(rows)
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.DescribePrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Prize.Id).Should(BeEquivalentTo(1))
			Expect(response.Prize.IssueId).Should(BeEquivalentTo(2))
			Expect(response.Prize.Link).Should(BeEquivalentTo("www"))
		})

		It("Test describe prize", func() {
			request := &desc.DescribePrizeV1Request{
				PrizeId: 1,
			}
			mock.ExpectQuery("SELECT (.+) FROM prizes WHERE").
				WithArgs(request.PrizeId).
				WillReturnError(errors.New("can't remove prize"))
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.DescribePrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test list prize", func() {
			request := &desc.ListPrizeV1Request{
				Limit:  3,
				Offset: 0,
			}
			rows := sqlmock.NewRows([]string{"id", "link", "IssueID"}).
				AddRow(1, "www", 2).AddRow(2, "www", 3)
			mock.ExpectQuery("SELECT id, link, IssueID FROM prizes LIMIT 3 OFFSET 0").
				WillReturnRows(rows)
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.ListPrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Prizes[0].Id).Should(BeEquivalentTo(1))
			Expect(response.Prizes[1].Id).Should(BeEquivalentTo(2))
			Expect(response.Prizes[0].IssueId).Should(BeEquivalentTo(2))
			Expect(response.Prizes[1].IssueId).Should(BeEquivalentTo(3))
			Expect(response.Prizes[0].Link).Should(BeEquivalentTo("www"))
			Expect(response.Prizes[0].Link).Should(BeEquivalentTo("www"))
		})

		It("Test incorrect list prize", func() {
			request := &desc.ListPrizeV1Request{
				Limit:  3,
				Offset: 0,
			}
			mock.ExpectQuery("SELECT id, link, IssueID FROM prizes LIMIT 3 OFFSET 0").
				WillReturnError(errors.New("can't list prize"))
			testApi = api.NewOcpPrizeApi(sqlxDB)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.ListPrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

	})
})
