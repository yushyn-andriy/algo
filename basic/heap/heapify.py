def left(i):
    return i*2

def right(i):
    return i*2 + 1

def parent(i):
    return i//2


def max_heapify(heap, i, heapsize):
    l, r = left(i), right(i)
    largest = i
    if l < heapsize and  heap[l] > heap[i]:
        largest = l
    if r < heapsize and heap[r] > heap[largest]:
        largest = r
    
    if largest != i:
        heap[i], heap[largest] = heap[largest], heap[i]
        max_heapify(heap, largest, heapsize)



def build(heap):
    for i in range(len(heap) //2, 0, -1):
        max_heapify(heap, i, len(heap))


def heapsort(array):
    res = []
    build(array)
    heapsize = len(array)
    for i in range(len(array)-1, 0, -1):
        res.append(array[1])
        array[i], array[1] = array[1], array[i]
        heapsize -= 1
        max_heapify(array, 1, heapsize)

    return res

def extract_max(heap, hz):
    print()
    print(heap, hz)
    hz = hz - 1
    key = heap[1]
    heap[1], heap[hz - 1] = heap[hz-1], heap[1]
    max_heapify(heap, 1, hz)
    return key, hz


def insert(heap, val):
    
    el = heap.pop(0)
    heap.insert(0, val)
    heap.insert(0, el)
    max_heapify(heap, 1, len(heap))


def heap_increase_key(heap, i, k):
    if k<heap[i]:
        raise ValueError
    heap[i] = k
    while k > heap[parent(i)] and i>1:
        heap[parent(i)], heap[i] = heap[i], heap[parent(i)]
        i = parent(i)



l = [0, 7, 4, 1, 8, 9, 4, 6, 7]

# print(l)
# build(l)
# print(l)

print(heapsort(l))
build(l)
print(l)
insert(l, 5)
heap_increase_key(l, 5, 10)
print(l)
print(heapsort(l))

# print('=======================')
# k, hs = extract_max(l, len(l))
# print(k, hs)
# k, hs = extract_max(l, hs)
# print(k, hs)
# k, hs = extract_max(l, hs)
# print(k, hs)

"""
               9
           /       \
          8         6
       /    \     /   \
      7      4   4     1
    /
   7
"""
