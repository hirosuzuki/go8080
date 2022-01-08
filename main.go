package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hirosuzuki/go8080/i8080"
)

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
	buf [65536]byte
}

func (m *IOPort) Read(addr uint16) byte {
	return m.buf[addr]
}

func (m *IOPort) Write(addr uint16, value byte) {
	m.buf[addr] = value
}

func main() {
	memory := &Memory{}
	cpu := i8080.CPU{Memory: memory, IOPort: &IOPort{}}
	cpu.Reset()
	prog, err := ioutil.ReadFile("sample.bin")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(prog); i++ {
		memory.Write(uint16(i), prog[i])
	}
	cpu.Exec(100, func(p *i8080.CPU) {
		fmt.Println(p.Status())
	})
}
