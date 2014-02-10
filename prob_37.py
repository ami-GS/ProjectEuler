import math
import time

s = time.time()

N = 1000000
idx = [1]*N

idx[0], idx[1] = 0, 0

for i in range(2, int(N**0.5)):
    if idx[i]:
        for j in range(i*2, N, i):
            idx[j] = 0

primes = [i for i in range(N) if idx[i] == 1]

count = 0
ans = 0
for prime in primes[4:]:
    pDigit = int(math.log(prime, 10))+1
    for i in range(1,pDigit):
        #print str(prime)[i:], str(prime)[:-i]
        if not idx[int(str(prime)[i:])]:
            break
        if not idx[int(str(prime)[:-i])]:
            break
        if i == pDigit-1:
            count += 1
            ans +=  prime
#            print prime
    if count == 11:
        break

print ans, count, time.time()-s
