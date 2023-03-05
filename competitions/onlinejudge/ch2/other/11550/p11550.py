import sys


stdin = sys.stdin


def check_incident_matrix_correctness(matrix):
    n, m = len(matrix), len(matrix[0])
    branches_registry = set()
    for c in range(m):
        non_zero_rows = []
        column_sum = 0
        for r in range(n):
            v = matrix[r][c]
            column_sum += v

            # column sum should not exceed 2
            if column_sum > 2:
                return False

            if v == 1:
                non_zero_rows.append(r)
        

        # no branches lead to nowhere
        if len(non_zero_rows) == 1 or len(non_zero_rows) == 0:
            return False


        # there are should not be duplicated branches
        if len(non_zero_rows) == 2:
            coords = tuple(non_zero_rows)
            if coords in branches_registry:
                return False
            else:
                branches_registry.add(coords)
        

    return True

if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for _ in range(t):
        row = stdin.readline().strip().split()
        n, m = list(map(int, row))
        
        matrix = []
        for _ in range(n):
            row = list(map(int, stdin.readline().strip().split()))
            matrix.append(row)


        incident = check_incident_matrix_correctness(matrix=matrix)
        print('Yes' if incident else 'No')
