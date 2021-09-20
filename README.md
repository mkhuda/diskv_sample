# Description

[Diskv](https://github.com/peterbourgon/diskv) wrapper using versioning method.

# Installing

Install

```bash
$ go get github.com/mkhuda/diskv_wrapper
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/peterbourgon/diskv"
    "github.com/mkhuda/diskv_wrapper"
)

func main() {

    // init disk
	disk := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: diskv_wrapper.AdvanceTransform,
		InverseTransform:  diskv_wrapper.InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

    // write the v1 version into main `version` key
	diskv_wrapper.WriteVersion(disk, "v1")

    // get path to write
    path := diskv_wrapper.GetKeyVersion(currentDisk, version, key)

    // write path of version, key and data into disk
	diskv_wrapper.Write(disk, path, "testdata")
}
```

# Credits

Diskv Creator: [Peter Bourgon](https://github.com/peterbourgon)
