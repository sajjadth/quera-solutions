import math


def isPrime(n):
    if n <= 1:
        return False
    if n <= 3:
        return True
    if n % 2 == 0 or n % 3 == 0:
        return False
    i = 5
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if n % i == 0:
            return False
    return True


def primeFactors(n):
    if isPrime(n):
        print(n)
        return

    i = 2
    isFirstFactor = True
    while i <= n:
        number = 0
        while n % i == 0:
            number += 1
            n //= i
        if number != 0:
            if isFirstFactor:
                isFirstFactor = False
            else:
                print("*", end='')
            if number != 1:
                print(f"{i}^{number}", end='')
            else:
                print(f"{i}", end='')
        i += 1


if __name__ == "__main__":
    n = int(input())
    primeFactors(n)
