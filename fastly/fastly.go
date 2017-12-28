package fastly

import (
	"github.com/sethvargo/go-fastly/fastly"
	"time"
)

// Client is a interface for fastly real time analytics api.
type Client interface {
	GetLatestMetrics() (*fastly.RealtimeStatsResponse, error)
}

type clientImpl struct {
	serviceID string
	client    *fastly.RTSClient
}

// New creates a fastly api client.
func New(serviceID string) Client {
	return &clientImpl{
		serviceID: serviceID,
		client:    fastly.NewRealtimeStatsClient(),
	}
}

func (f clientImpl) GetLatestMetrics() (*fastly.RealtimeStatsResponse, error) {
	return f.client.GetRealtimeStats(
		&fastly.GetRealtimeStatsInput{
			Service:   f.serviceID,
			Timestamp: uint64(time.Now().Second()),
			Limit:     1,
		})
}
