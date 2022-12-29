import sys



"""
Hajj
Umrah
Hajj
Umrah
*
Case 1: Hajj-e-Akbar
Case 2: Hajj-e-Asghar
Case 3: Hajj-e-Akbar
Case 4: Hajj-e-Asghar

"""

H = {
    'Hajj': 'Hajj-e-Akbar',
    'Umrah': 'Hajj-e-Asghar',
}
if __name__ == '__main__':
    c = 1
    while True:
        line = sys.stdin.readline().strip()
        if line == '*':
            break

        print(f'Case {c}: {H.get(line, "")}')
        c += 1
