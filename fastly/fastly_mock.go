package fastly

import (
	"github.com/sethvargo/go-fastly/fastly"
	"math/rand"
)

type fastlyMock struct{}

// NewMock creates a fastly mock client
func NewMock() Client {
	return &fastlyMock{}
}

func (f fastlyMock) GetLatestMetrics() (*fastly.RealtimeStatsResponse, error) {
	return &fastly.RealtimeStatsResponse{
		Data: []*fastly.RealtimeData{
			{
				Aggregated: &fastly.Stats{
					Requests: uint64(rand.Intn(10000)),
				},
			},
		},
	}, nil
}
