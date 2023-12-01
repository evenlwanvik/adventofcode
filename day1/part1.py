

with open('./input.txt') as f:
    lines = f.readlines()
    sum = 0
    for l in lines:
        first = None
        Last = None
        for c in l:
            if c.isdigit():
                if first is None:
                    first = c
                last = c
        sum += int(first+last)
        print(f"line: {l}, first: {first}, last: {last}")
print(sum)
