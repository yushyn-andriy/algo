import sys



stdin = sys.stdin


def find(i, parent):
    if parent[i] == -1:
        return i
    parent[i] = find(parent[i], parent)
    return parent[i]


def is_connected(a, b, parent):
    i = find(a, parent)
    j = find(b, parent)
    return i == j


def union(a, b, parent, rank):
    s1 = find(a, parent)
    s2 = find(b, parent)

    if s1 != s2:
        if rank[s1] < rank[s2]:
            parent[s1] = s2
            rank[s2] += rank[s1]
        else:
            parent[s2] = s1
            rank[s1] += rank[s2]


if __name__ == '__main__':
    t = int(sys.stdin.readline().strip())
    stdin.readline()
    for i in range(t):
        n = int(stdin.readline().strip())
        success = 0
        unsuccess = 0
        parent = [-1 for _ in range(n)]
        rank = [1 for _ in range(n)]
        for row in stdin:
            row = row.strip()
            if not row:
                break
            L = row.split()
            op= L[0]
            a, b = list(map(lambda x: int(x) - 1, L[1:]))
            if op == 'c':
                union(a, b, parent, rank)
            elif op == 'q':
                res = is_connected(a, b, parent)
                if res:
                    success += 1
                else:
                    unsuccess += 1

        print(f'{success},{unsuccess}')
        if i == t - 1:
            continue
        else:
            print()