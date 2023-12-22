from sys import argv

class Brick:
    def __init__(self, c: str):
        self.c = [int(coor) for coor in c.replace("~", ",").split(",")]
        self.supports = set()
        self.supported_by = set()
    def stack(self, another):
        return (
            max(self.c[0], another.c[0]) <= min(self.c[3], another.c[3])
            and max(self.c[1], another.c[1]) <= min(self.c[4], another.c[4])
        )
    def __str__(self) -> str:
        return str(self.c) + str(self.supports) + str(self.supported_by)
    def __repr__(self) -> str:
        return self.__str__()

bricks = []
with open(argv[1], "r") as f:
    lines = f.readlines()
    for line in lines:
        bricks.append(Brick(line.strip("\n")))

bricks.sort(key=lambda x: x.c[2])
# print(bricks)
# Fall
for i, b1 in enumerate(bricks):
    height = 1
    for b2 in bricks[:i]:
        if b1.stack(b2):
            height = max(height, b2.c[5] + 1)
    b1.c[5], b1.c[2] = height + b1.c[5] - b1.c[2], height

bricks.sort(key=lambda x: x.c[2])

for i, b1 in enumerate(bricks):
    for j, b2 in enumerate(bricks[:i]):
        if b1.stack(b2) and b1.c[2] == b2.c[5] + 1:
            b1.supported_by.add(j)
            b2.supports.add(i)
total = 0
for brick in bricks:
    if all(len(bricks[i].supported_by) >= 2 for i in brick.supports):
        total += 1
print(total)

def cascade(brick):
    falling = [i for i in brick.supports if len(bricks[i].supported_by) == 1]
    fell = set(falling)
    while falling:
        below = falling.pop(0)
        for above in bricks[below].supports:
            if above not in fell and bricks[above].supported_by.issubset(fell):
                falling.append(above)
                fell.add(above)
    return len(fell)


total = 0
for brick in bricks:
    total += cascade(brick)
print(total)

