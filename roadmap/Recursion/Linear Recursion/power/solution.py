def power(x, n):
    if n == 0:
        return 1
    return x * power(x, n - 1)

if __name__ == '__main__':
    x = int(input().strip())
    print(power(x, 4))

