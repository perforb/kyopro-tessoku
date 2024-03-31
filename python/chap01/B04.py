n = input()
bits = [*n]
answer = 0
for i, j in zip(range(len(bits)), range(len(bits) - 1, -1, -1)):
    power_of_two = 1 << j
    answer += power_of_two * int(bits[i])
print(answer)
