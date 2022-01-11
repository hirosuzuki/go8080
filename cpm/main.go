package main

import (
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
	// https://www.autometer.de/unix4fun/z80pack/
	return m.buf[addr]
}

func (m *Memory) Write(addr uint16, value byte) {
	m.buf[addr] = value
}

/**
 *      Used I/O ports:
 *
 *       0 - console status
 *       1 - console data
 *
 *       2 - printer status
 *       3 - printer data
 *
 *       4 - auxiliary status
 *       5 - auxiliary data
 *
 *      10 - FDC drive
 *      11 - FDC track
 *      12 - FDC sector (low)
 *      13 - FDC command
 *      14 - FDC status
 *
 *      15 - DMA destination address low
 *      16 - DMA destination address high
 */

var DiskImage []byte

type IOPort struct {
	Memory      *Memory
	DiskImage   *[]byte
	ConsBuf     uint8
	FdcDrive    uint8
	FdcTrack    uint8
	FdcSector   uint8
	FdcStatus   uint8
	DmaAddressL uint8
	DmaAddressH uint8
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
	case 10: // FDC drive
		return m.FdcDrive
	case 11: // FDC track
		return m.FdcTrack
	case 12: // FDC sector (low)
		return m.FdcSector
	case 13: // FDC command
	case 14: // FDC status
		return m.FdcStatus
	case 15: // DMA destination address low
		return m.DmaAddressL
	case 16: // DMA destination address high
		return m.DmaAddressH
	}
	return 0
}

func (m *IOPort) Write(addr uint16, value byte) {
	switch addr & 255 {
	case 1: // console data
		putChar(value)
	case 10: // FDC drive
		m.FdcDrive = value
	case 11: // FDC track
		m.FdcTrack = value
	case 12: // FDC sector (low)
		m.FdcSector = value
	case 13: // FDC command
		m.FdcStatus = 0
		if m.FdcTrack > 77 {
			m.FdcStatus = 2
			return
		}
		if m.FdcSector > 26 {
			m.FdcStatus = 3
			return
		}
		pos := (int(m.FdcTrack)*26 + int(m.FdcSector) - 1) * 128
		addr := (uint16(m.DmaAddressH) << 8) + uint16(m.DmaAddressL)
		switch value {
		case 0: // read
			// log.Printf("Disk Read %06X -> %04X\n", pos, addr)
			for i := 0; i < 128; i++ {
				m.Memory.buf[addr+uint16(i)] = (*m.DiskImage)[pos+i]
			}
		case 1: // write
			// log.Printf("Disk Write %06X <- %04X\n", pos, addr)
			for i := 0; i < 128; i++ {
				(*m.DiskImage)[pos+i] = m.Memory.buf[addr+uint16(i)]
			}
		}
	case 14: // FDC status
	case 15: // DMA destination address low
		m.DmaAddressL = value
	case 16: // DMA destination address high
		m.DmaAddressH = value
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

	// https://www.autometer.de/unix4fun/z80pack/ftp/
	diskImage, err := ioutil.ReadFile("cpm13.dsk")
	if err != nil {
		log.Fatal(err)
	}

	defer setRawMode()()

	memory := &Memory{}
	for i := 0; i < 128; i++ {
		memory.Write(uint16(i), diskImage[i])
	}

	io := &IOPort{}
	io.Memory = memory
	io.DiskImage = &diskImage

	cpu := i8080.CPU{Memory: memory, IOPort: io}
	cpu.Reset()

	for {
		cpu.Exec(10000, func(p *i8080.CPU) {
		})
		if cpu.Halted {
			break
		}
	}
}
