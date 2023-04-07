import sys


stdin = sys.stdin



if __name__ == '__main__':
    lines = stdin.readlines()
    
    M = len(max(lines, key=lambda x: len(x))) - 1
    L = []
    for line in lines:
        line = line.rstrip('\n')
        m = len(line)
        diff = M - m 
        if diff != 0:
            line += ' ' * diff
        L.append(line)


    for i in range(M):
        for j in range(len(L) - 1, -1, -1):
            row = L[j]
            print(row[i], end='')    
        print()

