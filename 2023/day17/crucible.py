import heapq
from sys import argv

grid = []
with open(argv[1], "r") as f:
    lines = f.readlines()
    for line in lines:
        grid.append([int(c) for c in line if c != "\n"])

width = len(grid[0])
height = len(grid)

offsets = {">": (0, 1), "<": (0, -1), "^": (-1, 0), "v": (1, 0)}
opposites = {">": "<", "<": ">", "^": "v", "v": "^"}

traverse = []
heapq.heappush(traverse, [0, 0, 0, "."])
passed = set()
while traverse:
    l, r, c, d = heapq.heappop(traverse)

    if r == height - 1 and c == width - 1:
        print(l)
        break
    
    if (r, c, d) in passed:
        continue
    
    passed.add((r, c, d))
    for dir, (dr, dc) in offsets.items():
        if dir == opposites.get(d[0], "."):
            continue
        if dir == d[0] and len(d) >= 3:
            continue
        newR = r + dr
        newC = c + dc
        if newR < 0 or newR >= height or newC < 0 or newC >= width:
            continue
        newD = dir if dir != d[0] else d + dir
        heapq.heappush(traverse, [l + grid[newR][newC], newR, newC, newD])

traverse = []
heapq.heappush(traverse, [0, 0, 0, "."])
passed = set()
while traverse:
    l, r, c, d = heapq.heappop(traverse)

    if r == height - 1 and c == width - 1 and len(d) >= 4:
        print(l)
        break
    
    if (r, c, d) in passed:
        continue
    
    passed.add((r, c, d))
    for dir, (dr, dc) in offsets.items():
        if dir == opposites.get(d[0], "."):
            continue
        if dir == d[0] and len(d) >= 10:
            continue
        if dir != d[0] and len(d) < 4 and d != ".":
            continue
        newR = r + dr
        newC = c + dc
        if newR < 0 or newR >= height or newC < 0 or newC >= width:
            continue
        newD = dir if dir != d[0] else d + dir
        heapq.heappush(traverse, [l + grid[newR][newC], newR, newC, newD])

