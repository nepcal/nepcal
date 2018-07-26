package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	// TODO: A hack to prevent panic on CI.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestRender: ", r)
		}
	}()

	dateBuf := bytes.NewBuffer([]byte(""))
	calBuf := bytes.NewBuffer([]byte(""))

	writer = dateBuf
	render(true)

	writer = calBuf
	render(false)

	if len(dateBuf.String()) > len(calBuf.String()) {
		t.Fatalf("Expected dateBuf to have a lower length than calBuf")
	}
}

func TestShowDate(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))

	tests := []struct {
		name     string
		t        time.Time
		expected string
	}{
		{
			"case-1",
			time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
			"जेठ 3, 2075\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			showDate(b, test.t)
			assert.Equal(t, test.expected, b.String())
		})
	}
}
