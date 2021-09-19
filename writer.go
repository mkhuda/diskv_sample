package diskv_wrapper

import (
	"github.com/peterbourgon/diskv/v3"
)

func WriteVersion(d *diskv.Diskv, version string) {
	d.WriteString("version", version)
}

func Write(d *diskv.Diskv, key, value string) {
	d.WriteString(key, value)
}
