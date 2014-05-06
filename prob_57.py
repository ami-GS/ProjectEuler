import sys
N = 1000
sys.setrecursionlimit(N+1)
numerator = 5
denominator = 2
ans = 0

def calc(numerator, denominator, n):
    global ans
    if n == 2:
        return denominator+numerator, denominator
    else:
        numerator = numerator + denominator*2
        if len(str(numerator+denominator)) > len(str(numerator)):
            ans += 1
        return calc(denominator, numerator, n-1)

calc(denominator, numerator, N) 
print ans
