import time

s = time.time()
N = 10000
pIdx = [1]*N

pIdx[0], pIdx[1] = 0, 0

for i in range(2, int(N**0.5)):
    if pIdx[i]:
        for j in range(i*2,N,i):
            pIdx[j] = 0

#make odd list except for primes
odds = [i for i in range(2,N) if pIdx[i] == 0 and i%2 != 0]

for odd in odds:
    ansFlag = True
    #below loop search for appropriate 'n'
    for i in range(1, int(((odd-3)/2)**0.5)+1):
        #this checks if it is prime
        if pIdx[odd-2*(i**2)] == 1:
            ansFlag = False
            break
    if ansFlag:
        print odd, odd-2*(i**2), i, time.time() - s
        break
