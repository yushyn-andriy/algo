import sys



stdin = sys.stdin



if __name__ == '__main__':
    while True:
        sb = stdin.readline().strip()
        if sb == '0 0':
            break
        S, B = [int(x) for x in sb.split()]

        left, right = [0]*100005, [0]*100005
        for i in range(1, S + 2):
            left[i] = i - 1
            right[i] = i + 1
        right[S] = -1
        left[1] = -1
        for i in range(B):
            line = stdin.readline().strip()
            l, r = [int(x) for x in line.split()]
            left[right[r]] = left[l]
            right[left[l]] = right[r]
            if left[l] != -1:
                print(f'{left[l]}', end='')
            else:
                print('*', end='')
            if right[r] != -1:
                print(f' {right[r]}')
            else:
                print(' *')

        print('-')

