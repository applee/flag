[![Build Status](https://travis-ci.org/applee/flag.svg?branch=master)](https://travis-ci.org/applee/flag)

# Usage

```go
import "github.com/applee/flag"

var (
    config string
    age int
)

flag.StringVar(&config, flag.DefaultConfigFlagName, "", "path to config toml file")
flag.IntVar(&age, "age", 24, "help message for age")

log.Fatal(flag.Parse())
```