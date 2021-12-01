def input_read(filename: str):
    input_list = []
    with open(filename, "r") as f:
        for line in f:
            input_list.append(line)
    return input_list
