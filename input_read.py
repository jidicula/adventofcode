def input_read():
    input_list = []
    with open("input.txt", "r") as f:
        for line in f:
            input_list.append(line)
    return input_list
