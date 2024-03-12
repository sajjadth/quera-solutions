inp = list(map(int, input().split()))

n, m = inp[0], inp[1]
counter = 0
prev = ""

for i in range(0, n):
    row = list(input())
    for r in row:
        if prev == "" or prev != r:
            prev = r
            counter += 1

print(counter)
