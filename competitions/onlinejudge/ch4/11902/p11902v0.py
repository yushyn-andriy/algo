import sys


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

# DFS modified for dominator problem
def DFS(g: Graph, root: Vertex, discovered, exc=None, y=None):
    for e in g.incident_vertices(root):
        #print('\troot:', root, e)
        if exc is not None and exc == e:
            #print('\t\tcontinue', e)
            continue
        
        if y is not None and y == e:
            discovered[e] = None
            return


        if e not in discovered:
            discovered[e] = None
            DFS(g, e, discovered, exc)


def is_dominator(g, root, x, y):
    
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

    # should be reachable from root to y
    DFS(g, root, discovered, exc=None, y=y)
    #print(discovered)
    if y not in discovered:
        #print(f'1. {y} not in {discovered}')
        return False

    # should not be reachable without intermediate x node
    discovered = {root: None}
    DFS(g, root, discovered, exc=x, y=y)
    #print(discovered)
    if y not in discovered:
        #print(f'2. {y} not in {discovered}')
        return True

    # if reachable that there is another way
    if y in discovered:
        #print(f'3. {y} in {discovered}')
        return False


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for c in range(n):
        g = Graph(directed=True)
        nodes = int(stdin.readline().strip())
        for i in range(nodes):
            v1 = Vertex(i)
            g.insert_vertex(v1)
            
            row = [int(x) for x in stdin.readline().strip().split()]
            for j, v in enumerate(row):
                if v == 1:
                    v2 = Vertex(j)
                    g.insert_vertex(v2)
                    g.insert_edge(v1, v2)
        
        print(f'Case {c+1}:')

        root = Vertex(0)
        vertices = list(g.vertices())
        for i in range(nodes):
            x = Vertex(i)
            answers = []
            for j in range(nodes):
                y = Vertex(j)
                dom = is_dominator(g, root, x, y)
                answer = 'Y' if dom else 'N'
                answers.append(answer)
            
            s = f'|{"|".join(answers)}|'
            print(f'+{"-" * (len(s) - 2)}+')
            print(s)
        print(f'+{"-" * (len(s) - 2)}+')

        #print(is_dominator(g, root, Vertex(0), Vertex(1)), c)
