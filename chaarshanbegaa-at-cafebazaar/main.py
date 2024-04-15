import math

n = int(input())
distances = []


def calculatePoints(distances):
    points = 0
    for distance in distances:
        if distance <= 10:
            points += 10
        elif distance <= 30:
            points += 9
        elif distance <= 50:
            points += 8
        elif distance <= 70:
            points += 7
        elif distance <= 90:
            points += 6
        elif distance <= 110:
            points += 5
        elif distance <= 130:
            points += 4
        elif distance <= 150:
            points += 3
        elif distance <= 170:
            points += 2
        elif distance <= 190:
            points += 1
    return points


for i in range(0, n):
    x, y = list(map(int, input().split()))
    distance = math.sqrt(x**2 + y**2)
    distances.append(distance)

print(calculatePoints(distances))
