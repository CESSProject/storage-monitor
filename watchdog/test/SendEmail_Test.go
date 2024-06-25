package test

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail()
	if err != nil {
		t.Fatal("Send email with html template failed")
	}
}
