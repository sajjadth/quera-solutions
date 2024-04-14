firstMatch = list(map(int, input().split()))
secondMatch = list(map(int, input().split()))

teamOne = firstMatch[0] + secondMatch[1]
teamTwo = firstMatch[1] + secondMatch[0]


if teamOne == teamTwo:
    if secondMatch[1] > firstMatch[1]:
        print("Persepolis")
    elif firstMatch[1] > secondMatch[1]:
        print("Esteghlal")
    else:
        print("Penalty")
elif teamOne > teamTwo:
    print("Persepolis")
else:
    print("Esteghlal")
