import sys



def get_answer(x):
    return str(int(abs( ( ( ( ( (x * 567) / 9) + 7492 ) * 235 )  / 47)  - 498)) )[-2]


if __name__ == '__main__':
    n = int(sys.stdin.readline())
    for _ in range(n):
        x = int(sys.stdin.readline())
        print(get_answer(x))
