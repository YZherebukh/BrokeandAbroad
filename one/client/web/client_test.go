package web_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrokeandAbroad/one/client/web"
	"github.com/stretchr/testify/assert"
)

var (
	testURL    = "/test/url"
	testRespOK = []byte(`OK`)
)

func TestGet(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// Test request parameters
			assert.Equal(t, req.URL.String(), testURL)
			// Send response to be tested
			rw.Write(testRespOK)
		}))
		// Close the server when test finishes
		defer server.Close()

		ctx := context.Background()

		b, err := web.New(*server.Client()).Get(ctx, server.URL+testURL)
		assert.Nil(t, err)
		assert.Equal(t, len(testRespOK), len(b))
	})

	t.Run("positive", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// Test request parameters
			assert.Equal(t, req.URL.String(), testURL)
			// Send response to be tested
			rw.Write(testRespOK)
		}))
		// Close the server when test finishes
		defer server.Close()

		ctx := context.Background()

		b, err := web.New(*server.Client()).Get(ctx, server.URL+testURL)
		assert.Nil(t, err)
		assert.Equal(t, len(testRespOK), len(b))
	})
}
