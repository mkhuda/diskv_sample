package diskv_wrapper

import (
	"github.com/peterbourgon/diskv/v3"
)

func Init() *diskv.Diskv {
	disk := diskv.New(diskv.Options{
		BasePath:          "localdisk",
		AdvancedTransform: AdvanceTransform,
		InverseTransform:  InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	return disk
}

func WriteVersion(d *diskv.Diskv, version string) {
	d.WriteString("version", version)
}

func Write(d *diskv.Diskv, key, value string) {
	d.WriteString(key, value)
}
