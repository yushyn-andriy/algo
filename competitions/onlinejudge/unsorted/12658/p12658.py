import sys


stdin = sys.stdin



STR_TO_NUMBER = {
    '.*...*...*...*...*..': 1, 
    '***...*.***.*...***.': 2, 
    '***...*.***...*.***.': 3,
}

def solve(matrix, n):
    numbers = ['' for _ in range(n)]
    results = []
    for row in matrix:
        for j, i in enumerate(range(0, n * 4, 4)):
            numbers[j] += row[i:i+4]
    results = [str(STR_TO_NUMBER[s]) for s in numbers]    
    return results


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    matrix = []
    for _ in range(5):
        matrix.append(stdin.readline().strip())

    print(''.join(solve(matrix, n)))
