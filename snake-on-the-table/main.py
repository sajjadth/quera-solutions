n, m = map(int, input().split())

result = []

for i in range(1, n + 1):
    row = list(map(str, range((i - 1) * m + 1, i * m + 1)))
    if i % 2 == 0:
        row = row[::-1]
    result.append(row)

for r in result:
    print(" ".join(r))
