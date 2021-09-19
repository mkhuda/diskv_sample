package utils

import (
	"fmt"
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

func VersionUse(d *diskv.Diskv) string {
	var version string

	for key := range d.Keys(nil) {
		if key == "version" {
			v, _ := d.Read(key)
			version = string(v)
		}
	}
	fmt.Printf("current version of storage is %v\n", version)
	return version
}

func ReadByVersions(v string, d *diskv.Diskv) {
	var values []string

	for key := range d.Keys(nil) {
		version := strings.Split(key, "/")[0]
		if version == v {
			val, _ := d.Read(key)
			values = append(values, string(val))
		}
	}
	fmt.Printf("total values in %v is %v\n", v, len(values))
}

func ReadAllKeys(d *diskv.Diskv) {
	var keyCount int
	for key := range d.Keys(nil) {
		_, err := d.Read(key)
		if err != nil {
			panic(fmt.Sprintf("key %s had no value", key))
		}
		keyCount++
	}
	fmt.Printf("there are %d total keys on the disk \n", keyCount)
}
