# procparts

Simple /proc/partitions parser library for Go.

## Usage

```
package main

import (
    procparts "github.com/dselans/procparts"
    "fmt"
    "os"
)

func main() {
    parts, err := procparts.GetPartitions()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for _, part := range parts {
        fmt.Printf("Found partition: %v Size: %v\n", part.Name, part.Blocks)
    }
}
// Found partition: name Size: #blocks
// Found partition: sr0 Size: 1048575
// Found partition: sda Size: 43008000
// Found partition: sda1 Size: 877568
// Found partition: sda2 Size: 1 
```
