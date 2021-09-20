package diskv_wrapper

import (
	"github.com/peterbourgon/diskv/v3"
)

func WriteVersion(d *diskv.Diskv, version string) error {
	err := d.WriteString("version", version)
	if err != nil {
		JustError("Write version error: ", err)
	}
	return err
}

func Write(d *diskv.Diskv, key string, value string) error {
	err := d.WriteString(key, value)
	if err != nil {
		JustError("Write error: ", err)
	}
	return err
}
