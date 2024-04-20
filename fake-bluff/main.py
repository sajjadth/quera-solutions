n = int(input())

output = 0
paper = ""

for i in range(0, 2*n):
    inp = input().replace(" ", "")

    if i % 2 == 0:
        paper = inp
    elif inp != paper:
        output += 1

print(output)
