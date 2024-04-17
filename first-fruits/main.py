a = list(map(int, input().split()))

idealWatermelon = 0

for i in a:
    if i >= 80:
        idealWatermelon += 1

print("Mamma mia!" if idealWatermelon >=
      3 else "Mamma mia!!" if idealWatermelon >= 1 else "Mamma mia!!!")
