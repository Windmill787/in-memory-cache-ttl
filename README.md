# In memory cache package

In memory storage of `interface{}`\
Access by `string` key\
Delete by `string` key\
Set expiration `time.Duration`

## Example

```go
package main

import "github.com/Windmill787/in-memory-cache-ttl"

func main() {
    // create cache
    cache := inmemorycachettl.NewCache()

    // set new value
    cache.Set("userId", 123, time.Second * 2)

    // get value by key
    cache.Get("userId")

    // delete value by key
    cache.Delete("userId")
}

```
