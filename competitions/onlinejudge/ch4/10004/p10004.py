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
    # if colored same as previous node 
    if node != parent and colors[node] == color:
        return False
    
    opp_color = opposite_color(color)

    # if unvisited color node to opposite color
    if colors[node] == UNVISITED:
        colors[node] = opp_color
    # already visited and colored to the wright color
    elif colors[node] == opp_color:
        return True

    for vertex in g.incident_vertices(node):
        res = helper_dfs(g, vertex, opp_color, colors, node)
        if res is False:
            return False
    return True


def is_bipartite(g, root, n):
    colors = [UNVISITED] * 201
    return helper_dfs(g, root, BLACK, colors, -100)



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

        res = is_bipartite(g, root, n)
        print('BICOLORABLE.' if res else 'NOT BICOLORABLE.')
