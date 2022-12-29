import sys
from copy import deepcopy


stdin = sys.stdin



class Pattern:
    def __init__(self, pattern):
        self._p = pattern
        self._n = len(pattern)
        self._origin = deepcopy(pattern)

    def reset(self):
        self._p = deepcopy(self._origin)

    def rot90(self):
        _p = deepcopy(self._p)
        for i in range(self._n):
            #print(''.join(self._p[i]))
            for j in range(self._n):
                # print(f'({i}, {j}) -> ({j}, {self._n - i - 1}) == {self._p[i][j]}')
                _p[j][self._n - i - 1] = self._p[i][j]
            #print('-'*10)
        self._p = _p
        return Pattern(deepcopy(self._p))

    def reflect(self):
        _p = deepcopy(self._p)
        for i in range(self._n//2):
            for j in range(self._n):
                self._p[self._n - i - 1][j], self._p[i][j] = (
                        self._p[i][j],
                        self._p[self._n - i -1][j])
        #self._p = _p
        return Pattern(deepcopy(self._p))


    def __eq__(self, other):
        return self._p == other

    def __repr__(self):
        return '\n'.join([''.join(row) for row in self._p])


if __name__ == '__main__':
    _id = 1
    while True:
        n = stdin.readline().strip()
        if not n:
            break
        n = int(n)
        p1, p2 = [], []
        for i in range(n):
            c1, c2 = stdin.readline().strip().split()
            p1.append(list(c1))
            p2.append(list(c2))

        p1, p2 = Pattern(p1), Pattern(p2)
        if p1 == p2:
            print(f'Pattern {_id} was preserved.')
        else:
            eq = False
            for refl in [False, True]:
                if eq:
                    break
                if refl:
                    p1.reset()
                    p1.reflect()
                    eq = p1 == p2
                    if eq:
                        print(f'Pattern {_id} was reflected vertically.')
                        break

                for deg in [90, 180, 270]:
                    p1.rot90()
                    eq = p1 == p2
                    if eq and not refl:
                        print(f'Pattern {_id} was rotated {deg} degrees.')
                        break
                    elif eq and refl:
                        print(f'Pattern {_id} was reflected vertically and rotated {deg} degrees.')
                        break
            if not eq:
                print(f'Pattern {_id} was improperly transformed.')

        _id += 1
