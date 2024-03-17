package channel_responses

import (
	"context"
	"fmt"
	"github.com/cassa10/go_concurrent_samples/util"
	"log"
	"time"
)

const thirdPartyTimeSlow = time.Millisecond * 300

func Sample3Timeout() {
	logger := log.Default()
	ctx := context.Background()
	context.WithValue(ctx, "sarasa", "hola")
	userID := 666
	timeoutMs := 200
	start := time.Now()
	res, err := fetchUserExist(ctx, timeoutMs, userID)
	took := time.Since(start)
	if err != nil {
		logger.Println(fmt.Sprintf("Error: %v", err))
	}
	logger.Println(fmt.Sprintf("Res: %v", res))
	logger.Println(fmt.Sprintf("Took: %v", took))
}

func fetchUserExist(ctx context.Context, timeoutInMs int, userID int) (bool, error) {
	ctx, cancelFn := context.WithTimeout(ctx, time.Millisecond*time.Duration(timeoutInMs))
	defer cancelFn()
	resCh := make(chan util.SimpleRes[bool])
	go func() {
		val, err := fetchThirdPartyWhichCanBeSlow(userID)
		resCh <- util.SimpleRes[bool]{
			Result: val,
			Error:  err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("timeout when fetch third party")
		case resp := <-resCh:
			return resp.Result, resp.Error
		}
	}
}

func fetchThirdPartyWhichCanBeSlow(userID int) (bool, error) {
	time.Sleep(thirdPartyTimeSlow)
	if userID == 666 {
		return true, nil
	}
	return false, nil
}
