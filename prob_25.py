import math

def func():
    N = 1000
    f_1, f_2 = 1, 1
    t = 0
    idx = 2
    while t < N-1:
        f_2, f_1 = f_1, f_1+f_2
        t = math.log10(f_1)
        idx += 1
    print(idx, t, f_1)
        

if __name__ == "__main__":
    func()