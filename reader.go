package diskv_wrapper

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
	return version
}

func ReadByVersions(v string, d *diskv.Diskv) []string {
	var values []string

	for key := range d.Keys(nil) {
		version := strings.Split(key, "/")[0]
		if version == v {
			val, _ := d.Read(key)
			values = append(values, string(val))
		}
	}
	return values
}

func ReadAllKeys(d *diskv.Diskv) int {
	var keyCount int
	for key := range d.Keys(nil) {
		_, err := d.Read(key)
		if err != nil {
			panic(fmt.Sprintf("key %s had no value", key))
		}
		keyCount++
	}
	return keyCount
}
