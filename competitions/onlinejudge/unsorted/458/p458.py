import sys



stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        for encrypted in row[:len(row)-1]:
            decrypted = chr(ord(encrypted) - 7)
            print(decrypted, end='')
        print()

