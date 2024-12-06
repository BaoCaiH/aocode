import re

# with open("2024/day03/example.txt", "r") as f:
with open("2024/day03/input.txt", "r") as f:
    lines = f.readlines()

# print(lines)

sum = 0
for line in lines:
    matches = re.findall("mul\([0-9]+,[0-9]+\)", line)
    # print(matches)
    for match in matches:
        pair = match.lstrip("mul(").rstrip(")").split(",")
        left = int(pair[0])
        right = int(pair[1])
        sum += left * right

print(sum)

sum = 0
do = True
for line in lines:
    matches = re.findall("mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)", line)
    print(matches)
    for match in matches:
        if match == "do()":
            do = True
            continue
        if match == "don't()":
            do = False
            continue
        if do:
            pair = match.lstrip("mul(").rstrip(")").split(",")
            left = int(pair[0])
            right = int(pair[1])
            sum += left * right
print(sum)
# 119285584