import itertools
import time

s = time.time()
count = 0
ans = 0
for i in itertools.permutations([0,1,2,3,4,5,6,7,8,9], 10):
    if i[3]%2 != 0 or (i[2]+i[3]+i[4])%3 != 0 or i[5]%5 != 0 or abs(i[4]*10+i[5]-i[6]*2)%7 != 0 or abs(i[5]-i[6]+i[7])%11 != 0:
        continue
    if (100*i[6]+10*i[7]+i[8])%13 != 0 or (100*i[7]+10*i[8]+i[9])%17 != 0:
        continue
    count += 1
    ans += int("".join(map(str, i)))#probably faster than below
#    ans += sum([i[idx]*10**(9-idx) for idx in range(10)])
print ans, count, time.time() - s
