package main

import (
	"flag"

	"github.com/seponik/fileguard/pkg/fileguard"
)

// TODO:
// 1. Ask key from user

func main() {
	encrypt := flag.Bool("e", false, "Example: ./fileguard -e hello.txt")
	decrypt := flag.Bool("d", false, "Example: ./fileguard -d hello.txt")
	flag.Parse()

	file := flag.Arg(0)

	switch {
	case *encrypt && *decrypt:
		flag.Usage()
		return
	case *encrypt:
		fileguard.EncryptFile(file, "SuperSecret")
	case *decrypt:
		fileguard.DecryptFile(file, "SuperSecret")
	default:
		flag.Usage()
		return
	}
}
