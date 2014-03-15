import sys,time
s = time.time()
N = int(sys.argv[1])
E = [1]*N; E[0] = 2

j = 1
for i in range(2,N,3):
    E[i] = 2*j
    j += 1

def func(n, den, num=1):
    if n > 0:
        return func(n-1, E[n-1]*den+num, den)
    else:
        return den, num

ans = func(N-1, E[N-1])
ans2 = 0
for i in range(len(str(ans[0]))):
    ans2 += int(str(ans[0])[i])
print(ans)
print(ans2, time.time() - s)
