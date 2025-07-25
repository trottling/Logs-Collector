package elastic

import (
	"context"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v9/esapi"
)

// withRetry executes the provided action with a small retry loop.
// It retries on network errors or 5xx responses with exponential backoff.
func withRetry(ctx context.Context, action func() (*esapi.Response, error)) (*esapi.Response, error) {
	delay := 100 * time.Millisecond
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		// Check if context already cancelled
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		res, err := action()
		if err == nil && res != nil && !res.IsError() {
			return res, nil
		}
		if err == nil && res != nil {
			lastErr = fmt.Errorf("elasticsearch error: %s", res.Status())
			res.Body.Close()
		} else if err != nil {
			lastErr = err
		}
		if attempt == 2 {
			break
		}
		select {
		case <-time.After(delay):
			delay *= 2
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("request failed")
	}
	return nil, lastErr
}
