package main

import (
	"strings"
	"sync"
	"testing"
)

func Test_RaceCondition(t *testing.T) {
	var m sync.Mutex

	name = "sasuke"

	wg.Add(1)
	go updateName("audi", &m)
	go updateName("subaru", &m)
	wg.Wait()

	if !strings.Contains(name, "subaru") {
		t.Error("data not match")
	}
}
