# with open("2024/day04/example.txt", "r") as f:
with open("2024/day04/input.txt", "r") as f:
    lines = f.readlines()

def find_4(x, y, w, h):
    res = []

    if (x + 3) < w:
        res.append([
            (x, y),
            (x + 1, y),
            (x + 2, y),
            (x + 3, y),
        ])
    if (x - 3) > -1:
        res.append([
            (x, y),
            (x - 1, y),
            (x - 2, y),
            (x - 3, y),
        ])
    if (y + 3) < h:
        res.append([
            (x, y),
            (x, y + 1),
            (x, y + 2),
            (x, y + 3),
        ])
    if (y - 3) > -1:
        res.append([
            (x, y),
            (x, y - 1),
            (x, y - 2),
            (x, y - 3),
        ])
    if (x + 3) < w and (y + 3) < h:
        res.append([
            (x, y),
            (x + 1, y + 1),
            (x + 2, y + 2),
            (x + 3, y + 3),
        ])
    if (x - 3) > -1 and (y + 3) < h:
        res.append([
            (x, y),
            (x - 1, y + 1),
            (x - 2, y + 2),
            (x - 3, y + 3),
        ])
    if (x + 3) < w and (y - 3) > -1:
        res.append([
            (x, y),
            (x + 1, y - 1),
            (x + 2, y - 2),
            (x + 3, y - 3),
        ])
    if (x - 3) > -1 and (y - 3) > -1:
        res.append([
            (x, y),
            (x - 1, y - 1),
            (x - 2, y - 2),
            (x - 3, y - 3),
        ])
    return res

def find_4_around(x, y, w, h):
    res = []
    if x < 1 or x >= w - 1 or y < 1 or y >= h - 1:
        return res
    res.append((x - 1, y - 1)) # top left
    res.append((x + 1, y + 1)) # bottom right
    res.append((x - 1, y + 1)) # bottom left
    res.append((x + 1, y - 1)) # top right
    return res

for line in lines:
    line = line.strip("\n")
sum = 0
h = len(lines)
w = len(lines[0])
for y in range(h):
    for x in range(w):
        possible = find_4(x, y, w, h)
        for letters in possible:
            if lines[letters[0][1]][letters[0][0]] != "X":
                continue
            if lines[letters[1][1]][letters[1][0]] != "M":
                continue
            if lines[letters[2][1]][letters[2][0]] != "A":
                continue
            if lines[letters[3][1]][letters[3][0]] != "S":
                continue
            sum += 1

print(sum)

sum = 0

for y in range(h):
    for x in range(w):
        # print(lines[y][x])
        if lines[y][x] != "A":
            continue
        possible = find_4_around(x, y, w, h)
        if len(possible) == 0:
            continue
        top_left = lines[possible[0][1]][possible[0][0]]
        bottom_right = lines[possible[1][1]][possible[1][0]]
        bottom_left = lines[possible[2][1]][possible[2][0]]
        top_right = lines[possible[3][1]][possible[3][0]]
        # print(top_left, top_right, bottom_left, bottom_right)
        if top_left == bottom_right:
            continue
        if bottom_left == top_right:
            continue
        m_count = 0
        s_count = 0
        for l in [top_left, top_right, bottom_left, bottom_right]:
            if l == "M":
                m_count += 1
            if l == "S":
                s_count += 1
        if m_count == 2 and s_count == 2:
            sum += 1
print(sum)
