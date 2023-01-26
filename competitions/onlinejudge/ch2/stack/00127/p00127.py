import sys
from collections import namedtuple



stdin = sys.stdin



class DLL:
    def __init__(self, pile):
        self.pile = pile
        self.left = None
        self.right = None



Card = namedtuple('Card', 'rank suit')

def traverse(node):
    count = 0
    l = []
    while node is not None:
        count += 1
        l.append(len(node.pile))
        node = node.right
    return count, l


def is_eq(a, b):
    return a.suit == b.suit or a.rank == b.rank


def game(deck):
    root = DLL(deck[0])
    prev = root
    for pile in deck[1:]:
        node = DLL(pile)
        prev.right = node
        node.left = prev
        prev = node


    node = root
    while node is not None:
        if node.left is None:
            node = node.right
            continue

        
        left_node = node.left
        for _ in range(2):
            if left_node is None:
                break
            left_node = left_node.left
        else:
            if left_node and is_eq(left_node.pile[-1], node.pile[-1]):
                left_node.pile.append(node.pile.pop())
                if not node.pile:
                    node.left.right = node.right
                    if node.right is not None:
                        node.right.left = node.left
                    
                node = left_node
                continue


        if is_eq(node.left.pile[-1], node.pile[-1]):
            node.left.pile.append(node.pile.pop())
            left_node = node.left
            if not node.pile:
                node.left.right = node.right
                if node.right is not None:
                    node.right.left = node.left
            
            node = left_node
            continue

        node = node.right
    
    count, length = traverse(root)
    return count, [str(x) for x in length]



if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        if row == '#':
            break
        
        
        deck = [[Card(*list(x))] for x in row.split()]
        row = stdin.readline().strip()
        for x in row.split():
            deck.append([Card(*list(x))])
        
        count, length = game(deck)
        print(f'{count} piles remaining: {" ".join(length)}')
