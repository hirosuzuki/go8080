package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	i8080 "github.com/hirosuzuki/go8080/cpu"
)

//go:embed minicpm.bin
var minicpm []byte

type Memory struct {
	buf [65536]byte
}

func (m *Memory) Read(addr uint16) byte {
	return m.buf[addr]
}

func (m *Memory) Write(addr uint16, value byte) {
	m.buf[addr] = value
}

type IOPort struct {
}

func (m *IOPort) Read(addr uint16) byte {
	return 0
}

func (m *IOPort) Write(addr uint16, value byte) {
	switch addr & 255 {
	case 1: // console data
		var buf []byte = []byte{value}
		os.Stdout.Write(buf)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s binary\n", os.Args[0])
		os.Exit(1)
	}

	binary, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	memory := &Memory{}

	copy(memory.buf[:], minicpm)
	copy(memory.buf[0x100:], binary)

	io := &IOPort{}

	cpu := i8080.Intel8080{Memory: memory, IOPort: io}
	cpu.Reset()
	cpu.Reg.PC = 0x100

	for !cpu.Halted {
		cpu.Exec(20, func(p *i8080.Intel8080) {
			// log.Println(cpu.Status())
		})
	}

}
