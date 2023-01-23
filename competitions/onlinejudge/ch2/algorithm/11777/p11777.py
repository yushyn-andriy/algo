import sys
import bisect





grades = list('FDCBA')
percents = [60, 70, 80, 90]


stdin = sys.stdin


def avmax2(l):
    max1 = max(l)
    l.pop(l.index(max1))
    max2 = max(l)
    return (max1 + max2) / 2


def get_grade(scores, grades=grades, percents=percents):
    score = sum(scores[:4]) + avmax2(scores[4:])
    return grades[bisect.bisect(percents, score)]




if __name__ == '__main__':
    c = int(stdin.readline())
    for i in range(1, c+1):
        scores = [int(x) for x in stdin.readline().strip().split()]
        print(f'Case {i}: {get_grade(scores)}')