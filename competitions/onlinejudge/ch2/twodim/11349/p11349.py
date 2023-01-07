import sys


stdin = sys.stdin


# O(N**2) time | O(N**2) space
def solve(m):
    """
    Two main cases when N is even and N is odd.
    """
    N = len(m)
    if N == 1:
        if m[0][0] >= 0:
            return True
        else:
            return False

    for i in range(N//2):
        for j in range(len(m[0])):
            # check positivity
            if m[i][j] < 0:
                return False

            # check diagonal symmetric element
            if m[i][j] != m[N - 1 - i][N - 1 - j]:
                return False




    if N % 2 != 0:
        middle_row = N // 2
        for i in range(N):
            if m[middle_row][i] < 0 or m[middle_row][i] != m[middle_row][N - 1 - i]:
                return False            

    return True

if __name__ == '__main__':
    T = int(stdin.readline().strip())
    for i in range(1, T + 1):
        dim = int(stdin.readline().strip().split('=')[1])
        m = []
        for _ in range(dim):
            m.append([int(x) for x in stdin.readline().strip().split()])

        is_symm = solve(m)
        result = 'Symmetric'
        if not is_symm:
            result = 'Non-symmetric'
        print(f'Test #{i}: {result}.')
