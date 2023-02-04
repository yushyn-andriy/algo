import sys
from itertools import permutations


stdin = sys.stdin


def backtrack(s1, s2, s1_idx, s2_idx, ops, sta):
    if s2_idx == len(s2):
        print(' '.join(ops))
        return

    if s1_idx < len(s1):
        ops.append('i')
        sta.append(s1[s1_idx])
        backtrack(s1, s2, s1_idx+1, s2_idx, ops, sta)
        sta.pop()
        ops.pop()

    if sta and sta[-1] == s2[s2_idx]:
        temp = sta.pop()
        ops.append('o')
        backtrack(s1, s2, s1_idx, s2_idx+1, ops, sta)
        ops.pop()
        sta.append(temp)
    return 



if __name__ == '__main__':
    for row in stdin:
        s1 = row.strip()
        s2 = stdin.readline().strip()
        print('[')
        seqs = backtrack(s1, s2, 0, 0, [], [])
        print(']')
        