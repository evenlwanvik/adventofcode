number_decoder = {
    "zero": "0",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
}

with open('./input.txt') as f:
    lines = f.readlines()
    sum = 0
    for l in lines:
        numbers = []
        # Start at first character
        for i, c_i in enumerate(l):
            current_string = c_i
            try:
                int(c_i)
                numbers.append(c_i) # Check if character is a number
                continue
            except ValueError:
                pass
            # Check encoder after appending each following character
            for j, c_j in enumerate(l[i + 1:]):
                current_string += c_j
                if current_string in number_decoder:
                    numbers.append(number_decoder[current_string])
                    break
        sum += int(numbers[0] + numbers[-1])

print(sum)