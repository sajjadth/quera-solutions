import math


def isPrime(number):
    if number < 2:
        return False
    if number == 2:
        return True
    if number % 2 == 0:
        return False
    for i in range(3, int(math.sqrt(number)) + 1, 2):
        if number % i == 0:
            return False
    return True


def hardPasswords(n):
    primeNumbers = [[2, 3, 5, 7]]
    numbers = [1, 3, 5, 7, 9]

    for i in range(1, n):
        primeNumbers.append([
            s * 10 + num for s in primeNumbers[i - 1] for num in numbers if isPrime(s * 10 + num)
        ])

    print(*primeNumbers[n - 1], sep='\n')


if __name__ == "__main__":
    n = int(input())
    hardPasswords(n)
