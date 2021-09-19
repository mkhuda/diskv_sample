package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	diskv_utils "github.com/mkhuda/diskv_wrapper"
)

func main() {
	var version, key, value string
	var basePath = "example_data"

	fmt.Println("::Using Diskv")

	currentDisk := diskv_utils.Init(basePath)

	fmt.Println("::::Existing")
	existingVersionUse := diskv_utils.VersionUse(currentDisk)
	diskv_utils.ReadByVersions(existingVersionUse, currentDisk)
	diskv_utils.ReadAllKeys(currentDisk)

	fmt.Println("::::version")
	fmt.Print("Enter version (v1, v2, v...): ")

	// Input version name {v1, v2, v...}
	_version, err := fmt.Scanln(&version)
	if err != nil {
		diskv_utils.JustError(_version, err)
		return
	}

	diskv_utils.WriteVersion(currentDisk, version)

	fmt.Println("::::key")
	fmt.Print("Enter key: ")

	// Input key name
	_key, err := fmt.Scanln(&key)
	if err != nil {
		diskv_utils.JustError(_key, err)
		return
	}

	fmt.Println("::::value")
	fmt.Print("Enter the value: ")

	// Input value
	valueReader := bufio.NewReader(os.Stdin)
	valueBytes, _, err := valueReader.ReadLine()
	value = string(valueBytes)
	if err != nil {
		diskv_utils.JustError(valueBytes, err)
		return
	}

	path := strings.Join([]string{version, key}, "/")

	currentDisk.WriteString(path, value)

}
