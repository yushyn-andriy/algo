import sys


stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        if not row:
            print()
            continue
        

        s = 0
        for ch in row.strip():
            if ch.isdigit():
                s += int(ch)
            elif ch == 'b':
                print(' ' * s, end='')
                s = 0
            elif ch.isalpha() or ch == '*':
                print(ch * s, end='')
                s = 0
            elif ch == '!':
                print()
                s = 0
        
        print()