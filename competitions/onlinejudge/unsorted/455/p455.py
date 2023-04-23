import sys


stdin = sys.stdin


def eat():
    stdin.readline()


def period(s):
    for k in range(1, len(s) + 1):
        flag = True
        if len(s) % k != 0:
            continue
    
        prev = s[0:k]
        for i in range(k, len(s), k):
            curr = s[i:i+k]
            if prev != curr:
                flag = False
                break
            prev = s[i:i+k]         

        if flag:
            break

    return k if flag else len(s)



if __name__ == '__main__':
    n = int(sys.stdin.readline().strip())
    eat()
    
    for i in range(n):
        row = stdin.readline().strip()
        eat()
        print(period(row))
        if i == n - 1:
            continue
        print()
