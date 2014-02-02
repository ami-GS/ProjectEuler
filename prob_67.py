import time

with open("./triangle.txt") as f:
    triangle = "".join(f.readlines())

s = time.time()
triList = [map(int, li.split(" ")) for li in triangle.split("\n") if len(li) >= 2]

i = 0
while i < len(triList)-1:
    tmpSum = 0
    for j in range(i+1):
        if tmpSum < triList[i][j] + triList[i+1][j]:
            triList[i+1][j] += triList[i][j]
        else:
            triList[i+1][j] = tmpSum
        tmpSum = triList[i+1][j+1] + triList[i][j]        
    triList[i+1][i+1] = tmpSum
    i += 1

print max(triList[len(triList)-1])
print time.time() - s
