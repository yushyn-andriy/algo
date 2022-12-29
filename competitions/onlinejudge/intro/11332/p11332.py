import sys


stdin = sys.stdin

# O(n^2) time | O(n^2) space - probably
def f(n):
    if len(n) == 1:
        return n
    
    s = 0
    for digit in n:
        s += int(digit)
    
    return f(str(s))


if __name__ == '__main__':
    while True:
        n = stdin.readline().strip()
        if n == '0':
            break
        print(f(n))
