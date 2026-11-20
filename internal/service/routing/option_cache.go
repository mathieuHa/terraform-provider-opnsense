package routing

import (
	"context"
	"strings"
	"time"
)

const gatewayListCacheWait = 60 * time.Second

func retryWhileGatewayListStale(ctx context.Context, op func() error) error {
	deadline := time.Now().Add(gatewayListCacheWait)
	for {
		err := op()
		if err == nil || !strings.Contains(err.Error(), "not in list") || time.Now().After(deadline) {
			return err
		}
		select {
		case <-ctx.Done():
			return err
		case <-time.After(5 * time.Second):
		}
	}
}
