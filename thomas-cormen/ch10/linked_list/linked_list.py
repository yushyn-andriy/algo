

class DoubleLinkedList:
    head = None

    def __init__(self, key, data=None):
        self.key = key
        self.data = data
        self.next = None
        self.prev = None

    def insert(self, x):
        if self.head is None:
            self.head = x
        else:
            x.next = self.head
            self.head.prev = x
            self.head = x

    def delete(self, key):
        x = self.search(key)
        if x:
            if x.prev is None and x.next is None:
                self.head = None
            elif x.prev is None:
                x.next.prev = None
                self.head = x.next
            elif x.next is None:
                x.prev.next = None
            else:
                x.prev.next = x.next
                x.next.prev = x.prev

    def search(self, key):
        current = self.head
        while current != None:
            if current.key == key:
                return current
            current = current.next
        return None

    def traverse(self, func):
        current = self.head
        while current != None:
            func(current)
            current = current.next

    def __repr__(self) -> str:
        return f'DoubleLinkedList({self.key}, {self.data})'

if __name__ == '__main__':
    dl = DoubleLinkedList(1, 'first')
    dl.head = dl

    dl.insert(DoubleLinkedList(2, 'second'))
    dl.insert(DoubleLinkedList(3, 'third'))
    print(dl.search(1))
    dl.traverse(lambda x: print(x.key, x.data))

    dl.delete(1)
    dl.delete(3)
    dl.delete(2)
    dl.insert(DoubleLinkedList(3, 'third'))
    dl.delete(3)
    dl.traverse(lambda x: print(x.key, x.data))
