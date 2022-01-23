all: tester cpm  invaders AllRoms

clean:
	rm cpm invaders tester images/*.COM images/*.rom images/*.dsk cmd/tester/minicpm.bin

run-invaders: invaders images/SpaceInvaders.rom
	./invaders images/SpaceInvaders.rom

invaders: cmd/invaders/main.go
	go build -o invaders cmd/invaders/main.go

run-cpm: cpm images/cpm13.dsk
	./cpm images/cpm13.dsk

cpm: cmd/cpm/main.go
	go build -o cpm cmd/cpm/main.go

test: tester images/TST8080.COM images/CPUTEST.COM images/8080PRE.COM images/8080EXM.COM
	./tester images/TST8080.COM
	@echo
	./tester images/CPUTEST.COM
	@echo
	./tester images/8080PRE.COM
	@echo
	./tester images/8080EXM.COM
	@echo

tester: cmd/tester/main.go cmd/tester/minicpm.bin
	go build -o tester cmd/tester/main.go

cmd/tester/minicpm.bin: cmd/tester/minicpm.z80
	pasmo cmd/tester/minicpm.z80 cmd/tester/minicpm.bin

AllRoms: images/cpm13.dsk images/SpaceInvaders.rom images/8080EXM.COM images/8080PRE.COM images/CPUTEST.COM images/TST8080.COM

images/cpm13.dsk:
	wget -O images/cpm13.dsk https://github.com/udo-munk/z80pack/raw/master/cpmsim/disks/library/cpm13.dsk

images/SpaceInvaders.rom:
	wget -O images/SpaceInvaders.rom https://github.com/gergoerdi/clash-spaceinvaders/raw/master/image/SpaceInvaders.rom

images/8080EXM.COM:
	wget -O images/8080EXM.COM https://github.com/superzazu/8080/raw/master/cpu_tests/8080EXM.COM

images/8080PRE.COM:
	wget -O images/8080PRE.COM https://github.com/superzazu/8080/raw/master/cpu_tests/8080PRE.COM

images/CPUTEST.COM:
	wget -O images/CPUTEST.COM https://github.com/superzazu/8080/raw/master/cpu_tests/CPUTEST.COM

images/TST8080.COM:
	wget -O images/TST8080.COM https://github.com/superzazu/8080/raw/master/cpu_tests/TST8080.COM
