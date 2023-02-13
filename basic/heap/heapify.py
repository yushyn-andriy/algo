def left(i):
    return i*2

def right(i):
    return i*2 +1

def parent(i):
    return i//2


def max_heapify(heap, i):
    l, r = left(i), right(i)
    print(l, r)
    largest = i
    if l < len(heap) and  heap[l] > heap[i]:
        largest = l
    if r < len(heap) and heap[r] > heap[largest]:
        largest = r
    
    if largest != i:
        heap[i], heap[largest] = heap[largest], heap[i]
        max_heapify(heap, largest)



def build(heap):
    for i in range(len(heap) //2, 0, -1):
        max_heapify(heap, i)


l = [0, 7, 4, 1, 8, 9, 4, 6, 7]

print(l)
build(l)
print(l)
"""
               9
           /       \
          8         6
       /    \     /   \
      7      4   4     1
    /
   7
"""



