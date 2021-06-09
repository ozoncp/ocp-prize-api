package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-prize-api/internal/flusher IFlusher

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-prize-api/internal/repo IRepo

//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-prize-api/internal/saver ISaver
