depth, path = map(str, input().split())

depth = int(depth)

res = 0

for i in path:
    res *= 2
    if i == 'L':
        res += 1

for i in range(depth-len(path)):
    res += 2**(depth-i)

print(res+1)
