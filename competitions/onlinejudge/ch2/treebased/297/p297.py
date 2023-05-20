import sys
from collections import deque


stdin = sys.stdin


class Quadtree:
    __slots__ = ('value', 'children')

    def __init__(self, value) -> None:
        self.value = value
        self.children = []

    def add_child(self, child):
        self.children.append(child)

    def __repr__(self) -> str:
        return f'{self.value}'

    def __str__(self) -> str:
        return f'{self.value}'


def print_tree(t, level=0):
    print('\t'* level + f'{t}')
    for child in t.children:
        print_tree(child, level=level+1)



def create_quedtree(t, s, level=0):
    counter = 0
    while s and counter < 4:
        ch = s.popleft()
        counter += 1
        node = Quadtree(ch)
        t.add_child(node)
        if ch == 'p':    
            create_quedtree(node, s, level + 1)


def eat(s):
    counter = 0
    while s and counter < 4:
        counter += 1
        ch = s.popleft()
        if ch == 'p':
            eat(s)


def create_merged_quedtree(root, d1, d2, level=0):
    counter = 0
    while (d1 or d2) and counter < 4:
        counter += 1
        if d1 and d2:
            ch1 = d1.popleft()
            ch2 = d2.popleft()
            if ch1 == 'p' and ch2 == 'p':
                node = Quadtree('p')
                create_merged_quedtree(node, d1, d2, level+1)
            elif ch1 == 'p' and ch2 == 'e':
                node = Quadtree('p')
                create_quedtree(node, d1)
            elif ch2 == 'p' and ch1 == 'e':
                node = Quadtree('p')
                create_quedtree(node, d2)
            elif ch1 == 'p' and ch2 == 'f':
                eat(d1)
                node = Quadtree('f')
            elif ch2 == 'p' and ch1 == 'f':
                eat(d2)
                node = Quadtree('f')
            else:
                if 'f' in [ch1, ch2]:
                    node = Quadtree('f')
                else:
                    node = Quadtree('e')
            root.add_child(node)
        elif d1:
            ch1 = d1.popleft()
            node = Quadtree(ch1)
            root.add_child(node)
            if ch1 == 'p':    
                create_merged_quedtree(node, d1, d2, level + 1)
        elif d2:
            ch2 = d2.popleft()
            node = Quadtree(ch2)
            root.add_child(node)
            if ch2 == 'p':    
                create_merged_quedtree(node, d1, d2, level + 1)


def total_black(root, power=5):
    count = 0
    if root.value == 'f':
        count += 4 ** power
    elif root.value == 'p':
        for child in root.children:
            count += total_black(child, power - 1)
    return count


if __name__ == '__main__':
    n = int(stdin.readline().strip())

    for _ in range(n):
        d1 = deque(list(stdin.readline().strip()))
        d2 = deque(list(stdin.readline().strip()))
        

        t = Quadtree('root')
        create_merged_quedtree(t, d1, d2)
        print(f'There are {total_black(t.children[0])} black pixels.')
