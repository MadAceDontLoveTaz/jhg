package utils

import (
	"math/rand"
	"time"
)

var (
	validNameRunes    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	validNameRunesLen = len(validNameRunes)
	randCh            = make(chan int, 1024)
	src               = rand.NewSource(time.Now().UnixNano())
)

func Init() {
	go func() {
		for {
			randCh <- int(src.Int63() % int64(validNameRunesLen))
		}
	}()
}

func RandomName() string {
	b := make([]rune, 7)
	for i := range b {
		b[i] = validNameRunes[<-randCh]
	}
	return "TekBoot_" + string(b)
}
