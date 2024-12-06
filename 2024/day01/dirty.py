with open("2024/day01/input.txt", "r") as f:
    lines = f.readlines()

left = []
right = []
for line in lines:
    p = line.split("   ")
    left.append(int(p[0]))
    right.append(int(p[1]))
left1 = sorted(left)
right1 = sorted(right)

sum = 0
for i, n in enumerate(left):
    sum += abs(n - right[i])

print(sum)

leftD = {}
rightD = {}
for i in range(len(left)):
    if left[i] not in leftD.keys():
        leftD[left[i]] = 0
    if right[i] not in rightD.keys():
        rightD[right[i]] = 0
    leftD[left[i]] += 1
    rightD[right[i]] += 1

sum = 0
for k, v in leftD.items():
    sum += k * v * rightD.get(k, 0)

print(sum)
