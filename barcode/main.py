barcode = []

output = 0

for _ in range(9):
    inp = list(input())
    barcode.append(inp)

pattern = [
    [1, 1, 1, 2, 2, 2, 1, 1, 1],
    [1, 0, 1, 2, 2, 2, 1, 0, 1],
    [1, 1, 1, 2, 2, 2, 1, 1, 1],
    [2, 2, 2, 2, 2, 2, 2, 2, 2],
    [2, 2, 2, 2, 2, 2, 2, 2, 2],
    [2, 2, 2, 2, 2, 2, 2, 2, 2],
    [1, 1, 1, 2, 2, 2, 1, 1, 1],
    [1, 0, 1, 2, 2, 2, 1, 0, 1],
    [1, 1, 1, 2, 2, 2, 1, 1, 1]
]

for i in range(9):
    for j in range(9):
        barcode[i][j] = int(barcode[i][j])
        if pattern[i][j] != 2 and barcode[i][j] != 2 and barcode[i][j] != pattern[i][j]:
            print(0)
            exit()
        if pattern[i][j] == 2 and barcode[i][j] == 2:
            output += 1

print(2**output)
