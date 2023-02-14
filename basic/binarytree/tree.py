class Node:
    def __init__(self, key, value=None, parent=None) -> None:
        self.key = key
        self.value = value
        self.left = None
        self.right = None
        self.parent = parent

    def __repr__(self) -> str:
        return f'Node({self.key})'


class Tree:
    def __init__(self, root: Node=None) -> None:
        self.root =  root


    def insert(self, key, value=None):
        z = Node(key, value) 
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


if __name__ == '__main__':
    tree = Tree()

    tree.insert(10)
    tree.insert(2)
    tree.insert(6)
    tree.insert(7)
    tree.insert(1)
    tree.insert(15)

    sorted_list = []
    def add(r, *args):
        global sorted_list
        sorted_list.append(r)


    tree.inorder_walk(print)
    tree.inorder_walk(add)

    print(sorted_list)


    print(tree.search(1))
    print(tree.search(7))
    print(tree.search(712))

    print('='*25)

    print(tree.min(), tree.max())

    print('='*25)
    print(tree.successor(tree.search(2)))
    print(tree.predsessor(tree.search(6)))
