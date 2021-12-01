import re


def input_read():
    input_str = ""
    with open("input.txt", "r") as f:
        for line in f:
            input_str += line
    return input_str


passports_str = input_read()

passports_raw_list = passports_str.split("\n\n")
valid_passports_list = []

valid = 0
for passport in passports_raw_list:
    passport_fields = re.split(r"\s+", passport)
    passport_fields_dict = {}
    for field in passport_fields:
        if field != "":
            field_split = field.split(":")
            k = field_split[0]
            v = field_split[1]
            passport_fields_dict[k] = v
    if len(passport_fields_dict.keys()) == 7 and "cid" not in passport_fields_dict:
        valid += 1
        valid_passports_list.append(passport_fields_dict)
    if len(passport_fields_dict.keys()) == 8:
        valid += 1
        valid_passports_list.append(passport_fields_dict)
print(valid)

true_valid = 0
for passport in valid_passports_list:
    birth_yr = False
    issue_yr = False
    exp_yr = False
    height = False
    hair_clr = False
    eye_clr = False
    passport_id = False
    if 1920 <= int(passport["byr"]) <= 2002:
        birth_yr = True
    else:
        continue
    if 2010 <= int(passport["iyr"]) <= 2020:
        issue_yr = True
    else:
        continue
    if 2020 <= int(passport["eyr"]) <= 2030:
        exp_yr = True
    else:
        continue
    if re.match(r"\d+cm", passport["hgt"]):
        if 150 <= int(re.split(r"cm", passport["hgt"])[0]) <= 193:
            height = True
        else:
            continue
    if re.match(r"\d+in", passport["hgt"]):
        if 59 <= int(re.split(r"in", passport["hgt"])[0]) <= 76:
            height = True
        else:
            continue
    if re.match(r"\#[0-9a-f]{6}", passport["hcl"]):
        hair_clr = True
    else:
        continue
    if passport["ecl"] in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth"):
        eye_clr = True
    else:
        continue
    if re.match(r"\d{9}", passport["pid"]):
        passport_id = True
    else:
        continue
    if (
        birth_yr
        and issue_yr
        and exp_yr
        and height
        and hair_clr
        and eye_clr
        and passport_id
    ):
        true_valid += 1
print(true_valid)
