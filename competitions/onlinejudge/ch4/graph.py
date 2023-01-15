import sys
from collections import deque


stdin = sys.stdin


class Vertex:
    """..."""
    __slots__ = ('_element', )

    def __init__(self, x=None):
        self._element = x
    
    def __hash__(self) -> int:
        return hash(id(self._element))
    
    def __eq__(self, other):
        if isinstance(other, Vertex):
            return self._element == getattr(other, 'element')()
        return False

    def element(self):
        return self._element

    def __repr__(self):
        return f'Vertex({self._element})'

class Edge:
    """..."""
    __slots__ = ('_origin', '_destination', '_element')
    
    def __init__(self, v, u, x=None):
        self._origin = v
        self._destination = u
        self._element = x
    
    def __hash__(self) -> int:
        return hash((self._origin, self._destination))

    def __repr__(self):
        return f'Edge(origin={self._origin}, destination={self._destination})'
    
    def element(self):
        return self._element
    
    def endpoints(self):
        return (self._origin, self._destination)
    
    def opposite(self):
        return self._destination


class Graph:
    """..."""
    def __init__(self, directed=False):
        self._outgoing = {}
        self._incoming = {} if directed else self._outgoing

    def is_directed(self):
        return self._outgoing is not self._incoming
    
    def vertex_count(self):
        return len(self._outgoing)
    
    def vertices(self):
        return self._outgoing.keys()
    
    def edge_count(self):
        total = sum(len(self._outgoing[v]) for v in self._outgoing)
        return total if self.is_directed() else total // 2

    def edges(self):
        result = set()
        for secondary_map in self._outgoing.values():
            result.update(secondary_map.values())
        return result
    
    def get_edge(self, u, v):
        return self._outgoing[u].get(v)
    
    def degree(self, v, outgoing=True):
        adj = self._outgoing if outgoing else self._incoming
        return len(adj[v])

    def incident_edges(self, v, outgoing=True) -> list[Vertex]:
        adj = self._outgoing if outgoing else self._incoming
        for edge in adj[v].values():
            yield edge
    
    def incident_vertices(self, v, outgoing=True):
        adj = self._outgoing if outgoing else self._incoming
        for vertex in adj[v].keys():
            yield vertex

    def insert_vertices(self, vertices):
        for v in vertices:
            self.insert_vertex(v)
        return self.vertices()

    def insert_vertex(self, v):
        if v not in self._outgoing:
            self._outgoing[v] = {}
        if self.is_directed() and v not in self._incoming:
            self._incoming[v] = {}
        return v

    def insert_edge(self, u, v, x=None):
        if self.is_directed():
            e = Edge(u, v, x)
            self._outgoing[u][v] = e
        else:
            e1 = Edge(u, v, x)
            e2 = Edge(v, u, x)
            self._outgoing[u][v] = e1
            self._outgoing[v][u] = e2


def DFS(g: Graph, u: Vertex, discovered, f=None):
    for e in g.incident_vertices(u):
        if e not in discovered:
            if f is not None:
                f(e)
            discovered[e] = None
            DFS(g, e, discovered, f)

def BFS(g, root, discovered, f=None):
    d = deque([root])
    while len(d) != 0:
        node = d.popleft()
        if f:
            f(node)
        for el in g.incident_vertices(node):
            if el not in discovered:
                discovered[el] = True
                d.append(el)



def pprint(e):
    print(e.element())

if __name__ == '__main__':
    n = int(stdin.readline().strip())
    g = Graph()
    for i in range(n):
        vals = [int(x) for x in stdin.readline().strip().split()]
        v, u = Vertex(vals[0]), Vertex(vals[1])
        g.insert_vertex(v)
        g.insert_vertex(u)
        g.insert_edge(v, u)
    
    discovered = {}
    #DFS(g, list(g.vertices())[0], discovered, pprint)
    BFS(g, list(g.vertices())[0], discovered, print)