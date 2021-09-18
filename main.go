package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

func main() {
	var version, key, value string

	fmt.Println("::Using Diskv")

	currentDisk := diskvInit()

	fmt.Println("::::version")
	fmt.Print("Enter version (v1, v2, v...): ")

	// Input version name {v1, v2, v...}
	_version, err := fmt.Scanln(&version)
	if err != nil {
		justError(_version, err)
		return
	}

	defer currentDisk.WriteString("version", version)

	fmt.Println("::::key")
	fmt.Print("Enter key: ")

	// Input key name
	_key, err := fmt.Scanln(&key)
	if err != nil {
		justError(_key, err)
		return
	}

	fmt.Println("::::value")
	fmt.Print("Enter the value: ")

	// Input value
	_value, err := fmt.Scan(&value)
	if err != nil {
		justError(_value, err)
		return
	}

	path := strings.Join([]string{version, key}, "/")

	currentDisk.WriteString(path, value)
}

// Diskv init
func diskvInit() *diskv.Diskv {
	transform := func(key string) *diskv.PathKey {
		path := strings.Split(key, "/")
		last := len(path) - 1
		return &diskv.PathKey{
			Path:     path[:last],
			FileName: path[last],
		}
	}
	inverseTransform := func(pathKey *diskv.PathKey) (key string) {
		return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-4]
	}

	disk := diskv.New(diskv.Options{
		BasePath:          "localdisk",
		AdvancedTransform: transform,
		InverseTransform:  inverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	return disk
}

func justError(id int, err error) {
	fmt.Printf("Error: %v %v %v", id, os.Stderr, err)
}
