n, m = map(int, input().split())

straightStreets = list(map(int, input().split()))
circularStreets = list(map(int, input().split()))

if (circularStreets.count(1) != 0 or circularStreets.count(0) != 0) and (straightStreets.count(0) != 0 and straightStreets.count(1) != 0):
    print("YES")
else:
    print("NO")
