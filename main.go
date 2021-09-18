package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/peterbourgon/diskv/v3"
)

func main() {
	var key, value string

	fmt.Println("::Using Diskv")

	currentDisk := initDiskv()

	fmt.Println("::::key")
	fmt.Print("Enter key: ")

	// Input key name
	k, err := fmt.Scanln(&key)
	if err != nil {
		justError(k, err)
		return
	}

	fmt.Println("::::value")
	fmt.Print("Enter the value: ")

	// Input value
	v, err := fmt.Scanln(&value)
	if err != nil {
		justError(v, err)
		return
	}
	writeBuf, sync := bytes.NewBufferString(value), true
	if err := currentDisk.WriteStream(key, writeBuf, sync); err != nil {
		fmt.Printf("Write error: %v", err)
	}
}

func initDiskv() *diskv.Diskv {
	flatTransform := func(s string) []string { return []string{} }

	disk := diskv.New(diskv.Options{
		BasePath:     "localdisk",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

	return disk
}

func justError(id int, err error) {
	fmt.Printf("Error: %v %v %v", id, os.Stderr, err)
}
