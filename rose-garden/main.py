m, n = map(int, input().split())

garden = [0]*m

for _ in range(0, n):
    row = list(input())

    for j, rose in enumerate(row):
        if rose == "W":
            garden[j] += 1

print("".join("B" if count % 2 == 0 else "F" for count in garden))
