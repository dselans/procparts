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
        fmt.Printf("Found partition: %v Blocks: %v\n", part.Name, part.Blocks)
    }
}

// Found partition: sr0 Blocks: 1048575
// Found partition: sda Blocks: 43008000
// Found partition: sda1 Blocks: 877568
// Found partition: sda2 Blocks: 1 
```
