package gelx

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicGelXAPI provides the gelX RPC service that can be
// use publicly without security implications.
type PublicGelXAPI struct {
	t        *GelX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicGelXAPI create a new RPC gelX service.
func NewPublicGelXAPI(t *GelX) *PublicGelXAPI {
	api := &PublicGelXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the GelX sub-protocol version.
func (api *PublicGelXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
