package diskv_wrapper

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

// WriteVersion used to write version file under disk. It contain version info that used in current.
func WriteVersion(d *diskv.Diskv, version string) error {
	err := d.WriteString("version", version)
	if err != nil {
		JustError("Write version error: ", err)
	}
	return err
}

// WritePath will return a key to be written under `vx` folder
func WritePath(d *diskv.Diskv, version string, key string) string {
	path := strings.Join([]string{version, key}, "/")

	return path
}

// Write is function to fill the data under keys
func Write(d *diskv.Diskv, key string, value string) error {
	err := d.WriteString(key, value)
	if err != nil {
		JustError("Write error: ", err)
	}
	return err
}
