"""
Disjoint Set Union.
"""


"""

Let there be 4 elements 0, 1, 2, 3

Initially, all elements are single element subsets.
0 1 2 3 

Do Union(0, 1)
   1   2   3  
  /
 0

Do Union(1, 2)
     2   3   
    /
   1
 /
0

Do Union(2, 3)
         3    
        /
      2
     /
   1
 /
0
"""
# O(NlogN)
"""
Let us see the above example with union by rank
Initially, all elements are single element subsets.
0 1 2 3 

Do Union(0, 1)
   1   2   3  
  /
 0

Do Union(1, 2)
   1    3
 /  \
0    2

Do Union(2, 3)
    1    
 /  |  \
0   2   3

"""


# O(N) time
def find(i, parent):
    if parent[i] == -1:
        return i
    parent[i] = find(parent[i], parent)
    return parent[i]


# O(N) time
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


# O(N) time
def contains_cycle(edge_list, V, parent, rank):
    
    for i, j in edge_list:
        s1, s2 = find(i, parent), find(j, parent)
        if s1 != s2:
            union(s1, s2, parent, rank)
        else:
            return True
    return False




rank = [1, 1, 1, 1, 1]
parent = [-1, -1, -1, -1, -1]
edge_list = [
    (1, 2),
    (3, 2),
    (3, 4),
    (3, 1),
]



union(3, 4, parent, rank)
print(parent)

union(2, 3, parent, rank)
print(parent)

union(1, 2, parent, rank)
print(parent)

print('rank', rank)

v= 4
parent = [-1] * (4+1)
print(contains_cycle(edge_list, v, parent, rank))
print(parent)
