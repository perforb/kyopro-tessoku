n, s = map(int, input().split())
a = list(map(int, input().split()))
answer = False
for i in range(0, 1 << n):
    sum = 0
    for j in range(1, n + 1):
        power_of_two = 1 << (j - 1)
        if int(i / power_of_two) % 2 == 1:
            sum += a[j - 1]
    if sum == s:
        answer = True
        break
print(answer)
