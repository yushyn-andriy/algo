import sys
import bisect


stdin = sys.stdin


def abs_diff_sum(array, A):
    return sum([abs(x - A) for x in array])




if __name__ == '__main__':
    for row in stdin:
        n = int(row.strip())
        X = []
        for _ in range(n):
            X.append(int(stdin.readline().strip()))

        X = sorted(X)
        if n % 2 == 1:
            i = n // 2
        else:
            i = n // 2 - 1 
        
        r1 = X[i]
        r2 = abs(bisect.bisect_left(X, r1) - bisect.bisect_right(X, r1))
        if n % 2 == 1:
            r3 = 1
        else:
            r3 = 1 if X[i] == X[i+1] else X[i+1] - X[i] + 1
            if r3 != 1:
                otherMedian = X[i]
                r2 += abs(bisect.bisect_left(X, otherMedian) - bisect.bisect_right(X, otherMedian))
        
        print(r1,r2,r3)
