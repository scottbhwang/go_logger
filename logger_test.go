package go_logger

import (
	"context"
	"github.com/scottbhwang/go_logger/logger"
	"math/big"
	"strconv"
	"sync"
	"testing"
)

func TestSample(t *testing.T) {
	logger.Level = logger.DebugLevel
	logger.Info("hello ")
	logger.Debug("world")
}

func TestPrintStruct(t *testing.T) {
	obj := struct {
		Name   string   `json:"name,omitempty"`
		Amount *big.Int `json:"amount"`
		Email  string   `json:"email,omitempty"`
	}{
		Name:   "Scott",
		Amount: big.NewInt(123567),
		Email:  "scottbhwang@gmail.com",
	}
	logger.Info(logger.SprintPretty(obj))
	logger.Info(logger.Sprint(obj))
}

func TestPrefix(t *testing.T) {
	ctx := logger.AppendPrefix(context.Background(), "id", "AF8LH5")
	logger.WithContext(ctx).Info(" yo~ ")
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		index := i
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			ctx = logger.AppendPrefix(ctx, "index", strconv.Itoa(index))
			logger.WithContext(ctx).Info(" oh~ ")
		}(ctx)
	}
	wg.Wait()
	logger.WithContext(ctx).Info(" done~ ")
}
