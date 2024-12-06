# with open("2024/day02/example.txt", "r") as f:
with open("2024/day02/input.txt", "r") as f:
    lines = f.readlines()

levels = [line.replace("\n", "").split(" ") for line in lines]

sum = 0
def validate(level):
    size = len(level)
    for i, _ in enumerate(level):
        if i == 0 and int(level[i]) < int(level[i + 1]):
            increasing = True
        elif i == 0:
            increasing = False
        if i == size - 1:
            return True, -1
        if increasing and int(level[i + 1]) - int(level[i]) <= 3 and int(level[i + 1]) > int(level[i]):
            continue
        if not increasing and int(level[i]) - int(level[i + 1]) <= 3 and int(level[i]) > int(level[i + 1]):
            continue
        return False, i

for level in levels:
    increasing = None
    size = len(level)
    res, i = validate(level)
    if res:
        sum += 1
    else:
        res0, _ = validate(level[:i-1] + level[i:])
        res1, _ = validate(level[:i] + level[i+1:])
        res2, _ = validate(level[:i+1] + level[i+2:])
    
        if res0 or res1 or res2:
            sum += 1
        else:
            print(level)

print(sum)
