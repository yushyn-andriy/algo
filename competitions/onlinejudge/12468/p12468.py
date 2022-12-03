import sys



stdin = sys.stdin



def get_min_presses(ch1, ch2):
    if ch2 == ch1:
        return 0
    
    p1, p2 = 0, 0
    if ch2 < ch1:
        p1 = ch1 - ch2
        p2 = 99 - ch1 + 1 + ch2
        return min([p1, p2])
    
    p1 = ch2 - ch1
    p2 = ch1 + 1 + 99 - ch2
    return min([p1, p2])


if __name__ == '__main__':
    while True:
        ch1, ch2 = [int(x) for x in stdin.readline().strip().split()]
        if ch1 == -1 and ch2 == -1:
            break
        print(get_min_presses(ch1, ch2))
