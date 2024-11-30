package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

func main() {
	var time, memory, threads, length int
	var file string
	flag.IntVar(&time, "time", 1, "time")
	flag.IntVar(&memory, "memory", 1024, "memory")
	flag.IntVar(&threads, "threads", 1, "threads")
	flag.IntVar(&length, "length", 32, "length")
	flag.StringVar(&file, "file", "argon2.key", "file name")
	flag.Parse()

	fmt.Println("Generating argon2id key with parameters:")
	fmt.Println("- time:", time)
	fmt.Println("- memory:", memory)
	fmt.Println("- threads:", threads)
	fmt.Println("- key length:", length)
	fmt.Println("- key file:", file)
	fmt.Println()

	fmt.Print("Enter pass: ")
	fd := int(os.Stdin.Fd())
	pass, err := term.ReadPassword(fd)
	if err != nil {
		fmt.Println("Error entering pass")
	}
	fmt.Println()

	fmt.Print("Enter salt: ")
	salt, err := term.ReadPassword(fd)
	if err != nil {
		fmt.Println("Error entering salt")
	}
	fmt.Println()
	fmt.Println()

	key := argon2.IDKey(pass, salt, uint32(time), uint32(memory), uint8(threads), uint32(length))

	os.WriteFile(file, key, 0600)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file")
	}

	keyTest := make([]byte, length)
	err = binary.Read(f, binary.LittleEndian, keyTest)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	if bytes.Equal(key, keyTest) {
		fmt.Printf("%s key written and verified\n", file)
	} else {
		fmt.Printf("Error verifying key written to %s\n", file)
	}

	f.Close()
}
