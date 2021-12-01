import itertools
from collections import Counter

pw_db = []
valid = 0

with open("input.txt", "r") as f:
    for line in f:
        pw_db.append(line)

for record in pw_db:
    words = record.split()
    num_range = words[0].split("-")
    min_freq = int(num_range[0])
    max_freq = int(num_range[1])
    char = words[1][0]
    pw = words[2]
    char_count = Counter(pw)[char]
    if char_count >= min_freq and char_count <= max_freq:
        valid += 1

print(valid)

from operator import xor

toboggan_valid = 0
for record in pw_db:
    words = record.split()
    num_range = words[0].split("-")
    pos1 = int(num_range[0]) - 1
    pos2 = int(num_range[1]) - 1
    char = words[1][0]
    pw = words[2]
    if xor(pw[pos1] == char, pw[pos2] == char):
        toboggan_valid += 1
print(toboggan_valid)
