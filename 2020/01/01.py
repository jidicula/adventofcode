import itertools

num_list = []

with open("input.txt", "r") as f:
    for line in f:
        num_list.append(int(line))

# num_list.sort()

# for n in num_list:
#     diff = 2020 - n
#     if diff in num_list:
#         print(n * diff)
#         break


# triplets = itertools.permutations(num_list, 3)

# for n in triplets:
#     if sum(n) == 2020:
#         print(n[0] * n[1] * n[2])
#         break


doublets = itertools.combinations(num_list, 2)
for n in doublets:
    if sum(n) == 2020:
        print(n[0] * n[1])
        break

triplets = itertools.combinations(num_list, 3)
for n in triplets:
    if sum(n) == 2020:
        print(n[0] * n[1] * n[2])
        break
