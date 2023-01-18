import sys
from collections import deque, namedtuple

stdin = sys.stdin



Info = namedtuple('Info', ['id', 'degree'])


class Graph:
    def __init__(self, directed=False):
        self.outgoing = {}
        self.ingoing = {}
        self.directed = directed

    def insert_vertex(self, v):
        if v not in self.outgoing:
            self.outgoing[v] = {}
        if v not in self.ingoing:
            self.ingoing[v] = {}
    
    def insert_edge(self, u, v):
        self.outgoing[u][v] = True
        self.ingoing[v][u] = True 
    
    def remove_ingoing(self, v, u):
        if u in self.ingoing[v]:
            del self.ingoing[v][u]

    def incident_vertices(self, v):
        return list(self.outgoing[v].keys())


    def degree(self, v):
        return len(self.ingoing[v].keys())


def get_ordered_by_degrees(g, visited):
    degrees = {}
    for v in range(1, n+1):
        if v not in visited:
            degrees[v] = g.degree(v)
    return sorted([
        Info(v, d)
        for v, d in degrees.items()
    ], key=lambda x: x.degree)


def topological_ordering(g):
    result = []
    visited = set()
    ordered_nodes = get_ordered_by_degrees(g, visited)


    root = ordered_nodes[0]
    if root.degree != 0:
        return result


    d = deque([root.id])
    while len(d) > 0:
        root = d.popleft()
        result.append(root)
        visited.add(root)

        for nbr in g.incident_vertices(root):
            g.remove_ingoing(nbr, root)

        ordered_nodes = get_ordered_by_degrees(g, visited)
        for node in ordered_nodes:
            if node.degree == 0 and node.id not in visited:
                d.append(node.id)
                visited.add(node.id)

    return result


if __name__ == '__main__':
    while True:
        n, m = [int(x) for x in stdin.readline().strip().split()]
        if n == 0:
            break

        g = Graph(directed=True)
        for v in range(1, n+1):
            g.insert_vertex(v)
        for _ in range(m):
            u, v = [int(x) for x in stdin.readline().strip().split()]
            g.insert_edge(u, v)


        ordered = topological_ordering(g) 
        print(' '.join([str(x) for x in ordered]))