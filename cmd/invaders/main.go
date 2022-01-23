package main

// https://www.computerarcheology.com/Arcade/SpaceInvaders/
// https://github.com/JimKnowler/SpaceInvaders8080/blob/master/src/main.cpp
// https://qiita.com/zakuroishikuro/items/15d1a69178895edf9a21

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	i8080 "github.com/hirosuzuki/go8080/cpu"
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
	shiftOffset uint8
	shiftData   uint16
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

func (m *IOPort) Read(addr uint16) byte {
	switch addr & 255 {
	case 3: // shift register read
		return uint8(m.shiftData << uint16(m.shiftOffset&7) >> 8)
	}
	return 0
}

func (m *IOPort) Write(addr uint16, value byte) {
	switch addr & 255 {
	case 2: // shift register offset
		m.shiftOffset = value
	case 4: // push to high byte of shift register
		m.shiftData = (m.shiftData >> 8) | (uint16(value) << 8)
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

	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON
	termios.Cflag &^= unix.CSIZE
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
		fmt.Printf("Usage: %s romfile\n", os.Args[0])
		os.Exit(1)
	}

	binary, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer setRawMode()()

	memory := &Memory{}
	copy(memory.buf[:], binary)

	io := &IOPort{}

	cpu := i8080.Intel8080{Memory: memory, IOPort: io}
	cpu.Reset()
	counter := 0

	width := 112
	height := 64
	var buf []rune = make([]rune, (width+1)*height)
	for y := 0; y < height-1; y++ {
		buf[y*(width+1)+width] = rune(10)
	}

	for !cpu.Halted {
		cpu.Exec(10000, func(p *i8080.Intel8080) {
			// log.Println(p.Status())
		})
		counter += 1

		if counter%2 == 0 {
			cpu.Interrupt(0x08)
		} else {
			cpu.Interrupt(0x10)
		}

		if counter%2 == 0 {
			for i := 0; i < 112; i++ {
				for j := 0; j < 32; j++ {
					m1 := memory.buf[0x2400+i*2*32+j]
					m2 := memory.buf[0x2400+i*2*32+32+j]
					m := (m1 & 0x08 >> 3) | (m1 & 0x04 >> 1) | (m1 & 0x02 << 1) | (m1 & 0x01 << 6) | (m2 & 0x08 >> 0) | (m2 & 0x04 << 2) | (m2 & 0x02 << 4) | (m2 & 0x01 << 7)
					buf[((31-j)*2+1)*(width+1)+i] = rune(0x2800 + int(m))
					m1 = m1 >> 4
					m2 = m2 >> 4
					m = (m1 & 0x08 >> 3) | (m1 & 0x04 >> 1) | (m1 & 0x02 << 1) | (m1 & 0x01 << 6) | (m2 & 0x08 >> 0) | (m2 & 0x04 << 2) | (m2 & 0x02 << 4) | (m2 & 0x01 << 7)
					buf[((31-j)*2+0)*(width+1)+i] = rune(0x2800 + int(m))
				}
			}

			// https://rosettacode.org/wiki/Terminal_control/Hiding_the_cursor#Go
			fmt.Print("\x1b[?25l\033[0;0H")
			fmt.Print(string(buf), counter)
			fmt.Print("\x1b[?25h")
		}

	}
}
