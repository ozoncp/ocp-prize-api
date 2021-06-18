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
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/producer"

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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.CreatePrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test multi create prize", func() {
			prizesToAdd := []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}
			descPrizes := make([]*desc.Prize, 0, len(prizesToAdd))
			for _, prize := range prizesToAdd {
				descPrizes = append(descPrizes, &desc.Prize{
					Id:      prize.ID,
					IssueId: prize.IssueID,
					Link:    prize.Link,
				})
			}
			request := &desc.MultiCreatePrizeV1Request{
				Prizes: descPrizes,
			}
			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1).AddRow(2).AddRow(3).AddRow(4).AddRow(5)
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizesToAdd[0].Link, prizesToAdd[0].IssueID,
					prizesToAdd[1].Link, prizesToAdd[1].IssueID,
					prizesToAdd[2].Link, prizesToAdd[2].IssueID,
					prizesToAdd[3].Link, prizesToAdd[3].IssueID,
					prizesToAdd[4].Link, prizesToAdd[4].IssueID).WillReturnRows(rows)
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.MultiCreatePrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.PrizeIds[0]).Should(BeEquivalentTo(1))
		})

		It("Test incorrect multi create prize", func() {
			prizesToAdd := []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}
			descPrizes := make([]*desc.Prize, 0, len(prizesToAdd))
			for _, prize := range prizesToAdd {
				descPrizes = append(descPrizes, &desc.Prize{
					Id:      prize.ID,
					IssueId: prize.IssueID,
					Link:    prize.Link,
				})
			}
			request := &desc.MultiCreatePrizeV1Request{
				Prizes: descPrizes,
			}
			mock.ExpectQuery("INSERT INTO prizes").
				WithArgs(prizesToAdd[0].Link, prizesToAdd[0].IssueID,
					prizesToAdd[1].Link, prizesToAdd[1].IssueID,
					prizesToAdd[2].Link, prizesToAdd[2].IssueID,
					prizesToAdd[3].Link, prizesToAdd[3].IssueID,
					prizesToAdd[4].Link, prizesToAdd[4].IssueID).
				WillReturnError(errors.New("can't multi add prizes"))
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.MultiCreatePrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test update prize", func() {
			request := &desc.UpdatePrizeV1Request{
				Id:      1,
				Link:    "www",
				IssueId: 1,
			}
			mock.ExpectExec("UPDATE prizes SET").
				WithArgs(request.Link, request.IssueId, request.Id).WillReturnResult(sqlmock.NewResult(1, 1))
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.UpdatePrizeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Succeed).Should(BeEquivalentTo(true))
		})

		It("Test update prize", func() {
			request := &desc.UpdatePrizeV1Request{
				Id:      1,
				Link:    "www",
				IssueId: 1,
			}
			mock.ExpectExec("UPDATE prizes SET").
				WithArgs(request.Link, request.IssueId, request.Id).
				WillReturnError(errors.New("can't update prize"))
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.UpdatePrizeV1(ctx, request)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
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
			prod := producer.NewProducer("TestOcpPrizeApi")
			testApi = api.NewOcpPrizeApi(sqlxDB, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.ListPrizeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

	})
})
