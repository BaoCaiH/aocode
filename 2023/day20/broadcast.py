from copy import deepcopy
import math
from sys import argv

class Module:
    def __init__(self, name, typ, children):
        self.name = name
        self.t = typ
        self.children = children or []
        self.state = False
        self.memory = {}
    def receive(self, name, signal):
        if self.t == "%" and signal == False:
            self.state = not self.state
            return self.state
        if self.t == "&":
            self.memory[name] = signal
            return not all(self.memory.values())
        if self.t == "b":
            return False
        return None


modules = {}
with open(argv[1], "r") as f:
    lines = f.readlines()
    for line in lines:
        line = line.strip("\n")
        splits = line.split(" -> ")
        m = splits[0]
        children = splits[1].split(", ")
        if m == "broadcaster":
            modules["broadcaster"] = Module(m, "b", children)
        else:
            modules[m[1:]] = Module(m[1:], m[0], children)
        for child in children:
            if child not in modules:
                modules[child] = Module(child, "", None)
modules["button"] = Module("button", "b", ["broadcaster"])

for module in modules.values():
    for child in module.children:
        if modules[child].t == "&" or modules[child].t == "":
            modules[child].memory[module.name] = False

# for k, v in modules.items():
#     print(k, v.t, v.state, v.memory, v.children)

storage = deepcopy(modules)

def push_button(stop=False, target=None):
    low = 0
    high = 0
    sequence = [(str("button"), str("broadcaster"), False)]
    while sequence:
        fr, to, signal = sequence.pop(0)
        if stop is True and signal and fr == target:
            return True
            
        # print(fr, signal, to)
        if signal:
            high += 1
        else:
            low += 1
        module = modules[to]
        if module.t == "":
            continue
        send = module.receive(fr, signal)

        if send is not None:
            for child in module.children:
                sequence.append((to, child, send))
        # print(sequence)
    if stop:
        return False
    return high, low


modules = dict(storage)
low = 0
high = 0
for _ in range(1000):
    result = push_button()
    l = 0
    h = 0
    if isinstance(result, tuple):
        l, h = result
    low += l
    high += h
print(f"Part 1: {low * high}")

# Too slow, must be a pattern, probably another LCM
# There are 4 things that feed into "zr" which is the only input for "rx"
last_input = list(storage["rx"].memory.keys())[0]
for module in storage.values():
    for child in module.children:
        if storage[child].name == last_input:
            storage[child].memory[module.name] = False
next_to_last = list(storage[last_input].memory.keys())
cnts = []
for t in next_to_last:
    cnt = 0
    found = False
    modules = deepcopy(storage)
    while not found:
        cnt += 1
        found = push_button(True, t)
    cnts.append(cnt)
print(f"Part 2: {math.lcm(*cnts)}")

