a = list(map(int, input().split()))
b = list(map(int, input().split()))

counter = 0

for i in range(0, 8):
    if a[i] == b[i] and a[i] == 1:
        counter += 1

print(counter)
