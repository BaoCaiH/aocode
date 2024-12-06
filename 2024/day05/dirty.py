# with open("2024/day05/example.txt", "r") as f:
with open("2024/day05/input.txt", "r") as f:
    lines = f.readlines()

ordering = []
ordering_map = dict()
instructions = []
done = False
sum = 0
middle = 0
incorrects = []
for line in lines:
    if line == "\n":
        done = True
        continue
    if not done:
        ordering.append(line.strip("\n"))
        left = int(line.strip("\n").split("|")[0])
        right = int(line.strip("\n").split("|")[1])
        if ordering_map.get(left, None) is None:
            ordering_map[left] = set()
        ordering_map[left].add(right)
    else:
        instructions.append(line.strip("\n"))
        instr = [int(n) for n in line.strip("\n").split(",")]
        for i in range(len(instr)):
            if i == len(instr) - 1:
                sum += 1
                # print(instr[(len(instr) // 2)])
                middle += instr[(len(instr) // 2)]
                continue
            if not ordering_map.get(instr[i], set()).issuperset(instr[i+1:]):
                incorrects.append(instr)
                break

# print(ordering_map)
# print(instructions)

print(sum)
print(middle)

# print(incorrects)
def merge_sort(lst, left, right):
    if left < right:
        middle = left + ((right - left) // 2)

        merge_sort(lst, left, middle)
        merge_sort(lst, middle + 1, right)
        merge(lst, left, middle, right)

def merge(lst, left, middle, right):
    left_part = middle - left + 1
    right_part = right - middle

    left_lst = [0] * (left_part)
    right_lst = [0] * (right_part)

    # Copy data to temp arrays L[] and R[]
    for i in range(0, left_part):
        left_lst[i] = lst[left + i]

    for i in range(0, right_part):
        right_lst[i] = lst[middle + 1 + i]

    i = 0
    j = 0
    k = left

    while i < left_part and j < right_part:
        # print(left_lst[i], right_lst[j], ordering_map.get(right_lst[j], {}))
        if {left_lst[i]}.issubset(ordering_map.get(right_lst[j], {})):
            lst[k] = left_lst[i]
            i += 1
        else:
            lst[k] = right_lst[j]
            j += 1
        k += 1

    while i < left_part:
        lst[k] = left_lst[i]
        i += 1
        k += 1

    while j < right_part:
        lst[k] = right_lst[j]
        j += 1
        k += 1

middle = 0
for inc in incorrects:
    merge_sort(inc, 0, len(inc) - 1)
    middle += inc[(len(inc) // 2)]

print(middle)

# merge_sort(incorrects[0], 0, len(incorrects[0]) - 1)
# print(incorrects[0])
