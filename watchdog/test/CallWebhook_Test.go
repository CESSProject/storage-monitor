package test

import (
	"testing"
)

func TestCallWebhook(t *testing.T) {
	err := CallWebhook()
	if err != nil {
		t.Fatal("Call webhook failed")
	}
}
