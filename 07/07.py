import re
def input_read(filename: str):
    input_list = []
    with open(filename, "r") as f:
        for line in f:
            input_list.append(line)
    return input_list

rules = input_read("test_case.txt")
rules = input_read("input.txt")

def rule_parse(rule: str):
    split_rule = rule.split("contain")
    key = split_rule[0].split("bags")[0].strip()
    val = split_rule[1].strip().split(",")
    val = list(map(lambda x: re.split(r"bags?", x)[0], val))
    val = list(map(lambda x: re.split(r"\d+\s", x)[-1], val))
    val = tuple(map(lambda x: x.strip(), val))

    return key, val

rules_dict = dict(map(rule_parse, rules))

# print(rules_dict)
shiny_gold_holders = set()

def holdsgold(key):
    contained = rules_dict[key]
    for colour in contained:
        if colour in shiny_gold_holders:
            shiny_gold_holders.add(key)
            continue
        elif colour == "shiny gold":
            shiny_gold_holders.add(key)
            continue
        elif colour == "no other":
            continue
        else:
            holdsgold(colour)
    return

for key in rules_dict.keys():
    holdsgold(key)

for key in rules_dict.keys():
    holdsgold(key)
# print(shiny_gold_holders)
print(len(shiny_gold_holders))


# print(rules_dict.values())

# pt 2
rules = input_read("test_case_2.txt")
rules = input_read("input.txt")

def rule_parse_pt2(rule: str):
    split_rule = rule.split("contain")
    key = split_rule[0].split("bags")[0].strip()
    val = split_rule[1].strip().split(",")
    val = list(map(lambda x: re.split(r"bags?", x)[0], val))
    # val = list(map(lambda x: re.split(r"\d+\s", x)[-1], val))
    val = tuple(map(lambda x: x.strip(), val))

    return key, val

rules_dict = dict(map(rule_parse_pt2, rules))
# print(rules_dict)

def cost(colour: str):
    contents = rules_dict[colour]
    n = 0
    for bag in contents:
        if bag != "no other":   # recursive case
            n_bags = int(re.match(r"\d+", bag)[0])
            bag_colour = re.split(r"\d+\s", bag)[1]
            n += n_bags + n_bags * cost(bag_colour)
    return n


print(cost("shiny gold"))
