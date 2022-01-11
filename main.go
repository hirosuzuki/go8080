package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/hirosuzuki/go8080/i8080"
	"golang.org/x/sys/unix"
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
}

func pollReadChar() bool {
	var r syscall.FdSet
	var ti syscall.Timeval
	ti.Sec = 0
	ti.Usec = 1
	n, _ := syscall.Select(0, &r, nil, nil, &ti)
	return n > 0
}

func readChar() byte {
	buf := make([]byte, 1)
	bytes, _ := os.Stdin.Read(buf)
	if bytes > 0 {
		return buf[0]
	}
	return 0
}

func putChar(ch byte) {
	var buf []byte = []byte{ch}
	os.Stdout.Write(buf)
}

func (m *IOPort) Read(addr uint16) byte {
	switch addr & 255 {
	case 0: // console status
		if pollReadChar() {
			return 255
		}
		return 0
	case 1: // console data
		ch := readChar()
		return ch
	}
	return 0
}

func (m *IOPort) Write(addr uint16, value byte) {
	switch addr & 255 {
	case 1: // console data
		putChar(value)
	}
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

	termios.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	termios.Oflag &^= unix.OPOST
	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.Cflag &^= unix.CSIZE | unix.PARENB
	termios.Cflag |= unix.CS8
	termios.Cc[unix.VMIN] = 1
	termios.Cc[unix.VTIME] = 0

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
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s binary\n", os.Args[0])
		os.Exit(1)
	}

	binary, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer setRawMode()()

	memory := &Memory{}
	for i := 0; i < len(binary); i++ {
		memory.Write(uint16(i), binary[i])
	}

	io := &IOPort{}

	cpu := i8080.CPU{Memory: memory, IOPort: io}
	cpu.Reset()

	for {
		cpu.Exec(20000, func(p *i8080.CPU) {
		})
		if cpu.Halted {
			break
		}
	}
}
