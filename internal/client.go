package internal

import (
	"sync"

	"github.com/poteto0/go-nba-sdk/gns"
)

var (
	instance *gns.Client
	once     sync.Once
)

func GetGNSClient() *gns.Client {
	once.Do(func() {
		instance = gns.NewClient(nil)
	})
	return instance
}
