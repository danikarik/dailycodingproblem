package problem15

import (
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func pick(stream []int) int {
	var element int
	for i, e := range stream {
		if i == 0 {
			element = e
		} else {
			if random(1, i+1) == 1 {
				element = e
			}
		}
	}
	return element
}
