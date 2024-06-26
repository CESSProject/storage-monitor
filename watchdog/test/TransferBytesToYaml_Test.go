package test

import (
	"testing"
)

func TestTransferBytesToYaml(t *testing.T) {
	err := TransferBytesToYaml()
	if err != nil {
		t.Fatal("Test: Transfer bytes to yaml failed")
	}
}
