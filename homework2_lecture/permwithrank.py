import math

def identity_perm(n):
    return list(range(1, n + 1))

def unrank(n, r, pi):
    result = []
    for i in range(n):
        fact = math.factorial(n - 1 - i)
        index = r // fact
        result.append(pi[index])
        pi.pop(index)
        r %= fact
    return result

for i in range(24):
    p = identity_perm(4)  # Reset p for each rank
    print(f"Rank {i}: {unrank(4, i, p)}")
