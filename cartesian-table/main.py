def print_grid(x, y):
    for i in range(y * 2 + 1):
        if i % 2 == 0:
            for j in range(x * 2 + 1):
                if j % 2 == 0:
                    print(" ", end="")
                else:
                    print("_", end="")
        else:
            for j in range(x * 2 + 1):
                if j % 2 == 0:
                    print("|", end="")
                else:
                    print(" ", end="")
        print()


x, y = map(int, input().split())

print_grid(y, x)
