package diskv_wrapper

import (
	"testing"

	"github.com/mkhuda/diskv_wrapper"
	"github.com/peterbourgon/diskv/v3"
)

var (
	disk *diskv.Diskv = diskv_wrapper.Init("disk_test")
)

func write_disk_version() {
	diskv_wrapper.WriteVersion(disk, "v1")
}

func TestVersionUse(t *testing.T) {
	write_disk_version()
	existingVersionUse := diskv_wrapper.VersionUse(disk)
	if len(existingVersionUse) == 0 {
		t.Errorf("Expected has version to use")
	}
}
