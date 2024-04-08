nextMonthPetrol = int(input())
fuelLeft = int(input())

fuelLeft += 60

print(nextMonthPetrol * 1500 if nextMonthPetrol <=
      fuelLeft else ((fuelLeft * 1500) + ((nextMonthPetrol - fuelLeft) * 3000)))
