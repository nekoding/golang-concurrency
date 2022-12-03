package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	output := "hello"

	wg.Add(1)
	go updateMessage(output)
	wg.Wait()

	if msg != output {
		t.Error("data not match")
	}
}

func TestPrintMessage(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "testing"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Error("data not match")
	}

}

func TestMain(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("error data not match Hello, universe!")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("error data not match Hello, cosmos!")
	}

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("error data not match Hello, universe!")
	}

}
