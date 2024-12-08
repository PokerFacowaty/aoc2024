from itertools import product
with open("input", encoding="utf-8") as f:
    fileLines = f.readlines()

valid = []
for line in fileLines:
    target_number = int(line.split(":")[0])
    part_numbers = [int(x) for x in line.split(": ")[1].split(" ")]
    if len(part_numbers) == 2:
        operator_options = [('*',), ('+',)]
    else:
        operator_options = list(product(["*", "+"],
                                        repeat=len(part_numbers) - 1))
    for options in operator_options:
        number_so_far = part_numbers[0]
        i = 0
        bust = False
        while i < len(options):
            if options[i] == "*":
                number_so_far = number_so_far * part_numbers[i + 1]
            else:
                number_so_far = number_so_far + part_numbers[i + 1]

            i += 1
            if number_so_far > target_number:
                bust = True
                break
        if bust:
            continue
        if number_so_far == target_number:
            valid.append(number_so_far)
            break

print(sum(valid))
