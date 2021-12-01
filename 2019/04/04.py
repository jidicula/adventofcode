import re
from itertools import combinations_with_replacement
min_pw = 156218
max_pw = 652527
numbers = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]

foo = ["112233", "123444", "111122", "144444"]

doubled = []
for num in foo:  # range(min_pw, max_pw):
    num_str = str(num)
    if re.match(r".*(\d)\1{1}.*", num_str) and (re.match(
            r".*(\d)\1{2,}.*", num_str) is None):
        print(num_str, "match")
        doubled.append(num_str)
    else:
        print(num_str, "no match")
print(doubled)
combos = set(combinations_with_replacement(numbers, 6))
combos_str = set(map("".join, combos))
# print(combos_str)
passwords = set(doubled).intersection(combos_str)
print(passwords)

print(len(passwords))
