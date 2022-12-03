package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestGreeting(t *testing.T) {
	// pindahkan dulu fungsi stdout dari package os ke variabel stdout
	stdout := os.Stdout

	r, w, err := os.Pipe()

	if err != nil {
		panic(err)
	}

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go greeting("Hello", &wg)

	wg.Wait()

	err = w.Close()

	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(r)

	if err != nil {
		panic(err)
	}

	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello") {
		t.Error("program error")
	}

	t.Log("Success")

}
