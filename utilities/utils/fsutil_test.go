package utils

import (
	"testing"
	"fmt"
)

func TestGetMountpoint(t *testing.T) {
	GetMountpoint(".")
}

func TestGetRawDevice(t *testing.T) {
	raw, _ := GetRawDevice("")
	fmt.Printf("%v\n", raw)
}

func TestGetDiskPartitions(t *testing.T) {
	mps, _ := GetDiskPartitions(false)
	fmt.Printf("%v\n", mps)
}
