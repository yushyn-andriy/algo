import sys

stdin = sys.stdin


def soundex(w):
    mm = {
        'BFPV': '1', 'CGJKQSXZ': '2', 'DT': '3',
        'L': '4', 'MN': '5', 'R': '6',
    }
    code = ''
    prev_ch = '-'
    for c in w:
        found = False
        for k in mm.keys():
            v = mm[k]
            if c in k:
                found = True
                if prev_ch == v:
                    break
                code += v
                prev_ch = v
        if not found:
            prev_ch = '-'


    return code


if __name__ == '__main__':
    while True:
        w = stdin.readline().strip()
        if not w:
            break
        print(soundex(w))

