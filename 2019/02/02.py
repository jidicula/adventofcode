def run_smoke(memory):
    for address in range(0, len(memory), 4):
        instruction_pointer = memory[address]
        if instruction_pointer == 99:
            break
        loc1 = memory[address + 1]
        loc2 = memory[address + 2]
        loc3 = memory[address + 3]
        arg1 = memory[loc1]
        arg2 = memory[loc2]
        if instruction_pointer == 1:
            memory[loc3] = arg1 + arg2
        elif instruction_pointer == 2:
            # print(loc1, arg1)
            # print(loc2, arg2)
            memory[loc3] = arg1 * arg2
    return (memory[0])


for noun in range(99):
    for verb in range(99):
        memory = [
            1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1,
            5, 19, 23, 1, 23, 5, 27, 2, 27, 10, 31, 1, 5, 31, 35, 2, 35, 6, 39,
            1, 6, 39, 43, 2, 13, 43, 47, 2, 9, 47, 51, 1, 6, 51, 55, 1, 55, 9,
            59, 2, 6, 59, 63, 1, 5, 63, 67, 2, 67, 13, 71, 1, 9, 71, 75, 1, 75,
            9, 79, 2, 79, 10, 83, 1, 6, 83, 87, 1, 5, 87, 91, 1, 6, 91, 95, 1,
            95, 13, 99, 1, 10, 99, 103, 2, 6, 103, 107, 1, 107, 5, 111, 1, 111,
            13, 115, 1, 115, 13, 119, 1, 13, 119, 123, 2, 123, 13, 127, 1, 127,
            6, 131, 1, 131, 9, 135, 1, 5, 135, 139, 2, 139, 6, 143, 2, 6, 143,
            147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 9, 155, 0, 99, 2, 14, 0, 0
        ]
        # print(noun,verb)
        memory[1] = noun
        memory[2] = verb
        output = run_smoke(memory)
        if output == 19690720:
            print(100 * noun + verb)
            break
