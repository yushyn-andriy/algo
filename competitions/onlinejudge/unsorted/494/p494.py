import sys
import string


stdin = sys.stdin


if __name__ == '__main__':
    ascii = set(string.ascii_letters)
    for row in stdin:
        row = row.strip()
        
        counter = 0
        flag = False
        for ch in row:
            ascii_sym = ch in ascii
            if flag and not ascii_sym:
                flag = False
            elif not flag and ascii_sym:
                flag = True
                counter += 1
        print(counter)
