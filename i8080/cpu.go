package i8080

import (
	"fmt"
	"strings"
)

var FlagTable []uint8 = []uint8{
	0x46, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
}

type Opcode struct {
	bytes    int
	mnemonic string
}

var OpcodeTable []Opcode = []Opcode{
	/* 00 00000000 */ {1, "NOP"},
	/* 01 00000001 */ {3, "LXI B,nn"},
	/* 02 00000010 */ {1, "STAX B"},
	/* 03 00000011 */ {1, "INX B"},
	/* 04 00000100 */ {1, "INR B"},
	/* 05 00000101 */ {1, "DCR B"},
	/* 06 00000110 */ {2, "MVI B,n"},
	/* 07 00000111 */ {1, "RLC"},
	/* 08 00001000 */ {1, "-"},
	/* 09 00001001 */ {1, "DAD B"},
	/* 0A 00001010 */ {1, "LDAX B"},
	/* 0B 00001011 */ {1, "DCX B"},
	/* 0C 00001100 */ {1, "INR C"},
	/* 0D 00001101 */ {1, "DCR C"},
	/* 0E 00001110 */ {2, "MVI C,n"},
	/* 0F 00001111 */ {1, "RRC"},
	/* 10 00010000 */ {1, "-"},
	/* 11 00010001 */ {3, "LXI D,nn"},
	/* 12 00010010 */ {1, "STAX D"},
	/* 13 00010011 */ {1, "INX D"},
	/* 14 00010100 */ {1, "INR D"},
	/* 15 00010101 */ {1, "DCR D"},
	/* 16 00010110 */ {2, "MVI D,n"},
	/* 17 00010111 */ {1, "RAL"},
	/* 18 00011000 */ {1, "-"},
	/* 19 00011001 */ {1, "DAD D"},
	/* 1A 00011010 */ {1, "LDAX D"},
	/* 1B 00011011 */ {1, "DCX D"},
	/* 1C 00011100 */ {1, "INR E"},
	/* 1D 00011101 */ {1, "DCR E"},
	/* 1E 00011110 */ {2, "MVI E,n"},
	/* 1F 00011111 */ {1, "RAR"},
	/* 20 00100000 */ {1, "-"},
	/* 21 00100001 */ {3, "LXI H,nn"},
	/* 22 00100010 */ {3, "SHLD nn"},
	/* 23 00100011 */ {1, "INX H"},
	/* 24 00100100 */ {1, "INR H"},
	/* 25 00100101 */ {1, "DCR H"},
	/* 26 00100110 */ {2, "MVI H,n"},
	/* 27 00100111 */ {1, "DAA"},
	/* 28 00101000 */ {1, "-"},
	/* 29 00101001 */ {1, "DAD H"},
	/* 2A 00101010 */ {3, "LHLD nn"},
	/* 2B 00101011 */ {1, "DCX H"},
	/* 2C 00101100 */ {1, "INR L"},
	/* 2D 00101101 */ {1, "DCR L"},
	/* 2E 00101110 */ {2, "MVI L,n"},
	/* 2F 00101111 */ {1, "CMA"},
	/* 30 00110000 */ {1, "-"},
	/* 31 00110001 */ {3, "LXI SP,nn"},
	/* 32 00110010 */ {3, "STA nn"},
	/* 33 00110011 */ {1, "INX SP"},
	/* 34 00110100 */ {1, "INR M"},
	/* 35 00110101 */ {1, "DCR M"},
	/* 36 00110110 */ {2, "MVI M,n"},
	/* 37 00110111 */ {1, "STC"},
	/* 38 00111000 */ {1, "-"},
	/* 39 00111001 */ {1, "DAD SP"},
	/* 3A 00111010 */ {3, "LDA nn"},
	/* 3B 00111011 */ {1, "DCX SP"},
	/* 3C 00111100 */ {1, "INR A"},
	/* 3D 00111101 */ {1, "DCR A"},
	/* 3E 00111110 */ {2, "MVI A,n"},
	/* 3F 00111111 */ {1, "CMC"},
	/* 40 01000000 */ {1, "MOV B,B"},
	/* 41 01000001 */ {1, "MOV B,C"},
	/* 42 01000010 */ {1, "MOV B,D"},
	/* 43 01000011 */ {1, "MOV B,E"},
	/* 44 01000100 */ {1, "MOV B,H"},
	/* 45 01000101 */ {1, "MOV B,L"},
	/* 46 01000110 */ {1, "MOV B,M"},
	/* 47 01000111 */ {1, "MOV B,A"},
	/* 48 01001000 */ {1, "MOV C,B"},
	/* 49 01001001 */ {1, "MOV C,C"},
	/* 4A 01001010 */ {1, "MOV C,D"},
	/* 4B 01001011 */ {1, "MOV C,E"},
	/* 4C 01001100 */ {1, "MOV C,H"},
	/* 4D 01001101 */ {1, "MOV C,L"},
	/* 4E 01001110 */ {1, "MOV C,M"},
	/* 4F 01001111 */ {1, "MOV C,A"},
	/* 50 01010000 */ {1, "MOV D,B"},
	/* 51 01010001 */ {1, "MOV D,C"},
	/* 52 01010010 */ {1, "MOV D,D"},
	/* 53 01010011 */ {1, "MOV D,E"},
	/* 54 01010100 */ {1, "MOV D,H"},
	/* 55 01010101 */ {1, "MOV D,L"},
	/* 56 01010110 */ {1, "MOV D,M"},
	/* 57 01010111 */ {1, "MOV D,A"},
	/* 58 01011000 */ {1, "MOV E,B"},
	/* 59 01011001 */ {1, "MOV E,C"},
	/* 5A 01011010 */ {1, "MOV E,D"},
	/* 5B 01011011 */ {1, "MOV E,E"},
	/* 5C 01011100 */ {1, "MOV E,H"},
	/* 5D 01011101 */ {1, "MOV E,L"},
	/* 5E 01011110 */ {1, "MOV E,M"},
	/* 5F 01011111 */ {1, "MOV E,A"},
	/* 60 01100000 */ {1, "MOV H,B"},
	/* 61 01100001 */ {1, "MOV H,C"},
	/* 62 01100010 */ {1, "MOV H,D"},
	/* 63 01100011 */ {1, "MOV H,E"},
	/* 64 01100100 */ {1, "MOV H,H"},
	/* 65 01100101 */ {1, "MOV H,L"},
	/* 66 01100110 */ {1, "MOV H,M"},
	/* 67 01100111 */ {1, "MOV H,A"},
	/* 68 01101000 */ {1, "MOV L,B"},
	/* 69 01101001 */ {1, "MOV L,C"},
	/* 6A 01101010 */ {1, "MOV L,D"},
	/* 6B 01101011 */ {1, "MOV L,E"},
	/* 6C 01101100 */ {1, "MOV L,H"},
	/* 6D 01101101 */ {1, "MOV L,L"},
	/* 6E 01101110 */ {1, "MOV L,M"},
	/* 6F 01101111 */ {1, "MOV L,A"},
	/* 70 01110000 */ {1, "MOV M,B"},
	/* 71 01110001 */ {1, "MOV M,C"},
	/* 72 01110010 */ {1, "MOV M,D"},
	/* 73 01110011 */ {1, "MOV M,E"},
	/* 74 01110100 */ {1, "MOV M,H"},
	/* 75 01110101 */ {1, "MOV M,L"},
	/* 76 01110110 */ {1, "HLT"},
	/* 77 01110111 */ {1, "MOV M,A"},
	/* 78 01111000 */ {1, "MOV A,B"},
	/* 79 01111001 */ {1, "MOV A,C"},
	/* 7A 01111010 */ {1, "MOV A,D"},
	/* 7B 01111011 */ {1, "MOV A,E"},
	/* 7C 01111100 */ {1, "MOV A,H"},
	/* 7D 01111101 */ {1, "MOV A,L"},
	/* 7E 01111110 */ {1, "MOV A,M"},
	/* 7F 01111111 */ {1, "MOV A,A"},
	/* 80 10000000 */ {1, "ADD B"},
	/* 81 10000001 */ {1, "ADD C"},
	/* 82 10000010 */ {1, "ADD D"},
	/* 83 10000011 */ {1, "ADD E"},
	/* 84 10000100 */ {1, "ADD H"},
	/* 85 10000101 */ {1, "ADD L"},
	/* 86 10000110 */ {1, "ADD M"},
	/* 87 10000111 */ {1, "ADD A"},
	/* 88 10001000 */ {1, "ADC B"},
	/* 89 10001001 */ {1, "ADC C"},
	/* 8A 10001010 */ {1, "ADC D"},
	/* 8B 10001011 */ {1, "ADC E"},
	/* 8C 10001100 */ {1, "ADC H"},
	/* 8D 10001101 */ {1, "ADC L"},
	/* 8E 10001110 */ {1, "ADC M"},
	/* 8F 10001111 */ {1, "ADC A"},
	/* 90 10010000 */ {1, "SUB B"},
	/* 91 10010001 */ {1, "SUB C"},
	/* 92 10010010 */ {1, "SUB D"},
	/* 93 10010011 */ {1, "SUB E"},
	/* 94 10010100 */ {1, "SUB H"},
	/* 95 10010101 */ {1, "SUB L"},
	/* 96 10010110 */ {1, "SUB M"},
	/* 97 10010111 */ {1, "SUB A"},
	/* 98 10011000 */ {1, "SBB B"},
	/* 99 10011001 */ {1, "SBB C"},
	/* 9A 10011010 */ {1, "SBB D"},
	/* 9B 10011011 */ {1, "SBB E"},
	/* 9C 10011100 */ {1, "SBB H"},
	/* 9D 10011101 */ {1, "SBB L"},
	/* 9E 10011110 */ {1, "SBB M"},
	/* 9F 10011111 */ {1, "SBB A"},
	/* A0 10100000 */ {1, "ANA B"},
	/* A1 10100001 */ {1, "ANA C"},
	/* A2 10100010 */ {1, "ANA D"},
	/* A3 10100011 */ {1, "ANA E"},
	/* A4 10100100 */ {1, "ANA H"},
	/* A5 10100101 */ {1, "ANA L"},
	/* A6 10100110 */ {1, "ANA M"},
	/* A7 10100111 */ {1, "ANA A"},
	/* A8 10101000 */ {1, "XRA B"},
	/* A9 10101001 */ {1, "XRA C"},
	/* AA 10101010 */ {1, "XRA D"},
	/* AB 10101011 */ {1, "XRA E"},
	/* AC 10101100 */ {1, "XRA H"},
	/* AD 10101101 */ {1, "XRA L"},
	/* AE 10101110 */ {1, "XRA M"},
	/* AF 10101111 */ {1, "XRA A"},
	/* B0 10110000 */ {1, "ORA B"},
	/* B1 10110001 */ {1, "ORA C"},
	/* B2 10110010 */ {1, "ORA D"},
	/* B3 10110011 */ {1, "ORA E"},
	/* B4 10110100 */ {1, "ORA H"},
	/* B5 10110101 */ {1, "ORA L"},
	/* B6 10110110 */ {1, "ORA M"},
	/* B7 10110111 */ {1, "ORA A"},
	/* B8 10111000 */ {1, "CMP B"},
	/* B9 10111001 */ {1, "CMP C"},
	/* BA 10111010 */ {1, "CMP D"},
	/* BB 10111011 */ {1, "CMP E"},
	/* BC 10111100 */ {1, "CMP H"},
	/* BD 10111101 */ {1, "CMP L"},
	/* BE 10111110 */ {1, "CMP M"},
	/* BF 10111111 */ {1, "CMP A"},
	/* C0 11000000 */ {1, "RNZ"},
	/* C1 11000001 */ {1, "POP B"},
	/* C2 11000010 */ {3, "JNZ nn"},
	/* C3 11000011 */ {3, "JMP nn"},
	/* C4 11000100 */ {3, "CNZ nn"},
	/* C5 11000101 */ {1, "PUSH B"},
	/* C6 11000110 */ {2, "ADI n"},
	/* C7 11000111 */ {1, "RST 0"},
	/* C8 11001000 */ {1, "RZ"},
	/* C9 11001001 */ {1, "RET"},
	/* CA 11001010 */ {3, "JZ nn"},
	/* CB 11001011 */ {1, "-"},
	/* CC 11001100 */ {3, "CZ nn"},
	/* CD 11001101 */ {3, "CALL nn"},
	/* CE 11001110 */ {2, "ACI n"},
	/* CF 11001111 */ {1, "RST 1"},
	/* D0 11010000 */ {1, "RNC"},
	/* D1 11010001 */ {1, "POP D"},
	/* D2 11010010 */ {3, "JNC nn"},
	/* D3 11010011 */ {2, "OUT n"},
	/* D4 11010100 */ {3, "CNC nn"},
	/* D5 11010101 */ {1, "PUSH D"},
	/* D6 11010110 */ {2, "SUI n"},
	/* D7 11010111 */ {1, "RST 2"},
	/* D8 11011000 */ {1, "RC"},
	/* D9 11011001 */ {1, "-"},
	/* DA 11011010 */ {3, "JC nn"},
	/* DB 11011011 */ {2, "IN n"},
	/* DC 11011100 */ {3, "CC nn"},
	/* DD 11011101 */ {1, "-"},
	/* DE 11011110 */ {2, "SBI n"},
	/* DF 11011111 */ {1, "RST 3"},
	/* E0 11100000 */ {1, "RPO"},
	/* E1 11100001 */ {1, "POP H"},
	/* E2 11100010 */ {3, "JPO nn"},
	/* E3 11100011 */ {1, "XTHL"},
	/* E4 11100100 */ {3, "CPO nn"},
	/* E5 11100101 */ {1, "PUSH H"},
	/* E6 11100110 */ {2, "ANI n"},
	/* E7 11100111 */ {1, "RST 4"},
	/* E8 11101000 */ {1, "RPE"},
	/* E9 11101001 */ {1, "PCHL"},
	/* EA 11101010 */ {3, "JPE nn"},
	/* EB 11101011 */ {1, "XCHG"},
	/* EC 11101100 */ {3, "CPE nn"},
	/* ED 11101101 */ {1, "-"},
	/* EE 11101110 */ {2, "XRI n"},
	/* EF 11101111 */ {1, "RST 5"},
	/* F0 11110000 */ {1, "RP"},
	/* F1 11110001 */ {1, "POP PSW"},
	/* F2 11110010 */ {3, "JP nn"},
	/* F3 11110011 */ {1, "DI"},
	/* F4 11110100 */ {3, "CP nn"},
	/* F5 11110101 */ {1, "PUSH PSW"},
	/* F6 11110110 */ {2, "ORI n"},
	/* F7 11110111 */ {1, "RST 6"},
	/* F8 11111000 */ {1, "RM"},
	/* F9 11111001 */ {1, "SPHL"},
	/* FA 11111010 */ {3, "JM nn"},
	/* FB 11111011 */ {1, "EI"},
	/* FC 11111100 */ {3, "CM nn"},
	/* FD 11111101 */ {1, "-"},
	/* FE 11111110 */ {2, "CPI n"},
	/* FF 11111111 */ {1, "RST 7"},
}

type IO64K interface {
	Read(addr uint16) byte
	Write(addr uint16, value byte)
}

type CPU struct {
	Memory       IO64K
	IOPort       IO64K
	CanInterrupt bool
	Halted       bool
	FetchCount   int
	Instructions int
	Reg          struct {
		PC uint16
		SP uint16
		B  uint8
		C  uint8
		D  uint8
		E  uint8
		H  uint8
		L  uint8
		A  uint8
		F  uint8
	}
}

func (p *CPU) Read16(addr uint16) uint16 {
	nn := uint16(p.Memory.Read(addr)) | uint16(p.Memory.Read(addr+1))<<8
	return nn
}

func (p *CPU) Write16(addr uint16, v uint16) {
	p.Memory.Write(addr, uint8(v))
	p.Memory.Write(addr+1, uint8(v>>8))
}

func (p *CPU) Fetch8() uint8 {
	n := p.Memory.Read(p.Reg.PC)
	p.Reg.PC++
	return n
}

func (p *CPU) Fetch16() uint16 {
	nn := uint16(p.Memory.Read(p.Reg.PC)) | uint16(p.Memory.Read(p.Reg.PC+1))<<8
	p.Reg.PC += 2
	return nn
}

func (p *CPU) Pop16() uint16 {
	nn := uint16(p.Memory.Read(p.Reg.SP)) | uint16(p.Memory.Read(p.Reg.SP+1))<<8
	p.Reg.SP += 2
	return nn
}

func (p *CPU) Push16(v uint16) {
	p.Reg.SP -= 2
	p.Memory.Write(p.Reg.SP, uint8(v))
	p.Memory.Write(p.Reg.SP+1, uint8(v>>8))
}

func (p *CPU) GetB() uint8 {
	return p.Reg.B
}

func (p *CPU) SetB(v uint8) {
	p.Reg.B = v
}

func (p *CPU) GetC() uint8 {
	return p.Reg.C
}

func (p *CPU) SetC(v uint8) {
	p.Reg.C = v
}

func (p *CPU) GetD() uint8 {
	return p.Reg.D
}

func (p *CPU) SetD(v uint8) {
	p.Reg.D = v
}

func (p *CPU) GetE() uint8 {
	return p.Reg.E
}

func (p *CPU) SetE(v uint8) {
	p.Reg.E = v
}

func (p *CPU) GetH() uint8 {
	return p.Reg.H
}

func (p *CPU) SetH(v uint8) {
	p.Reg.H = v
}

func (p *CPU) GetL() uint8 {
	return p.Reg.L
}

func (p *CPU) SetL(v uint8) {
	p.Reg.L = v
}

func (p *CPU) GetA() uint8 {
	return p.Reg.A
}

func (p *CPU) SetA(v uint8) {
	p.Reg.A = v
}

func (p *CPU) GetF() uint8 {
	return p.Reg.F
}

func (p *CPU) SetF(v uint8) {
	p.Reg.F = v&0xd7 | 0x02
}

func (p *CPU) GetBC() uint16 {
	return uint16(p.Reg.C) + (uint16(p.Reg.B) << 8)
}

func (p *CPU) SetBC(v uint16) {
	p.Reg.C = uint8(v)
	p.Reg.B = uint8(v >> 8)
}

func (p *CPU) GetDE() uint16 {
	return uint16(p.Reg.E) + (uint16(p.Reg.D) << 8)
}

func (p *CPU) SetDE(v uint16) {
	p.Reg.E = uint8(v)
	p.Reg.D = uint8(v >> 8)
}

func (p *CPU) GetHL() uint16 {
	return uint16(p.Reg.L) + (uint16(p.Reg.H) << 8)
}

func (p *CPU) SetHL(v uint16) {
	p.Reg.L = uint8(v)
	p.Reg.H = uint8(v >> 8)
}

func (p *CPU) GetAF() uint16 {
	return uint16(p.Reg.F) + (uint16(p.Reg.A) << 8)
}

func (p *CPU) SetAF(v uint16) {
	p.Reg.F = uint8(v&0xd7 | 0x02)
	p.Reg.A = uint8(v >> 8)
}

func (p *CPU) GetSP() uint16 {
	return p.Reg.SP
}

func (p *CPU) SetSP(v uint16) {
	p.Reg.SP = v
}

func (p *CPU) GetPC() uint16 {
	return p.Reg.PC
}

func (p *CPU) SetPC(v uint16) {
	p.Reg.PC = v
}

func (p *CPU) GetR8(n uint8) uint8 {
	switch n {
	case 0:
		return p.GetB()
	case 1:
		return p.GetC()
	case 2:
		return p.GetD()
	case 3:
		return p.GetE()
	case 4:
		return p.GetH()
	case 5:
		return p.GetL()
	case 6:
		return p.Memory.Read(p.GetHL())
	case 7:
		return p.GetA()
	}
	return 0
}

func (p *CPU) SetR8(n uint8, v uint8) {
	switch n {
	case 0:
		p.SetB(v)
	case 1:
		p.SetC(v)
	case 2:
		p.SetD(v)
	case 3:
		p.SetE(v)
	case 4:
		p.SetH(v)
	case 5:
		p.SetL(v)
	case 6:
		p.Memory.Write(p.GetHL(), v)
	case 7:
		p.SetA(v)
	}
}

func (p *CPU) GetR16(n uint8) uint16 {
	switch n {
	case 0:
		return p.GetBC()
	case 1:
		return p.GetDE()
	case 2:
		return p.GetHL()
	case 3:
		return p.GetSP()
	}
	return 0
}

func (p *CPU) SetR16(n uint8, v uint16) {
	switch n {
	case 0:
		p.SetBC(v)
	case 1:
		p.SetDE(v)
	case 2:
		p.SetHL(v)
	case 3:
		p.SetSP(v)
	}
}

func (p *CPU) GetR16S(n uint8) uint16 {
	switch n {
	case 0:
		return p.GetBC()
	case 1:
		return p.GetDE()
	case 2:
		return p.GetHL()
	case 3:
		return p.GetAF()
	}
	return 0
}

func (p *CPU) SetR16S(n uint8, v uint16) {
	switch n {
	case 0:
		p.SetBC(v)
	case 1:
		p.SetDE(v)
	case 2:
		p.SetHL(v)
	case 3:
		p.SetAF(v)
	}
}

func (p *CPU) CC(n uint8) bool {
	switch n {
	case 0: // NZ
		return p.Reg.F&0x40 == 0
	case 1: // Z
		return p.Reg.F&0x40 != 0
	case 2: // NC
		return p.Reg.F&0x01 == 0
	case 3: // C
		return p.Reg.F&0x01 != 0
	case 4: // PO
		return p.Reg.F&0x04 == 0
	case 5: // PE
		return p.Reg.F&0x04 != 0
	case 6: // P
		return p.Reg.F&0x80 == 0
	case 7: // M
		return p.Reg.F&0x80 != 0
	}
	return false
}

func (p *CPU) DisableInterrupt() {
	p.CanInterrupt = false
}

func (p *CPU) EnableInterrupt() {
	p.CanInterrupt = true
}

func op_adc(v1 uint8, v2 uint8, v3 uint8) (uint8, uint8, uint8) {
	value1 := uint16(v1)
	value2 := uint16(v2)
	value3 := uint16(v3)
	result := value1 + value2 + value3
	xorval := result ^ value1 ^ value2
	// fmt.Println(value1, value2, value3, result, xorval)
	hcarry := xorval & 16
	carry := (xorval >> 8) & 1
	return uint8(result), uint8(carry), uint8(hcarry)
}

func op_sbc(v1 uint8, v2 uint8, v3 uint8) (uint8, uint8, uint8) {
	value1 := uint16(v1)
	value2 := uint16(v2)
	value3 := uint16(v3)
	result := value1 - value2 - value3
	xorval := result ^ value1 ^ value2
	// fmt.Println(value1, value2, value3, result, xorval)
	hcarry := (xorval & 16) ^ 16 // ???
	carry := (xorval >> 8) & 1
	return uint8(result), uint8(carry), uint8(hcarry)
}

func (p *CPU) Op8(n uint8, v uint8) {
	switch n {
	case 0:
		// ADD R0 / ADI n
		r, c, hc := op_adc(p.GetA(), v, 0)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 1:
		// ADC R0 / ACI n
		r, c, hc := op_adc(p.GetA(), v, p.GetF()&1)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 2:
		// SUB R0 / SUI n
		r, c, hc := op_sbc(p.GetA(), v, 0)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 3:
		// SBB R0 / SBI n
		r, c, hc := op_sbc(p.GetA(), v, p.GetF()&1)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 4:
		// ANA R0 / ANI n
		a := p.GetA()
		r := a & v
		p.SetA(r)
		hcarry := ((a | v) & 0x08) << 1
		p.SetF(FlagTable[r] | hcarry)
	case 5:
		// XRA R0 / XRI n
		var r uint8 = p.GetA() ^ v
		p.SetA(r)
		p.SetF(FlagTable[r])
	case 6:
		// ORA R0 / ORI n
		var r uint8 = p.GetA() | v
		p.SetA(r)
		p.SetF(FlagTable[r])
	case 7:
		// CMP R0 / CPI n
		r, c, hc := op_sbc(p.GetA(), v, 0)
		p.SetF(FlagTable[r] | c | hc)
	}
}

func (p *CPU) Op() {
	// https://pastraiser.com/cpu/i8080/i8080_opcodes.html
	// http://dunfield.classiccmp.org/r/8080.txt
	op := p.Fetch8()
	p.FetchCount += OpcodeTable[op].bytes
	op6 := op & 0x3f
	switch op >> 6 {
	case 0:
		switch op6 & 15 {
		case 0, 8:
			// NOP
		case 1:
			// LXI
			nn := p.Fetch16()
			p.SetR16(op6>>4, nn)
		case 2:
			switch op6 >> 4 {
			case 0, 1:
				// STAX
				p.Memory.Write(p.GetR16(op6>>4), p.GetA())
			case 2:
				// SHLD
				nn := p.Fetch16()
				p.Memory.Write(nn, p.GetL())
				p.Memory.Write(nn+1, p.GetH())
			case 3:
				// STA
				nn := p.Fetch16()
				p.Memory.Write(nn, p.GetA())
			}
		case 3:
			// INX
			p.SetR16(op6>>4, p.GetR16(op6>>4)+1)
		case 4, 12:
			// INR
			v := p.GetR8(op6>>3) + 1
			p.SetR8(op6>>3, v)
			var hcarry uint8 = 0
			if (v & 15) == 0 {
				hcarry = 0x10
			}
			p.SetF(FlagTable[v] | hcarry | (p.GetF() & 1))
		case 5, 13:
			// DCR
			v := p.GetR8(op6>>3) - 1
			p.SetR8(op6>>3, v)
			// ???
			var hcarry uint8 = 0x10
			if (v & 15) == 15 {
				hcarry = 0x00
			}
			p.SetF(FlagTable[v] | hcarry | (p.GetF() & 1))
		case 6, 14:
			// MVI
			v := p.Fetch8()
			p.SetR8(op6>>3, v)
		case 7, 15:
			switch op6 >> 3 {
			case 0:
				// RLC
				a := p.GetA()
				a = (a << 1) | ((a >> 7) & 1)
				p.SetA(a)
				p.SetF(p.GetF()&0xfe | (a & 1))
			case 1:
				// RRC
				a := p.GetA()
				p.SetF(p.GetF()&0xfe | (a & 1))
				a = (a >> 1) | ((a << 7) & 0x80)
				p.SetA(a)
			case 2:
				// RAL
				a := p.GetA()
				v := (a << 1) | (p.GetF() & 1)
				p.SetA(v)
				p.SetF(p.GetF()&0xfe | ((a >> 7) & 1))
			case 3:
				// RAR
				a := p.GetA()
				v := (a >> 1) | ((p.GetF() << 7) & 0x80)
				p.SetA(v)
				p.SetF(p.GetF()&0xfe | (a & 1))
			case 4:
				// DAA
				a := p.GetA()
				f := p.GetF()
				var adj uint8 = 0
				if (f&0x10) != 0 || (a&0x0f) >= 0x0a {
					adj += 0x06
				}
				if (f&0x01) != 0 || (a&0xf0) >= 0xa0 || ((a&0xf0) >= 0x90 && (a&0x0f) >= 0x0a) {
					adj += 0x60
					f |= 0x01
				}
				r, _, hc := op_adc(a, adj, 0)
				p.SetA(r)
				p.SetF(FlagTable[r] | (f & 0x01) | hc)
			case 5:
				// CMA
				p.SetA(p.GetA() ^ 0xff)
			case 6:
				// STC
				p.SetF(p.GetF() | 1)
			case 7:
				// CMC
				p.SetF(p.GetF() ^ 1)
			}
		case 9:
			// DAD
			v := p.GetR16(op6 >> 4)
			s := p.GetHL()
			d := s + v
			c := uint8(0)
			if d < s {
				c = 1
			}
			p.SetHL(d)
			p.SetF(p.GetF()&0xfe | c)
		case 10:
			switch op6 >> 4 {
			case 0, 1:
				// LDAX
				p.SetA(p.Memory.Read(p.GetR16(op6 >> 4)))
			case 2:
				// LHLD
				nn := p.Fetch16()
				p.SetL(p.Memory.Read(nn))
				p.SetH(p.Memory.Read(nn + 1))
			case 3:
				// LDA
				nn := p.Fetch16()
				p.SetA(p.Memory.Read(nn))
			}
		case 11:
			// DCX
			p.SetR16(op6>>4, p.GetR16(op6>>4)-1)
		}
	case 1:
		// MOV R1,R0 / HLT
		if op == 0x76 {
			p.Halted = true
			p.Reg.PC = p.Reg.PC - 1
		} else {
			p.SetR8(op6>>3, p.GetR8(op6&7))
		}
	case 2:
		// OP R0
		p.Op8(op6>>3, p.GetR8(op6&7))
	case 3:
		switch op6 & 15 {
		case 0, 8:
			// RET CC
			if p.CC(op6 >> 3) {
				p.SetPC(p.Pop16())
			}
		case 1:
			// POP
			p.SetR16S(op6>>4, p.Pop16())
		case 2, 10:
			// JP CC
			nn := p.Fetch16()
			if p.CC(op6 >> 3) {
				p.SetPC(nn)
			}
		case 3:
			switch op6 >> 4 {
			case 0:
				// JMP
				nn := p.Fetch16()
				p.SetPC(nn)
			case 1:
				// OUT
				a := p.GetA()
				nn := uint16(p.Fetch8()) + (uint16(a) << 8)
				p.IOPort.Write(nn, a)
			case 2:
				// XTHL
				v := p.Read16(p.GetSP())
				p.Write16(p.GetSP(), p.GetHL())
				p.SetHL(v)
			case 3:
				// DI
				p.DisableInterrupt()
			}
		case 4, 12:
			// CALL CC
			nn := p.Fetch16()
			if p.CC(op6 >> 3) {
				p.Push16(p.GetPC())
				p.SetPC(nn)
			}
		case 5:
			// PUSH
			p.Push16(p.GetR16S(op6 >> 4))
		case 6, 14:
			// OPI
			v := p.Fetch8()
			p.Op8(op6>>3, v)
		case 7, 15:
			// RST
			p.Push16(p.GetPC())
			p.SetPC(uint16(op6 & 0x38))
		case 9:
			switch op6 >> 4 {
			case 0:
				// RET
				p.SetPC(p.Pop16())
			case 2:
				// PCHL
				p.SetPC(p.GetHL())
			case 3:
				// SPHL
				p.SetSP(p.GetHL())
			}
		case 11:
			switch op6 >> 4 {
			case 1:
				// IN
				a := p.GetA()
				nn := uint16(p.Fetch8()) + (uint16(a) << 8)
				p.SetA(p.IOPort.Read(nn))
			case 2:
				// XCHG
				v := p.GetDE()
				p.SetDE(p.GetHL())
				p.SetHL(v)
			case 3:
				// EI
				p.EnableInterrupt()
			}
		case 13:
			if op == 0xcd {
				// CALL
				nn := p.Fetch16()
				p.Push16(p.GetPC())
				p.SetPC(nn)
			}
		}
	}
	p.Instructions += 1
}

func (p *CPU) Reset() {
	p.CanInterrupt = true
	p.Halted = false
	p.FetchCount = 0
	p.Instructions = 0
	p.Reg.PC = 0
	p.Reg.SP = 0
	p.Reg.A = 0
	p.Reg.F = 0
	p.Reg.B = 0
	p.Reg.C = 0
	p.Reg.D = 0
	p.Reg.E = 0
	p.Reg.H = 0
	p.Reg.L = 0
}

func (p *CPU) Interrupt(addr uint16) {
	if p.CanInterrupt {
		p.CanInterrupt = false
		p.Push16(p.GetPC())
		p.SetPC(addr)
	}
}

func (p *CPU) Exec(n int, debugFuc func(*CPU)) int {
	startFetchCount := p.FetchCount
	for i := 0; i < n; i++ {
		if debugFuc != nil {
			debugFuc(p)
		}
		p.Op()
		if p.Halted {
			break
		}
	}
	return p.FetchCount - startFetchCount
}

func (p *CPU) Status() string {
	hl := uint16(p.Reg.L) | (uint16(p.Reg.H) << 8)
	m := p.Memory.Read(hl)
	// op := p.Memory.Read(pc)
	f := p.Reg.F
	sf := [6]byte{'S', 'Z', 'A', 'P', 'C', 'I'}
	if f&0x80 == 0 {
		sf[0] = '-'
	}
	if f&0x40 == 0 {
		sf[1] = '-'
	}
	if f&0x10 == 0 {
		sf[2] = '-'
	}
	if f&0x04 == 0 {
		sf[3] = '-'
	}
	if f&0x01 == 0 {
		sf[4] = '-'
	}
	if !p.CanInterrupt {
		sf[5] = '-'
	}

	opr := p.Memory.Read(p.Reg.PC)
	opcodes := OpcodeTable[opr]
	mnemonic := opcodes.mnemonic
	opbuff := [3]uint8{opr, 0, 0}
	for i := 1; i < opcodes.bytes; i++ {
		opbuff[i] += p.Memory.Read(p.Reg.PC + uint16(i))
	}
	ops := ""
	for i := 0; i < opcodes.bytes; i++ {
		ops += fmt.Sprintf("%02X ", opbuff[i])
	}
	if opcodes.bytes == 2 {
		mnemonic = strings.Replace(mnemonic, "n", fmt.Sprintf("$%02X", opbuff[1]), 1)

	} else if opcodes.bytes == 3 {
		mnemonic = strings.Replace(mnemonic, "nn", fmt.Sprintf("$%02X%02X", opbuff[2], opbuff[1]), 1)
	}

	stacktop0 := p.Memory.Read(p.Reg.SP)
	stacktop1 := p.Memory.Read(p.Reg.SP + 1)
	s := fmt.Sprintf("%08d A=%02X F=%02X %s BC=%02X%02X DE=%02X%02X HL=%04X [%02X] SP=%04X [%02X%02X] PC=%04X | %-9s| %s", p.Instructions, p.Reg.A, f, sf, p.Reg.B, p.Reg.C, p.Reg.D, p.Reg.E, hl, m, p.Reg.SP, stacktop1, stacktop0, p.Reg.PC, ops, mnemonic)
	return s
}
