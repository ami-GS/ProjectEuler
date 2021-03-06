import itertools
import time

N = 28123
ansList = [False] * (N+1)
abundants = []

s = time.time()
for i in range(12,N-12+1):
    j = 1
    divisors = []
    while j <= i**0.5:
        if i % j == 0:
            divisors.append(j)
            divisors.append(i/j)
        if i/j - j <= 1:
            break
        j += 1
    if i < sum(list(set(divisors)))-i:
        abundants.append(i)

for twoAbun in itertools.combinations_with_replacement(abundants, 2):
    tmpSum = sum(twoAbun)
    if tmpSum > N:
        continue
    else:
        ansList[tmpSum] = True

ans = [i for i in range(1,len(ansList)) if not ansList[i]]
print sum(ans)
print time.time() - s

