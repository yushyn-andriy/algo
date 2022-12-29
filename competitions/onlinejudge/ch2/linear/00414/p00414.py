import sys



stdin = sys.stdin

# O(N) time | O(N) space
def get_total_void(rows):
    counts = []
    for row in rows:
        counts.append(row.count(' '))
    min_void = min(counts)
    total_void = 0
    for c in counts:
        if c > min_void:
            total_void += c - min_void
    return total_void


if __name__ == '__main__':
    while True:
        n_rows = int(stdin.readline().strip())
        if n_rows == 0:
            break
        rows = [None] * n_rows
        for i in range(n_rows):
            rows[i] = list(stdin.readline().strip('\n'))
        print(get_total_void(rows))

