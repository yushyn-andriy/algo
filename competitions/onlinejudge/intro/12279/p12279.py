import sys



def get_balance(l):
    nils = 0
    not_nils = 0
    for ch in l:
        if ch == '0':
            nils += 1
        else:
            not_nils += 1
    return not_nils - nils


if __name__ == '__main__':
    c = 1 
    while True:
        n = int(sys.stdin.readline().strip())
        if n == 0:
            break
        line = sys.stdin.readline().strip().split()
        print(f'Case {c}: {get_balance(line)}')
        c += 1
