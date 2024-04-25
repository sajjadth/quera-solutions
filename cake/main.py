import sys


def calculateArea(w, n):
    area = 0
    for _ in range(n):
        wi, li = map(int, sys.stdin.readline().rstrip().split())
        area += wi * li
    return area // w


if __name__ == "__main__":
    w = int(sys.stdin.readline().rstrip())
    n = int(sys.stdin.readline().rstrip())

    result = calculateArea(w, n)

    print(result)
