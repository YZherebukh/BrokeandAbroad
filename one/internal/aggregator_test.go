package internal_test

import (
	"fmt"
	"testing"

	"github.com/BrokeandAbroad/one/internal"
	"github.com/BrokeandAbroad/one/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate go install github.com/golang/mock/mockgen@v1.6.0
//go:generate mockgen -source aggregator.go -destination ./mock/aggregator_mock.go -package mock

var (
	testResp = []byte("https://google.com")
	errTest  = fmt.Errorf("error test")
)

func TestRun(t *testing.T) {
	t.Run("positive_10_calls", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		client := mock.NewMockClient(ctrl)

		client.EXPECT().Get(gomock.Any(), "https://google.com").Return(testResp, nil).Times(10)

		totalBytes := internal.Run(client, 10)
		assert.Equal(t, int64(len(testResp)*10), totalBytes)
	})
	t.Run("positive_0_calls", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		client := mock.NewMockClient(ctrl)

		totalBytes := internal.Run(client, 0)
		assert.Equal(t, int64(0), totalBytes)
	})

	t.Run("positive_11_calls_one_failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		client := mock.NewMockClient(ctrl)

		client.EXPECT().Get(gomock.Any(), "https://google.com").Return(testResp, nil).Times(10)
		client.EXPECT().Get(gomock.Any(), "https://google.com").Return(testResp, errTest).Times(1)

		totalBytes := internal.Run(client, 11)
		assert.Equal(t, int64(len(testResp)*10), totalBytes)
	})
}
