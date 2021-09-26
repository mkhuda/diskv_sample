package diskv_wrapper

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

// AdvanceTransform to transform key it provided for *diskv and used as Option in *diskv
func AdvanceTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last],
	}
}

// InverseTransform will inverse all written keys on disk. Used by *diskv to determine all keys on disk.
func InverseTransform(pathKey *diskv.PathKey) (key string) {
	var path, file []string
	file = []string{pathKey.FileName[:len(pathKey.FileName)]}
	path = append(pathKey.Path, file...)
	return strings.Join(path, "/")
}
