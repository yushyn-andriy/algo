import sys


stdin = sys.stdin



if __name__ == '__main__':
    n = int(stdin.readline().strip())

    for i in range(1, n+1):
        C, R = [int(x) for x in stdin.readline().strip().split()]
        diff = C - R 
        
        print(f'Case #{i}:', end='')
        if diff == 0:
            print(' 0')
            continue

        
        res = set()
        i = 1
        while i * i <= diff:
            if diff % i == 0:
                res.add(i)
                res.add(int(diff/i))

            i+=1
        
        res = sorted(res)
        for v in res:
            if v > R:
                print(f' {v}', end='')
        print()
    