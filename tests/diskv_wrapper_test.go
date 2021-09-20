package diskv_wrapper

import (
	"testing"

	"github.com/mkhuda/diskv_wrapper"
	"github.com/peterbourgon/diskv/v3"
)

var (
	disk    *diskv.Diskv = diskv_wrapper.Init("disk_test")
	version string       = "v1"
	key     string       = "versionkeytest"
	data    string       = "this is string of data"
	path    string       = diskv_wrapper.GetKeyVersion(disk, version, key)
)

func write_disk_version(t *testing.T) {
	err := diskv_wrapper.WriteVersion(disk, version)
	if err != nil {
		t.Errorf("Write version error")
	}
}

func write_data_string_version(t *testing.T) {
	err := diskv_wrapper.Write(disk, path, data)
	if err != nil {
		t.Errorf("Write data error")
	}
}

func TestVersionUse(t *testing.T) {
	write_disk_version(t)
	existingVersionUse := diskv_wrapper.VersionUse(disk)

	if len(existingVersionUse) == 0 {
		t.Errorf("Expected has version to use")
	}
}

func TestWriteVersionData(t *testing.T) {
	write_disk_version(t)
	write_data_string_version(t)
	keys := diskv_wrapper.ReadAllKeys(disk)
	if keys == 0 {
		t.Errorf("Expected some keys to read")
	}
}

func TestEraseAll(t *testing.T) {
	write_disk_version(t)
	write_data_string_version(t)
	err := disk.EraseAll()
	if err != nil {
		t.Errorf("Expected erasing all data")
	}
	existingVersionUse := diskv_wrapper.VersionUse(disk)
	if len(existingVersionUse) > 0 {
		t.Errorf("Expected no version stored after erased")
	}
}
