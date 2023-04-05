import sys



stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        v, t = list(map(int, row.split()))
        if t == 0:
            print(0)
        else:
            a = float(v) / t
            s = int(a * t ** 2 + v * t)
            print(s)
    
        
