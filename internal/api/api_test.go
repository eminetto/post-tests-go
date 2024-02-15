package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"log/slog"
)

type mockHandler struct{}

func (m mockHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestStart(t *testing.T) {
	t.Run("retorna erro ao executar ListenAndServe com porta inválida", func(t *testing.T) {
		logger := slog.Default()
		err := Start(logger, "abacate", mockHandler{})
		assert.Contains(t, err.Error(), "listen tcp: lookup tcp/abacate: nodename nor servname provided, or not known")
	})
}
