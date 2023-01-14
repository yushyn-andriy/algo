import sys


stdin = sys.stdin


class Graph:
    """..."""
    def __init__(self, directed=False):
        self._outgoing = {}
        self._incoming = {} if directed else self._outgoing

    def is_directed(self):
        return self._outgoing is not self._incoming
    
    def vertices(self):
        return self._outgoing.keys()

    def incident_vertices(self, v, outgoing=True):
        adj = self._outgoing if outgoing else self._incoming
        if v not in adj:
            return []
        return list(adj[v].keys())

    def insert_vertex(self, v):
        if v not in self._outgoing:
            self._outgoing[v] = {}
        if self.is_directed() and v not in self._incoming:
            self._incoming[v] = {}
        return v

    def insert_edge(self, u, v, x=None):
        if self.is_directed():
            self._outgoing[u][v] = True
        else:
            self._outgoing[u][v] = True
            self._outgoing[v][u] = True

# DFS modified for dominator problem
def DFS(g: Graph, root: int, discovered, exc=None, y=None):
    for e in g.incident_vertices(root):
        if exc is not None and exc == e:
            continue
        if y is not None and y == e:
            discovered[e] = None
            return
        if e not in discovered:
            discovered[e] = None
            DFS(g, e, discovered, exc)

def is_dominator(g: Graph, root: int, x: int, y:int):
    discovered = {root: None}
    if root == x:
        DFS(g, root, discovered, exc=None, y=y)
        if y in discovered:
            return True
    if x == y:
        DFS(g, root, discovered, exc=None, y=y)
        if y in discovered:
            return True
        else:
            return False
    DFS(g, root, discovered, exc=None, y=y)
    if y not in discovered:
        return False
    discovered = {root: None}
    DFS(g, root, discovered, exc=x, y=y)
    if y not in discovered:
        return True
    if y in discovered:
        return False


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for c in range(n):
        g = Graph(directed=True)
        nodes = int(stdin.readline().strip())
        for i in range(nodes):
            row = [int(x) for x in stdin.readline().strip().split()]
            for j, v in enumerate(row):
                if v == 1:
                    g.insert_vertex(i)
                    g.insert_vertex(j)
                    g.insert_edge(i, j)
        print(f'Case {c+1}:')
        for i in range(nodes):
            x = i
            answers = []
            for j in range(nodes):
                y = j
                answers.append('Y' if is_dominator(g, 0, x, y) else 'N')
            s = f'|{"|".join(answers)}|'
            print(f'+{"-" * (len(s) - 2)}+')
            print(s)
        print(f'+{"-" * (len(s) - 2)}+')


