import sys
from math import log2


stdin = sys.stdin



def parent(n):
    return n // 2


def solve(N, i, j):
    total_participants = 2**N
    last_level_nodes_count = total_participants // 2

    curr_node1 = (i+1) // 2 + (last_level_nodes_count - 1)
    curr_node2 = (j+1) // 2 + (last_level_nodes_count - 1)

    round = 1
    while curr_node1 != curr_node2:
        curr_node1 = parent(curr_node1)
        curr_node2 = parent(curr_node2)
        round += 1

    return round


if __name__ == '__main__':
    for row in stdin:
        N, i, j = list(map(int, row.split()))
        print(solve(N, i, j))
