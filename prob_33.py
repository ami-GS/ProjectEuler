
def gcd(a, b):
    if a % b == 0:
        return b
    else:
        gcd(b, a%b)

ans = []
for i in range(11,100):
    if i % 10 == 0 or i % 11 == 0:
        continue
    ten, one = i/10, i%10

    for j in range(1, one):
        if float(ten*10+j) / i == float(j) / one:
            ans.append([j,one])
            print (ten*10+j), "/", i

    for j in range(1, ten):
        if float(j*10+one) / i == float(j) / ten:
            ans.append([j,ten])
            print (j*10+one), "/", i

    for j in range(1,ten):
        if float(j*10+ten) / i == float(j) / one:
            ans.append([j,one])
            print (j*10+ten), "/", i

    if ten >= one:
        for j in range(1,10):
            if float(one*10+j) / i == float(j) / ten:
                ans.append([j,ten])
                print (one*10+j), "/", i

print reduce(lambda a,b: a*b, [an[1] for an in ans]) / gcd(reduce(lambda a,b: a*b, [an[1] for an in ans]),
                                                           reduce(lambda a,b: a*b, [an[0] for an in ans]))
