import sys



stdin = sys.stdin



class seq:
    def __init__(self, n):
        self._n = n
        self._p = None

    def __iter__(self):
        return self
    

    def __next__(self):
        n_o_d = len(self._n)
        self._p = self._n 
        self._n = str(n_o_d)
        
        if self._p == self._n:
            raise StopIteration
        else:
            return n_o_d


def solve(n):
    for i, _ in enumerate(seq(n), 1): ...
    return locals().get('i', 0) + 1


if __name__ == '__main__':
    for row in stdin:
        if row.strip() == 'END':
            break
        print(solve(row.strip()))
