package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/seponik/fileguard/pkg/fileguard"
)

func main() {
	encrypt := flag.Bool("e", false, "Example: ./fileguard -e hello.txt")
	decrypt := flag.Bool("d", false, "Example: ./fileguard -d hello.txt")
	flag.Parse()

	file := flag.Arg(0)

	if *encrypt && *decrypt {
		flag.Usage()
		os.Exit(2)
	}

	if file == "" {
		fmt.Println("No file provided. Please specify the file you would like to process.")
		flag.Usage()
		os.Exit(2)
	}

	var key string
	fmt.Print("Enter your FileGuard encryption/decryption key: ")
	fmt.Scan(&key)

	switch {
	case *encrypt:
		err := fileguard.EncryptFile(file, key)
		if err != nil {
			fmt.Println(err)
		}
	case *decrypt:
		err := fileguard.DecryptFile(file, key)
		if err != nil {
			fmt.Println(err)
		}
	default:
		flag.Usage()
		return
	}
}
