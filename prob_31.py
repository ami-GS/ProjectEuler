coins = [1,2,5,10,20,50,100,200]

def fn(rest, coins):
    if rest == 0 or len(coins) == 1:
        return 1
    else:
        largest = coins[-1]
        m = rest/largest
        total = 0
        for i in range(m+1):
            total += fn(rest - largest*i, coins[:-1])
        return total

print fn(200, coins)
