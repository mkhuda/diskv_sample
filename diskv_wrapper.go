package diskv_wrapper

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

func Init(basePath string) *diskv.Diskv {
	disk := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: AdvanceTransform,
		InverseTransform:  InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	return disk
}

func GetKeyVersion(d *diskv.Diskv, version string, key string) string {
	path := strings.Join([]string{version, key}, "/")

	return path
}
