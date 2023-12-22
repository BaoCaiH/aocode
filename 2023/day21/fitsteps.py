from sys import argv

grid = []
start = (-1, -1)
rocks = set()
with open(argv[1], "r") as f:
    lines = f.readlines()
    for r, line in enumerate(lines):
        grid.append(line.strip("\n").replace("S", "."))
        for c, char in enumerate(line):
            if char == "#":
                rocks.add((r, c))
            elif char == "S":
                start = (r, c)

# for line in grid:
#     print(line)
#
# print(rocks)
# print(start)
# print(len(grid))

def adjacent(coor, grid, rocks):
    out = []
    r, c = coor
    height = len(grid)
    width = len(grid[0])
    for x, y in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        nr = r + x
        nc = c + y
        if nr >= 0 and nr < height and nc >= 0 and nc < width and (nr, nc) not in rocks:
            out.append((nr, nc))
    return out

# pos = [start]
def walk(start, grid, rocks, steps):
    seen = set()
    pos = set()
    pos.add(start)
    tmp_pos = []
    save = len(seen)
    for i in range(steps + 1):
        add = i % 2 == steps % 2
        for p in pos:
            if add and p not in seen:
                seen.add(p)
            out = adjacent(p, grid, rocks)
            tmp_pos += out
        pos = set(tmp_pos)
        tmp_pos = []
        if len(seen) == save and add:
            break
        else:
            save = len(seen)
    return len(seen)

print(walk(start, grid, rocks, 64))

# Max steps is (half grid) + (grid width * 202300)
# It's a square
grid_size = len(grid)
assert (grid_size // 2 + grid_size * 202300) == 26501365
width = 26501365 // len(grid) - 1
reached_with_even = walk(start, grid, rocks, 6400)
reached_with_odd = walk(start, grid, rocks, 6401)

odd_grids = (width // 2 * 2 + 1) ** 2
even_grids = ((width + 1) // 2 * 2) ** 2

corner_boxes = (
    walk((grid_size - 1, start[1]), grid, rocks, grid_size - 1)
    + walk((0, start[1]), grid, rocks, grid_size - 1)
    + walk((start[0], grid_size - 1), grid, rocks, grid_size - 1)
    + walk((start[0], 0), grid, rocks, grid_size - 1)
)

left_over = (
    # Tiny
    (width + 1) * (
        walk((0, 0), grid, rocks, grid_size - 1 - grid_size // 2 - 1)
        + walk((0, grid_size - 1), grid, rocks, grid_size - 1 - grid_size // 2 - 1)
        + walk((grid_size - 1, 0), grid, rocks, grid_size - 1 - grid_size // 2 - 1)
        + walk((grid_size - 1, grid_size - 1), grid, rocks, grid_size - 1 - grid_size // 2 - 1)
    )
    # Big bois
    + width * (
        walk((0, 0), grid, rocks, grid_size * 2 - 1 - grid_size // 2 - 1)
        + walk((0, grid_size - 1), grid, rocks, grid_size * 2 - 1 - grid_size // 2 - 1)
        + walk((grid_size - 1, 0), grid, rocks, grid_size * 2 - 1 - grid_size // 2 - 1)
        + walk((grid_size - 1, grid_size - 1), grid, rocks, grid_size * 2 - 1 - grid_size // 2 - 1)
    )
)

print(reached_with_even * even_grids + reached_with_odd * odd_grids + corner_boxes + left_over)
# print(reached_with_even * even_grids)
# print(reached_with_odd * odd_grids)
# print(left_over)


# def adjacent_out(coor, grid, rocks):
#     out = []
#     r, c = coor
#     height = len(grid)
#     width = len(grid[0])
#     for x, y in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
#         nr = r + x
#         nc = c + y
#         if (nr % height, nc % width) not in rocks:
#             out.append((nr, nc))
#     return out
#
# pos = set()
# pos.add(start)
# tmp_pos = []
# for i in range(50):
#     for p in pos:
#         out = adjacent_out(p, grid, rocks)
#         tmp_pos += out
#     # pos = tmp_pos
#     pos = set(tmp_pos)
#     tmp_pos = []
#     print(i+1, len(pos))




