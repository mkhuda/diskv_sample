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
	disk := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: diskv_wrapper.AdvanceTransform,
		InverseTransform:  diskv_wrapper.InverseTransform,
		CacheSizeMax:      1024 * 1024,
	})

	diskv_wrapper.WriteVersion(disk, "v1")

    path := strings.Join([]string{"v1", "key"}, "/")

	diskv_wrapper.Write(disk, path, "testdata")
}
```
# Credits

Diskv Creator: [Peter Bourgon](https://github.com/peterbourgon)