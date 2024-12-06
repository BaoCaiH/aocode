# with open("2024/day06/example.txt", "r") as f:
with open("2024/day06/input.txt", "r") as f:
    lines = f.readlines()

obstacles = []
guard = None
direction = (0, -1)
for i, line in enumerate(lines):
    line = line.strip("\n")
    for j, thing in enumerate(line):
        if thing == "#":
            obstacles.append((j, i))
        elif thing == "^":
            guard = (j, i)

print(obstacles)
print(guard)
print(i + 1, j + 1)

paths = set()

def go_until_obstacle(obstacles, pos, dir, w, h):
    path = set()
    cnt = 10000
    while not (
        pos in obstacles
        or pos[0] < 0
        or pos[0] >= h
        or pos[1] < 0
        or pos[1] >= w
        or cnt < 0
    ):
        cnt -= 1
        path.add(pos)
        pos = (pos[0] + dir[0], pos[1] + dir[1])

    if pos in obstacles:
        pos = (pos[0] - dir[0], pos[1] - dir[1])
        if dir == (0, 1):
            dir = (-1, 0)
        elif dir == (1, 0):
            dir = (0, 1)
        elif dir == (0, -1):
            dir = (1, 0)
        else:
            dir = (0, -1)
    else:
        pos = (pos[0] - dir[0], pos[1] - dir[1])
        dir = (0, 0)
    return path, pos, dir

initial = guard
initial_direction = direction

while (direction != (0, 0)):
    path, guard, direction = go_until_obstacle(obstacles, guard, direction, j + 1, i + 1)
    # print(path)
    paths = paths.union(path)

# print(paths)
print(len(paths))

sum = 0
for place in paths:
    # print(place)
    new_obstacle = obstacles + [place]
    guard = initial
    direction = initial_direction
    turning = {guard: set(direction)}
    # if sum % 10 == 0:
    #     print(sum)
    while (direction != (0, 0)):
        _, guard, direction = go_until_obstacle(new_obstacle, guard, direction, j + 1, i + 1)
        # print(guard, direction)
        # paths = paths.union(path)
        if guard not in turning:
            turning[guard] = set(direction)
        elif direction not in turning[guard]:
            turning[guard].add(direction)
        else:
            sum += 1
            break

print(sum)
