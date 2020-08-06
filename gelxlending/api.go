package gelxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)


// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicGelXLendingAPI provides the gelX RPC service that can be
// use publicly without security implications.
type PublicGelXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicGelXLendingAPI create a new RPC gelX service.
func NewPublicGelXLendingAPI(t *Lending) *PublicGelXLendingAPI {
	api := &PublicGelXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicGelXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
