def input_read():
    input_list = []
    with open("input.txt", "r") as f:
        for line in f:
            input_list.append(line)
    return input_list


slope_map = input_read()
line_length = len(slope_map[0])


def count_trees(instructions: (int, int)) -> int:
    slope_map = input_read()
    tree_counter = 0
    right = instructions[0]
    down = instructions[1]
    line_length = len(slope_map[0])
    col = 0
    for line in range(0, len(slope_map), down):
        if slope_map[line][col % (line_length - 1)] == "#":
            tree_counter += 1
        col += right
    return tree_counter


print(count_trees((3, 1)))


paths = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
product = 1
for instruction in paths:
    trees = count_trees(instruction)
    print(trees)
    product *= trees
print(product)
