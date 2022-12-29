import sys

stdin = sys.stdin


def get_rings(N):
    return N // 2 + N % 2

def to_string(sq):
    return "\n".join([
        " ".join([str(y) for y in x])
        for x in sq
    ])


def transform(sq, ring, ops):
    N = len(sq)
    is_even = len(sq) % 2 == 0
    for op in ops:
        if is_even:
            width =  2 * ring
        else:
            width =  2 * ring - 1
        begin = int((N - width) / 2)
        end = begin + width - 1

        if op == 1: # upside down
            i = begin
            while i <= end:
                sq[begin][i], sq[end][i] = sq[end][i], sq[begin][i]
                i += 1
            i = begin + 1
            j = end - 1
            while i < j:
                sq[i][begin], sq[j][begin] = sq[j][begin], sq[i][begin]
                sq[i][end], sq[j][end] = sq[j][end], sq[i][end]
                i += 1
                j -= 1   
        elif op == 2: # left right flip
            i = begin
            while i <= end:
                sq[i][begin], sq[i][end] = sq[i][end], sq[i][begin]
                i += 1
            i = begin + 1
            j = end - 1
            while i < j:
                sq[begin][i], sq[begin][j] = sq[begin][j], sq[begin][i]
                sq[end][i], sq[end][j] = sq[end][j], sq[end][i]
                i += 1
                j -= 1
        elif op == 4: # flip through  inverse main diagonal
            i = begin
            j = end
            while i <= end:
                sq[begin][i], sq[j][end] = sq[j][end], sq[begin][i]
                i += 1
                j -= 1
            
            i = begin + 1
            j = end - 1
            while i <= end:
                sq[i][begin], sq[end][j] = sq[end][j], sq[i][begin]
                i += 1
                j -= 1
        elif op == 3:
            i = begin
            while i <= end:
                sq[begin][i], sq[i][begin] = sq[i][begin], sq[begin][i]
                i += 1
            
            i = begin + 1
            while i < end:
                sq[i][end], sq[end][i] = sq[end][i], sq[i][end] 
                i += 1


if __name__ == "__main__":
    M = int(stdin.readline().strip())
    for _ in range(M):
        N = int(stdin.readline().strip())
        sq = [None] * N
        for i in range(N):
            sq[i] = [int(x) for x in stdin.readline().strip().split()]
        T = get_rings(N)
        for j in range(0, T):
            arr = [int(x) for x in stdin.readline().strip().split()]
            r_n = arr[0]
            ops = arr[1:]
            transform(sq, T - j, ops)
        print(to_string(sq))
