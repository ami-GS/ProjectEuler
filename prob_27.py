import time

s = time.time()
N = 100000
idx = [1]*N

idx[0], idx[1] = 0, 0

for i in range(2, int(N**0.5)):
    if idx[i]:
        for j in range(i*2, N, i):
            idx[j] = 0

B = [i for i in range(1000) if idx[i] == 1] # calc prime

maxN = 0
ans = 0
for a in range(-999, 1000, 2):
    for b in B:
        i = 0
        while idx[i**2+a*i+b]:
            i += 1
        if maxN < i:
            maxN = i
            ans = a*b
            #print i,a,b

print maxN, ans
print time.time()-s
