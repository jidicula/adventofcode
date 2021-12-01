def input_read():
    input_list = []
    with open("input.txt", "r") as f:
        for line in f:
            input_list.append(line)
    return input_list


boarding_pass_list = input_read()
seat_ids = set()

for boarding_pass in boarding_pass_list:
    row = boarding_pass[:7]
    column = boarding_pass[7:-1]

    row_binary = row.replace("B", "1").replace("F", "0")
    row_dec = int(row_binary, base=2)

    column_binary = column.replace("R", "1").replace("L", "0")
    column_dec = int(column_binary, base=2)
    seat_id = row_dec * 8 + column_dec
    seat_ids.add(seat_id)

print(max(seat_ids))

# part 2
expected_ids = set(8 * row + col for row in range(127) for col in range(3))

seat_candidates = list(expected_ids - seat_ids)

for seat in seat_candidates:
    if seat + 1 in seat_ids and seat - 1 in seat_ids:
        print(seat)
        break
