def printPattern(n):
    for i in range(n):
        for j in range(n):
            if (i == 0 or i == n - 1 or j == 0 or j == n - 1
                    or i == j
                    or i+j == n-1
                    or j > n-1-i and i <= n//2
                    or j > i and i > n//2):
                print('#', end='')
            else:
                print(' ', end='')
        print()


n = int(input())

printPattern(n)
