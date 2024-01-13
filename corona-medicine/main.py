n = int(input())
k = int(input())
p = int(input())
q = int(input())

shekarestan = n - k
namakestan = p - q


print("Namakestan" if namakestan > shekarestan else (
    "Shekarestan" if shekarestan > namakestan else "Equal"))