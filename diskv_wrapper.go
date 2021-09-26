package diskv_wrapper

import (
	"errors"

	"github.com/peterbourgon/diskv/v3"
)

type DiskvWrapper struct {
	Disk *diskv.Diskv
}

// New will initialize diskv, if basepath is empty it will return nil, error
func New(basePath string) (*DiskvWrapper, error) {
	if basePath == "" {
		return nil, errors.New("please provide basepath for main store")
	}

	disk := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: AdvanceTransform,
		InverseTransform:  InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	wrapper := &DiskvWrapper{
		Disk: disk,
	}

	return wrapper, nil
}

// Init can be used as single method that will return *diskv.Diskv
func Init(basePath string) *diskv.Diskv {
	disk := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: AdvanceTransform,
		InverseTransform:  InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	return disk
}
