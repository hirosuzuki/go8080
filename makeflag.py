def calc(i: int) -> int:
    j = (i >> 4) ^ i
    j = (j >> 2) ^ j
    j = (j >> 1) ^ j
    return (i & 128) | (i == 0) * 64 | (j & 1 ^ 1) * 4 | 2

for i in range(0, 256, 16):
    for j in range(16):
        print("0x%02x" % calc(i + j), end=", ")
    print()

    