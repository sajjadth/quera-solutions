import math

inp = list(map(int, input().split()))
phones = list(map(int, input().split()))

n, x, y = inp[0], inp[1], inp[2]

phones.sort()
least_charge = phones[0]
phones.pop(0)

total_charge_required = sum(math.ceil((100-p)/y) * x for p in phones)

print("YES" if total_charge_required <= least_charge else "NO")
