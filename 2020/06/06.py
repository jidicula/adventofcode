def input_read():
    input_str = ""
    with open("input.txt", "r") as f:
        for line in f:
            input_str += line
    return input_str


groups = input_read().split("\n\n")
yes_count = 0

for group in groups:
    yes_count += len(set(group.replace("\n", "")))
print(yes_count)

# pt 2
import functools

correct_yes_count = 0
for group in groups:
    group_members = group.strip().split("\n")
    member_sets = list(map(set, group_members))
    group_yes_set = functools.reduce(lambda x, y: x & y, member_sets)
    correct_yes_count += len(group_yes_set)
print(correct_yes_count)
