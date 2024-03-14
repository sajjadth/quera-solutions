calendar = {0: 31, 1: 31, 2: 31, 3: 31, 4: 31, 5: 31,
            6: 30, 7: 30, 8: 30, 9: 30, 10: 30, 11: 29}
date = list(map(int, input().split("/")))
month, day = date[1], date[2]

output = (calendar[month-1] - day)+1

month += 1

while month <= 12:
    output += calendar[month-1]
    month += 1

print(output)
