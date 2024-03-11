t = int(input())
for _ in range(t):
    card = input()
    counter = 0
    prev = '1'
    for c in card:
        if c != '1' and prev == '1':
            counter += 1
        prev = c
    print(counter)
    