import sys


"""
+ - pos result        S = 1, S = 4, S = 78
- - neg result        S = S35
* - failed experiment S = 9S4
? - not completed     S = 190S
"""

94949435

POS, NEG, FAIL, NOT_COMPL = list('+-*?') 
m = {
    '1': POS,
    '4': POS,
    '78': POS,

    '135': NEG,
    '435': NEG,
    '7835': NEG,
    
    '914': FAIL,
    '944': FAIL,
    '9784': FAIL,

    '1901': NOT_COMPL,
    '1904': NOT_COMPL,
    '19078': NOT_COMPL,
}

def decrypt(s):
    r = m.get(s, None)
    if r:
        print(r)

    elif s[-1] == '5' and s[-2] == '3':
        print('-')
    elif s[0] == '9' and s[-1] == '4':
        print('*')
    elif s[0] == '1' and s[1] == '9' and s[2] == '0':
        print("?")
    else:
        print("+")


if __name__ == '__main__':
    n = int(sys.stdin.readline())
    for _ in range(n):
        enc = sys.stdin.readline().strip()
        decrypt(enc)
