package main

import (
	chip8 "chip-8"
	"flag"
)

func main() {
	romPath := flag.String("rom", "", "Path to ROM file")
	flag.Parse()

	if *romPath == "" {
		println("Please provide the path to the ROM file with the -rom flag.")
		return
	}

	chip8.Chip8(romPath)
}
