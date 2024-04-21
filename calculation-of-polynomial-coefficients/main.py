m, n = map(int, input().split())

coefficients = list(map(int, input().split()))


def polynomialPower(coefficients, m, n):
    result = coefficients.copy()

    tempResult = []

    for _ in range(n - 1):
        tempResult = [0] * (len(result) + m)

        for i in range(len(result)):
            for j in range(len(coefficients)):
                tempResult[i + j] += result[i] * coefficients[j]

        result = tempResult

    return result


print(" ".join(str(res) for res in polynomialPower(coefficients, m, n)))
