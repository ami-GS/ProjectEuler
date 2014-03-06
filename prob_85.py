import time
s = time.time()
N = 2000000

def Comb(n, x=2):
    def fact(x):
        if x == 1:
            return x
        else:
            return fact(x-1) * x
    
    return reduce(lambda x,y: x*y, [i for i in range(n-x+1, n+1)]) / fact(x)

o1 = 0
for i in range(1, 10000):
    if Comb(i,2) >= N:
        o1 = i
        break

ans = []
grid = []
closest = 10000
area = 0
for i in range(o1, 0, -1):
    for j in range(1, o1):
        calc = abs(N - Comb(i)*Comb(j))
        if Comb(i,2) * Comb(j,2) >= N:
            ans.append(abs(N - Comb(i,2)*Comb(j-1,2)))
            grid.append((i-1)*(j-2))
            break

print min(ans), grid[ans.index(min(ans))]
print time.time() - s
