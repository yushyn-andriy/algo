import sys
import heapq
from collections import deque


stdin = sys.stdin



if __name__ == '__main__':
    for row in stdin:
        k = int(row.strip())
        
        stack = deque([])
        queue = deque([])
        priority_queue = []
        ns, nsf = 0, 0
        nq, nqf = 0, 0
        np, npf= 0, 0
        n=0
        

        result = []
        for i in range(k):
            op, val = list(map(int, stdin.readline().strip().split()))
            if op == 1:
                stack.append(val)
                queue.append(val)
                heapq.heappush(priority_queue, -val)
                n+=1
            elif op == 2:
                try:
                    sval = stack.pop()
                    if sval == val and nsf == 0:
                        ns += 1
                    else:
                        nsf+=1
                except IndexError:
                    nsf += 1
                
                try:
                    qval = queue.popleft()
                    if qval == val and nqf == 0:
                        nq += 1
                    else:
                        nqf +=1
                except IndexError:
                    nqf += 1

                try:
                    pval = abs(heapq.heappop(priority_queue))
                    if pval == val and npf == 0:
                        np += 1
                    else:
                        npf += 1
                except IndexError:
                    npf +=1

        heapq.heappush(result, (-ns, nsf, 'stack'))
        heapq.heappush(result, (-np, npf,   'priority queue'))
        heapq.heappush(result, (-nq, nqf, 'queue'))
        r1, r2, r3 = heapq.heappop(result), heapq.heappop(result), heapq.heappop(result)
        if r1[1] != 0:
            print('impossible')
        else:
            if r2[1] == 0 and r1[0] == r2[0]:
                print('not sure')
            else:
                print(r1[2])
