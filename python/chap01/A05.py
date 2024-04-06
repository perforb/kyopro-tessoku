n, k = map(int, input().split())
counter = 0
for red in range(1, n + 1):
    for blue in range(1, n + 1):
        white = k - (red + blue)
        if 1 <= white <= n:
            counter += 1
print(counter)
