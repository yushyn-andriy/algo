"""
Disjoint Set Union.
"""


#      1
#     / \
#    2   3
#         \
#          4

# [-1, 1, 1, 3]
#   1  2  3  4


# O(N) time
def find(i, parent):
    if parent[i] == -1:
        return i
    return find(parent[i], parent)


# O(N) time
def union(a, b, parent):
    s1 = find(a, parent)
    s2 = find(b, parent)
    if s1 != s2:
        parent[s2] = s1


# O(N) time
def contains_cycle(edge_list, V, parent):
    
    for i, j in edge_list:
        s1, s2 = find(i, parent), find(j, parent)
        if s1 != s2:
            union(s1, s2, parent)
        else:
            return True
    return False





parent = [-1, -1, -1, -1, -1]
edge_list = [
    (1, 2),
    (3, 2),
    (3, 4),
    (3, 1),
]


v= 4
parent = [-1] * (4+1)
print(contains_cycle(edge_list, v, parent))
print(parent)
