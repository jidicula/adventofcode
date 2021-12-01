def input_read(filename: str):
    with open(filename, "r") as f:
        input_list = [line.strip() for line in f]
    return input_list


input_list = input_read("input.txt")

int_input_list = [int(x) for x in input_list]

count = 0
prev = int_input_list[0]
for i in int_input_list[1:]:
    if i > prev:
        count += 1
    prev = i
print(count)

count = 0
previous_window_sum = 0
for i in range(len(int_input_list) - len(int_input_list) % 3):
    window = int_input_list[i : i + 3]
    # print(window)
    window_sum = sum(window)
    if window_sum > previous_window_sum:
        count += 1
    previous_window_sum = window_sum
print(count - 1)
