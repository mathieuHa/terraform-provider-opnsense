package tools

import (
	"context"
	"strings"
	"time"
)

const gatewayListCacheWait = 60 * time.Second

var staleGatewayMessages = []string{"not in list", "valid gateway from the list"}

func RetryWhileGatewayListStale(ctx context.Context, op func() error) error {
	deadline := time.Now().Add(gatewayListCacheWait)
	for {
		err := op()
		if err == nil || !staleGatewayError(err) || time.Now().After(deadline) {
			return err
		}
		select {
		case <-ctx.Done():
			return err
		case <-time.After(5 * time.Second):
		}
	}
}

func staleGatewayError(err error) bool {
	for _, m := range staleGatewayMessages {
		if strings.Contains(err.Error(), m) {
			return true
		}
	}
	return false
}
