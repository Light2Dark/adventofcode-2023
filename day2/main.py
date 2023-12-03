# Contributed from Discord (not mine!!) (just a reference for me)
# Represents a simple solution in python using regex

import re

with open('day2/input.txt','r') as f:
    lines = (f.readlines())

sum = 0

for l in lines:
    blue_numbers = map(lambda x: int(x), re.findall(r'(\d+)\sblue', l))
    green_numbers = map(lambda x: int(x), re.findall(r'(\d+)\sgreen', l))
    red_numbers = map(lambda x: int(x), re.findall(r'(\d+)\sred', l))

    sum = sum + max(red_numbers) * max(blue_numbers) * max(green_numbers)

print(sum)