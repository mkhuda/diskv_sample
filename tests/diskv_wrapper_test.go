package diskv_wrapper

import (
	"testing"

	"github.com/mkhuda/diskv_wrapper"
	"github.com/peterbourgon/diskv/v3"
)

var (
	disk    *diskv.Diskv = diskv_wrapper.Init("disk_test")
	version string       = "v1-test"
	key     string       = "versionkeytest"
	data    string       = "this is string of data"
	path    string       = diskv_wrapper.WritePath(disk, version, key)
)

func write_disk_version(t *testing.T) {
	err := diskv_wrapper.WriteVersion(disk, version)
	if err != nil {
		t.Errorf("Expected can write version")
	}
}

func write_data_string_version(t *testing.T) {
	err := diskv_wrapper.Write(disk, path, data)
	if err != nil {
		t.Errorf("Expected can write data")
	}
}

func TestUsingNew(t *testing.T) {
	vKeyForNew := "version"

	diskv, err := diskv_wrapper.New("disk_test_new")
	if err != nil {
		t.Errorf("Expected can create new diskv: %v", err)
	}

	diskv_wrapper.WriteVersion(diskv.Disk, "v1")

	versionUseOnNew := diskv_wrapper.VersionUse(diskv.Disk)

	if versionUseOnNew != "v1" {
		t.Errorf("Expected key is %v", vKeyForNew)
	}

	errorErase := diskv.Disk.EraseAll()
	if errorErase != nil {
		t.Errorf("Expected erasing %v data: %v", vKeyForNew, errorErase)
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
