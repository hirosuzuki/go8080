package i8080

import (
	"fmt"
	"os"
	"strconv"
)

var FlagTable []uint8 = []uint8{
	0x42, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02,
	0x06, 0x02, 0x02, 0x06, 0x02, 0x06, 0x06, 0x02, 0x02, 0x06, 0x06, 0x02, 0x06, 0x02, 0x02, 0x06,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86,
	0x82, 0x86, 0x86, 0x82, 0x86, 0x82, 0x82, 0x86, 0x86, 0x82, 0x82, 0x86, 0x82, 0x86, 0x86, 0x82,
}

type IO64K interface {
	Read(addr uint16) byte
	Write(addr uint16, value byte)
}

type CPU struct {
	Memory IO64K
	IOPort IO64K
	Reg    struct {
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

func (p *CPU) Reset() {
	fmt.Println("Reset")
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
	p.Reg.SP += 2
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
	p.Reg.F = v
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
	p.Reg.F = uint8(v)
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

func (p *CPU) CC(n int) bool {
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

func op_adc(v1 uint8, v2 uint8, v3 uint8) (uint8, uint8, uint8) {
	value1 := uint16(v1)
	value2 := uint16(v2)
	value3 := uint16(v3)
	result := value1 + value2 + value3
	xorval := result ^ value1 ^ value2
	fmt.Println(value1, value2, value3, result, xorval)
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
	fmt.Println(value1, value2, value3, result, xorval)
	hcarry := xorval & 16
	carry := (xorval >> 8) & 1
	return uint8(result), uint8(carry), uint8(hcarry)
}

func (p *CPU) Op8(n uint8, v uint8) {
	switch n {
	case 0x00:
		// ADD R0 / ADI n
		r, c, hc := op_adc(p.GetA(), v, 0)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 0x08:
		// ADC R0 / ACI n
		r, c, hc := op_adc(p.GetA(), v, p.GetF()&1)
		p.SetA(r)
		p.SetF(FlagTable[r] | c | hc)
	case 0x10:
		// SUB R0 / SUI n
		r, c, hc := op_sbc(p.GetA(), v, 0)
		p.SetA(r)
		p.SetF(FlagTable[r] | (c ^ 1) | (hc ^ 16))
	case 0x18:
		// SBB R0 / SBI n
		r, c, hc := op_sbc(p.GetA(), v, p.GetF()&1)
		p.SetA(r)
		p.SetF(FlagTable[r] | (c ^ 1) | (hc ^ 16))
	case 0x20:
		// ANA R0 / ANI n
		var r uint8 = p.GetA() & v
		p.SetA(r)
		p.SetF(FlagTable[r])
	case 0x28:
		// XRA R0 / XRI n
		var r uint8 = p.GetA() ^ v
		p.SetA(r)
		p.SetF(FlagTable[r])
	case 0x30:
		// ORA R0 / ORI n
		var r uint8 = p.GetA() | v
		p.SetA(r)
		p.SetF(FlagTable[r])
	case 0x38:
		// CMP R0 / CPI n
	}
}

func (p *CPU) Op() {
	// https://pastraiser.com/cpu/i8080/i8080_opcodes.html
	op := p.Fetch8()
	op0 := op & 0x07
	op1 := op & 0x38
	op2 := op & 0xc0
	switch op2 {
	case 0x00:
		// NOP
	case 0x40:
		// MOV R1,R0
		p.SetR8(op1, p.GetR8(op0))
	case 0x80:
		// OP R0
		p.Op8(op1, p.GetR8(op0))
	case 0xc0:
		switch op0 {
		case 0x00:
			// RET CC
		case 0x01:
			switch op1 {
			case 0x00:
				// POP B
				p.SetBC(p.Pop16())
			case 0x08:
				// RET
				p.SetPC(p.Pop16())
			case 0x10:
				// POP D
				p.SetDE(p.Pop16())
			case 0x20:
				// POP H
				p.SetHL(p.Pop16())
			case 0x28:
				// PCHL
				p.SetPC(p.GetHL())
			case 0x30:
				// POP PSW
				p.SetAF(p.Pop16())
			case 0x38:
				// SPHL
				v := p.Read16(p.GetSP())
				p.Write16(p.GetSP(), p.GetHL())
				p.SetHL(v)
			}
		case 0x02:
			// JP CC
		case 0x03:
		case 0x04:
			// CALL CC
		case 0x05:
			switch op1 {
			case 0x00:
				// PUSH B
				p.Push16(p.GetBC())
			case 0x08:
				// CALL
				nn := p.Fetch16()
				p.Push16(p.GetPC())
				p.SetPC(nn)
			case 0x10:
				// PUSH D
				p.Push16(p.GetDE())
			case 0x20:
				// PUSH H
				p.Push16(p.GetHL())
			case 0x30:
				// PUSH PSW
				p.Push16(p.GetAF())
			}
		case 0x06:
			// OPI
			v := p.Fetch8()
			p.Op8(op1, v)
		case 0x07:
			// RST
			p.Push16(p.GetPC())
			p.SetPC(uint16(op1))
		}
	}
	fmt.Println("Fetch", op)
}

func (p *CPU) Exec() {
	fmt.Println("Exec")
	p.Op()
}

func Add() {
	arg1, _ := strconv.Atoi(os.Args[1])
	arg2, _ := strconv.Atoi(os.Args[2])
	arg3, _ := strconv.Atoi(os.Args[3])
	r, c8, c4 := op_adc(uint8(arg1), uint8(arg2), uint8(arg3))
	fmt.Println("adc", arg1, arg2, arg3, "->", r, c8, c4)
	r, c8, c4 = op_sbc(uint8(arg1), uint8(arg2), uint8(arg3))
	fmt.Println("sbc", arg1, arg2, arg3, "->", r, c8, c4)
}
