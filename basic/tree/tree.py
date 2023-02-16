"""
Красно-черные деревья (свойства)

1. Каждый узел либо красный либо белый
2. Корень черного цвета
3. Каждый лист черный
4. Если узел красный то оба его дочерних элемента черные
5. Для каждого узла все простые пути от него до листьев являющихся потомками 
   данного узла содержат одно и то же количество черных узлов.



Tree-Insert O(log(n)) time
Tree-Delete O(log(n)) time

"""

from timeit import timeit

import enum


class Color(enum.Enum):
    RED = 0
    BLACK = 1




class Node:
    def __init__(self, key, value=None, color=Color.RED, parent=None) -> None:
        self.key = key
        self.value = value
        self.left = None
        self.right = None
        self.color = color
        self.parent = parent

    def __repr__(self) -> str:
        return f'Node({self.key}, {self.color})'


class Tree:
    def __init__(self, root: Node=None) -> None:
        self.root =  root


    def insert(self, key, value=None, color=None):
        z = Node(key=key, value=value, color=color) 
        y = None
        x = self.root

        while x is not None:
            y = x
            if z.key < x.key:
                x = x.left
            else:
                x = x.right
        
        z.parent = y
        if y is None:
            self.root = z
        elif z.key < y.key:
            y.left = z
        else:
            y.right = z

        return z

    def _inorder_walk(self, node, func=None, level=0):
        if node is None:
            return
        self._inorder_walk(node.left, func, level=level+1)
        if func is not None:
            func(node, level)
        self._inorder_walk(node.right, func, level=level+1)


    def inorder_walk(self, func=None):
        self._inorder_walk(self.root, func)


    def _preorder_walk(self, node, func=None, level=0):
        if node is None:
            return

        if func is not None:
            func(node, level)
        self._preorder_walk(node.left, func, level=level+1)
        self._preorder_walk(node.right, func, level=level+1)


    def preorder_walk(self, func=None):
        self._preorder_walk(self.root, func)

    def _search(self, node, k):
        if node is None or node.key == k:
            return node
        
        if k < node.key:
            return self._search(node.left, k)
        elif k >= node.key:
            return self._search(node.right, k)


    def search(self, k):
        return self._search(self.root, k)

    def iterative_search(self, k):
        x = self.root
        y = x
        while x is not None and x.key != k:
            if k >= y.key:
                x = x.right
            else:
                x = x.left 
            y = x

        return y


    def min(self, node=None):
        x = self.root
        if node:
            x = node
        while x.left is not None:
            x = x.left
        return x
    
    def max(self, node=None):
        x = self.root
        if node:
            x = node
        while x.right is not None:
            x = x.right
        return x

    def successor(self, node):
        if node.right is not None:
            return self.min(node.right)

        y = node.parent
        while y is not None and node is y.right:
            node = y
            y = y.p
        return y


    def predsessor(self, node):
        if node.left is not None:
            return self.max(node.right)

        y = node.parent
        while y is not None and node is y.left:
            node = y
            y = y.p
        return y


    def transplant(self, u, v):
        if u.parent is None:
            self.root = v
        elif u.parent.left is u:
            u.parent.left = v
        elif u.parent.right is u:
            u.parent.right = v
        if v is not None:
            v.parent = u.parent


    def delete(self, z):
        if z.left is None:
            self.transplant(z, z.right)
        elif z.right is None:
            self.transplant(z, z.left)
        else:
            y = self.min(z.right)
            if y.parent is not z:
                self.transplant(y, y.right)
                y.right = z.right
                y.right.parent = y
            self.transplant(z, y)
            y.left = z.left
            y.left.parent = y

    def left_rotate(self, x):
        y = x.right
        x.right = y.left
        if y.left is not None:
            y.left.parent = x

        y.parent = x.parent
        if x.parent is None:
            self.root = y
        elif x is x.parent.left:
            x.parent.left = y
        elif x is x.parent.right:
            x.parent.right = y
        y.left = x

    def right_rotate(self, x):
        y = x.left
        x.left = y.right
        if y.right is not None:
            y.right.parent = x

        y.parent = x.parent
        if x.parent is None:
            self.root = y
        elif x.parent.left is x:
            x.parent.left = y
        elif x.parent.right is x:
            x.parent.right = y
        y.right = x


    def rb_fixup(self, z):
        while z.parent and z.parent.color == Color.RED:
            if z.parent is z.parent.parent.left:  # parent on the left branch
                y = z.parent.parent.right
                if y is not None and y.color == Color.RED: # just recolor and up 2 level
                    z.parent.color = Color.BLACK
                    y.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    z = z.parent.parent
                else:
                    if z is z.parent.right: # line it to the left branch
                        z = z.parent
                        self.left_rotate(z)
                    # 1 color parent to black
                    # 2 color grand parent to red
                    # 3 rotate to the right 
                    z.parent.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    self.right_rotate(z.parent.parent)
            else: # parent on right branch 
                y = z.parent.parent.left
                if y is not None and y.color == Color.RED: # just recolor and up 2 level
                    z.parent.color = Color.BLACK
                    y.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    z = z.parent.parent
                else:
                    if z is z.parent.left: # line it to the left branch
                        z = z.parent
                        self.right_rotate(z)
                    # 1 color parent to black
                    # 2 color grand parent to red
                    # 3 rotate to the right 
                    z.parent.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    self.left_rotate(z.parent.parent)
        
        self.root.color = Color.BLACK           

    def rb_insert(self, key, value=None):
        z = self.insert(key, value, color=Color.RED)
        self.rb_fixup(z)


def bt_tree(n=10**4):
    t = Tree()
    for v in range(n):
        t.insert(v)
    return t 


def rb_tree(n=10**4):
    t = Tree()
    for v in range(n):
        t.rb_insert(v)
    return t 



def search(tree, v):
    return tree.iterative_search(v)




if __name__ == '__main__':
    n = 10000
    bt = bt_tree(n)
    rb = rb_tree(n)

    print('BT search elapsed time:', timeit('search(bt, n-1)', number=10000, globals=globals()))
    print('RB search elapsed time:', timeit('search(rb, n-1)', number=10000, globals=globals()))
    # cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
    # BT search elapsed time: 6.366588499862701
    # RB search elapsed time: 0.011413399828597903
