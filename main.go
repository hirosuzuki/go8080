package main

import (
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
	i := i8080.CPU{Memory: &Memory{}, IOPort: &IOPort{}}
	i.Exec()
	i8080.Add()
}
