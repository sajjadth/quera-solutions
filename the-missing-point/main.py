def find_treasure(coordinates):
    x, y, z = 0, 0, 0

    for coord in coordinates:
        x ^= coord[0]
        y ^= coord[1]
        z ^= coord[2]

    return f"{x} {y} {z}"


coordinates = []
for _ in range(7):
    point = list(map(int, input().split()))
    coordinates.append(tuple(point))


print(find_treasure(coordinates))
