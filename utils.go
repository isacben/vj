package main

import (
    "fmt"
)

func usage() string {
	return fmt.Sprintf(`vj - JSON viewer version %v

Usage: vj [file]
   or: curl ... | vj

Arguments:
   -h, --help            print help
   -v, --version         print version`, version,
	)
}
