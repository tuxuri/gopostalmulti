# This is a fork of github.com/mzimmerman/gopostalmulti

It provides additional configuration, along with change of syntax to allow forking of commands in the background to parse addresses.

It doesn't require the Go library bindings for libpostal, just compiled binary of `cmd/gopostalcmd.c` as `gopostalcmdc` in the search path. Communication is via stdio + msgpack, so you could use generate a pure Go compiled `gopostalmulti`.

# gopostalmulti
libpostal is not multicore, so this forks commands in the background (as needed) to help parse additional addresses

# Usage

```package main

import "gopkg.in/tuxuri/gopostalmulti.v2"
import "fmt"

func main() {
	l := gopostalmulti.Libpostal{
		Path: "/usr/local/bin/gopostalmultic",
	}
	l.Init()
	data := l.Parse("Kuala Lumpur")
	fmt.Printf("%+v\n", data)
}
```
