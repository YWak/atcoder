import sys
sys.setrecursionlimit(10**6)

n = int(input())
a = list(map(int, input().split()))
c = [0] * 4

for p in a:
    c[p] += 1

def id(a, b, c):
    return (a * 301 + b) * 301 + c

memo = [-1] * 301 * 301 * 301
flag = [0] * 301 * 301 * 301

memo[id(0, 0, 0)] = 0.0
flag[id(0, 0, 0)] = 1

def dp(a, b, c):
    idx = id(a, b, c)
    if flag[idx] == 0:
        flag[idx] = 1
        p = (a + b + c) / n
        ans = 1 / p
        if a > 0:
            ans += dp(a-1, b, c) * a / n / p
        if b > 0:
            ans += dp(a+1, b-1, c) * b / n / p
        if c > 0:
            ans += dp(a, b+1, c-1) * c / n / p
        memo[idx] = ans

    return memo[idx]

print(dp(c[1], c[2], c[3]))
