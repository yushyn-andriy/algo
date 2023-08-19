import sys


def find_prime_numbers(n):
    d = 2
    factors = []
    while d*d<=n:
        while (n % d) == 0:
            factors.append(d)
            n //= d
        d += 1
    if n > 1:
        factors.append(n)
    return factors


def main():
    max_n = 0
    res = []
    number = 0
    for i in range(1, 1000001):
        l = find_prime_numbers(i)
        n = len(l)
        if n>= max_n:
            max_n, res = n, l
            number = i
        print(i, find_prime_numbers(i))

    print(number, max_n, res)
    return 0


if __name__ == '__main__':
    sys.exit(main())
