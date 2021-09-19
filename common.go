package diskv_wrapper

import (
	"fmt"
	"os"
)

func JustError(id interface{}, err error) {
	fmt.Printf("Error: %v %v %v", id, os.Stderr, err)
}
