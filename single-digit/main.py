n = list(input())


def shrinkTheNumber(n):
    if len(n) == 1:
        return n[0]
    return shrinkTheNumber(str(sum(int(digit) for digit in n)))


print(shrinkTheNumber(n))
