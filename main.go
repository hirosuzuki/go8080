package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hirosuzuki/go8080/i8080"
	"golang.org/x/sys/unix"
)

type Memory struct {
	buf [65536]byte
}

func (m *Memory) Read(addr uint16) byte {
	// https://www.autometer.de/unix4fun/z80pack/
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

func setRawMode() func() {
	// 参考: https://github.com/mattn/go-tty/blob/master/tty_unix.go#L92
	var ioctlReadTermios uint = 0x5401  // syscall.TCGETS
	var ioctlWriteTermios uint = 0x5402 // syscall.TCSETS

	termios, err := unix.IoctlGetTermios(0, ioctlReadTermios)
	if err != nil {
		log.Fatal(err)
	}

	backup := *termios

	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.ISIG

	if err := unix.IoctlSetTermios(0, ioctlWriteTermios, termios); err != nil {
		log.Fatal(err)
	}

	return func() {
		if err = unix.IoctlSetTermios(0, ioctlWriteTermios, &backup); err != nil {
			log.Fatal("Restore Termios Fail:", err)
		}
	}
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

	defer setRawMode()()

	for {
		buf := make([]byte, 1)
		n, err := os.Stdin.Read(buf)
		fmt.Println(buf[0], n, err)
		if buf[0] == 'e' {
			break
		}
	}
}
