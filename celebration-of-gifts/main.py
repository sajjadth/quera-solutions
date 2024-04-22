def main():
    n = int(input())
    names = {input(): 0 for _ in range(n)}

    for _ in range(n):
        giverName = input()
        balance, numGifts = map(int, input().split())

        if balance != 0 and numGifts != 0:
            giftAmount = balance // numGifts
            names[giverName] -= balance - (balance % numGifts)

            for _ in range(numGifts):
                receiver_name = input()
                names[receiver_name] += giftAmount
        else:
            for _ in range(numGifts):
                input()

    for name, balance in names.items():
        print(f"{name} {balance}")


if __name__ == "__main__":
    main()
