import sys
import heapq


stdin = sys.stdin



if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        if '0' == row:
            break
        heap = list(map(int, stdin.readline().strip().split()))
        heapq.heapify(heap)
        
        result = 0
        while len(heap)>1: 
            s = (heapq.heappop(heap) + heapq.heappop(heap))
            heapq.heappush(heap, s)
            result += s
        print(result) 
