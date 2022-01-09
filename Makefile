all: go8080 sample.bin

run: go8080 sample.bin
	./go8080 sample.bin

cpm: go8080
	./go8080 cpm13.dsk

cpm: go8080
	./go8080 cpm13.dsk

clean:
	rm go8080 sample.bin

go8080: i8080/cpu.go main.go
	go build -o go8080

sample.bin: sample.asm
	pasmo -8 -d sample.asm sample.bin sample.sym
