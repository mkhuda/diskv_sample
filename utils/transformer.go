package utils

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

func AdvanceTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last],
	}
}
func InverseTransform(pathKey *diskv.PathKey) (key string) {
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-4]
}
