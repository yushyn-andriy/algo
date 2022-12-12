import sys



if __name__ == '__main__':
    t = int(sys.stdin.readline())
    amount = 0
    for _ in range(t):
        args = sys.stdin.readline().strip().split()
        op = args[0]
        if op == 'donate':
            amount += int(args[1])
        elif op == 'report':
            print(amount)
