import sys

def tex_quotes(s):
    result = ''
    first = True
    for i, ch in enumerate(s):
        if ch != '"':
            result += ch
            continue

        if ch == '"' and first:
            result += '``'
            first = False
        elif ch == '"' and not first:
            result += '\'\''
            first = True
    return result


if __name__ == '__main__':
    print(tex_quotes(sys.stdin.read()), end='')
