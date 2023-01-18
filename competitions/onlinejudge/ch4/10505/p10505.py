import sys


stdin = sys.stdin


UNVISITED = '.'
WHITE = 'W'
BLACK = 'B'


def opposite(c):
    if c == WHITE:
        return BLACK
    return WHITE


class Graph:
    __slots__ = ('outgoing', )

    def __init__(self):
        self.outgoing = {}
    

    def incident_vertices(self, v):
        if v not in self.outgoing:
            return []
        return list(self.outgoing[v].keys())

    def insert_edge(self, u, v):
        if v not in self.outgoing:
            self.outgoing[v] = {}
        if u not in self.outgoing:
            self.outgoing[u] = {}
        self.outgoing[u][v] = True
        self.outgoing[v][u] = True


    def insert_vertex(self, v):
        if v not in self.outgoing:
            self.outgoing[v] = {}

    def vertices(self):
        return set(self.outgoing.keys())


def helper_dfs(g, vertex, discovered, color, parent, bw):
    discovered[vertex] = color
    if color == BLACK:
        bw[0] += 1
    else:
        bw[1] += 1

    is_bipartite = True
    for nbr in g.incident_vertices(vertex):
        if discovered[nbr] == UNVISITED:
            subprob = helper_dfs(g, nbr, discovered, opposite(color), vertex, bw)
            if subprob is False:
                is_bipartite = False
        elif nbr != parent and discovered[nbr] == color:
                is_bipartite = False

    return is_bipartite


def maximum_invited(g, discovered):
    max_invited = 0
    for vertex in g.vertices():
        if discovered[vertex] == UNVISITED:
            bw = [0, 0]
            is_bipartite = helper_dfs(g, vertex, discovered, BLACK, -1, bw)
            if is_bipartite is False:
                continue
            max_invited += max(bw)
    return max_invited


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    stdin.readline().strip()
    for _ in range(1, c+1):
        n = int(stdin.readline().strip())
        g = Graph()
        for node in range(1, n+1):
            row = [int(x) for x in stdin.readline().strip().split()]
            g.insert_vertex(node)
            enemies = row[1:]
            for e in enemies:
                if e > n:
                    continue
                g.insert_edge(node, e)
        stdin.readline()
        discovered = [UNVISITED] * 201
        print(maximum_invited(g, discovered))
