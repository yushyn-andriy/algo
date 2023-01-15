import sys
import string
from collections import deque


letters = string.ascii_uppercase

stdin = sys.stdin


class Graph:
    def __init__(self):
        self._outgoing = {}
    
    def insert_edge(self, x, y):
        if x not in self._outgoing:
            self._outgoing[x] = {}
        if y not in self._outgoing:
            self._outgoing[y] = {}
        self._outgoing[x][y] = True
        self._outgoing[y][x] = True
    
    def incident_vertices(self, x):
        result = set()
        if x not in self._outgoing:
            return result
        for inner_map in self._outgoing[x].keys():
            result.update(inner_map)
        return result


def BFS(g, discovered, root, f=None):
    d = deque([root])
    while len(d) > 0:
        node = d.popleft()
        if f:
            f(node)
        for new_node in g.incident_vertices(node):
            if new_node in discovered:
                continue
            else:
                discovered.add(new_node)
                d.append(new_node)



def count_maximum_subgraphs(g, largest_letter):
    discovered = set()
    count = 0
    for l in letters:
        if l > largest_letter:
            break
        if l in discovered:
            continue

        BFS(g, discovered, l)
        count += 1
    return count


     


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    stdin.readline()
    for i in range(n):
        g = Graph()
        largest_letter = stdin.readline().strip()
        for row in stdin:
            row = row.strip()
            if not row:
                break
            x, y = list(row)
            g.insert_edge(x, y)


        print(count_maximum_subgraphs(g, largest_letter))
        if i == n - 1:
            continue
        print()
