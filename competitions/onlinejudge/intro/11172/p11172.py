import sys



def get_relational_operator(a, b):
    if a > b:
        return '>'
    if a < b:
        return '<'
    return '='


if __name__ == '__main__':
    cases = int(sys.stdin.readline())
    for _ in range(cases):
        a, b = [int(i) for i in sys.stdin.readline().split()]
        print(get_relational_operator(a, b))
