n = int(input())
answer = ""
for i in range(9, -1, -1):
    x = 1 << i
    answer = answer + str(int(n / x) % 2)
print(answer)
