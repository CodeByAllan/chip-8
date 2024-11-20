# CHIP-8 Emulator

A CHIP-8 emulator written in Go, capable of running classic games from the 1970s with support for graphics, sound, and user input.

CHIP-8 was a simple virtual machine used to program video games in the 1970s. This emulator faithfully replicates its functionality, allowing you to play nostalgic games with modern tools.

---

## Documentation

Learn more about CHIP-8 and how this emulator works:

- [What is CHIP-8?](https://en.wikipedia.org/wiki/CHIP-8): Background information on CHIP-8.
- [Instruction Set](https://github.com/mattmikolay/chip-8/wiki/CHIP%E2%80%908-Instruction-Set): Details on how CHIP-8 programs are executed.
- [How to Write an Emulator](https://multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/): A guide for building your own CHIP-8 emulator.

---

## Usage 

To use the emulator, download the binary from the [Releases](https://github.com/CodeByAllan/chip-8/releases) section. Then run the following command:

```bash
chip8-windows-amd64-v*.*.*.exe -rom path/to/your-rom.ch8
```

## Stack

**Front-end:** [Raylib](https://www.raylib.com/)

**Back-end:** [Go](https://go.dev/)


## Screenshots

![pong](https://github.com/CodeByAllan/chip-8/blob/master/assets/pong.jpg)


## License

[GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/)


## Authors

- [@CodeByAllan](https://www.github.com/CodeByAllan)
