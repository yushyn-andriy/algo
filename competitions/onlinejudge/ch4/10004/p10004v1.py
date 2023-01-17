import sys


stdin = sys.stdin


UNVISITED = 0
WHITE = 1
BLACK = 2

def opposite_color(color):
    return 3 - color


class Graph:

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


def helper_dfs(g, node, color, colors, parent):
    colors[node] = color

    for vertex in g.incident_vertices(node):
        if colors[vertex] == UNVISITED:
            subprob = helper_dfs(g, vertex, opposite_color(color), colors, node)
            if subprob is False:
                return False
        elif vertex != parent and color == colors[vertex]:
            return False    
    return True


def is_bipartite(g, root):
    colors = [UNVISITED] * 201
    return helper_dfs(g, root, BLACK, colors, -1)



if __name__ == '__main__':
    while True:
        n = int(stdin.readline().strip())
        if n == 0:
            break
        e = int(stdin.readline())
        g = Graph()
        for i in range(e):
            u, v = [int(x) for x in stdin.readline().strip().split()]
            g.insert_edge(u, v)
            root = u
        res = is_bipartite(g, root)
        print('BICOLORABLE.' if res else 'NOT BICOLORABLE.')
