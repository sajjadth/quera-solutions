n = int(input())
m = int(input())

if n < m:
    n, m = m, n


print(m*n if m <= 2 else ((m+n-1)*2)-2)
