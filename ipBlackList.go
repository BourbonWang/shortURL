package main

import (
	"sync"
	"time"
)

var ipBlacklist sync.Map

func ipControllor() {
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ipBlacklist.Range(func(key, value interface{}) bool {
				ipBlacklist.Delete(key)
				return true
			})
		}
	}
}

func checkIP(ip string) bool {
	if value, ok := ipBlacklist.Load(ip); ok {
		times := value.(int)
		if times > 30 {
			return true
		}
		ipBlacklist.Store(ip, times+1)
	} else {
		ipBlacklist.Store(ip, 1)
	}
	return false
}
