package core

import (
	"fmt"
	"io"
	"os"
)

func (cpu *CPU) Load(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open ROM file: %w", err)
	}
	defer file.Close()
	romData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("unable to read ROM file: %w", err)
	}

	if len(romData)+0x200 > len(cpu.Mem) {
		return fmt.Errorf("ROM size exceeds available memory")
	}
	copy(cpu.Mem[0x200:], romData)

	return nil
}
