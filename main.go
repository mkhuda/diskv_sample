package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	diskv_utils "github.com/mkhuda/diskv_wrapper/utils"
	"github.com/peterbourgon/diskv/v3"
)

func main() {
	var version, key, value string

	fmt.Println("::Using Diskv")

	currentDisk := diskvInit()

	fmt.Println("::::Existing")
	existingVersionUse := diskv_utils.VersionUse(currentDisk)
	diskv_utils.ReadByVersions(existingVersionUse, currentDisk)
	diskv_utils.ReadAllKeys(currentDisk)

	fmt.Println("::::version")
	fmt.Print("Enter version (v1, v2, v...): ")

	// Input version name {v1, v2, v...}
	_version, err := fmt.Scanln(&version)
	if err != nil {
		justError(_version, err)
		return
	}

	diskv_utils.WriteVersion(currentDisk, version)

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
	valueReader := bufio.NewReader(os.Stdin)
	valueBytes, _, err := valueReader.ReadLine()
	value = string(valueBytes)
	if err != nil {
		justError(valueBytes, err)
		return
	}

	path := strings.Join([]string{version, key}, "/")

	currentDisk.WriteString(path, value)

}

// Diskv init
func diskvInit() *diskv.Diskv {
	disk := diskv.New(diskv.Options{
		BasePath:          "localdisk",
		AdvancedTransform: diskv_utils.AdvanceTransform,
		InverseTransform:  diskv_utils.InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	return disk
}

func justError(id interface{}, err error) {
	fmt.Printf("Error: %v %v %v", id, os.Stderr, err)
}
