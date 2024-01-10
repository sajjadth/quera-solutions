s, f, l, x = map(int, input().split(" "))

print("exam did not started!" if x < s else (
    "exam finished!" if x >= f else (l if f - x >= l else f-x)))
