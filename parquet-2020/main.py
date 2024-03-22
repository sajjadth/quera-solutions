def findRoomDimensions(black, yellow):
    total = black + yellow

    for width in range(3, int(total**0.5) + 1):
        if total % width == 0:
            length = total // width
            if (width + length) * 2 - 4 == black:
                return str(length) + " " + str(width)


inp = list(map(int, input().split()))
b, y, = inp[0], inp[1]

print(findRoomDimensions(b, y))
