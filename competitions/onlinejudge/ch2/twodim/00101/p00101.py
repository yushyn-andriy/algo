import sys



stdin = sys.stdin



class BlockWorld:
    def __init__(self, n):
        self.n = n
        self.loc = {}
        self.blocks = [None] * self.n
        for i in range(n):
            self.blocks[i] = [i]
            self.loc[i] = (i, 0)

    def _parse_cmd(self, raw):
        l = raw.split()
        cmd = (l[0], l[2])
        a, b = int(l[1]), int(l[3])
        return cmd, a, b

    def _initial(self, blk):
        blks = self.blocks
        loc = self.loc

        r, c = loc[blk]
        s = self.blocks[r]
        while len(s):
            if s[-1] == blk:
                break
            b = s.pop()
            blks[b].append(b)
            loc[b] = (b, len(blks[b]) -1 )

    def _move_onto(self, a, b):
        self._initial(a)
        self._initial(b)
        r1, _ = self.loc[a]
        r2, _ = self.loc[b]
        a = self.blocks[r1].pop()
        self.blocks[r2].append(a)
        self.loc[a] = (r2, len(self.blocks[r2]) - 1)

    def _move_over(self, a, b):
        self._initial(a)
        r1, _ = self.loc[a]
        r2, _ = self.loc[b]
        a = self.blocks[r1].pop()
        self.blocks[r2].append(a)
        self.loc[a] = (r2, len(self.blocks[r2]) - 1)


    def _pile_onto(self, a, b):
        self._initial(b)
        r1, _ = self.loc[a]
        r2, _ = self.loc[b]

        stack = self.blocks[r1]
        l = []
        while len(stack):
            bl = stack.pop()
            l.append(bl)
            if bl == a:
                break
        while len(l):
            bl = l.pop()
            self.blocks[r2].append(bl)
            self.loc[bl] = (r2, len(self.blocks[r2]) - 1)



    def _pile_over(self, a, b):
        r1, c1 = self.loc[a]
        r2, c2 = self.loc[b]


        stack = self.blocks[r1]
        l = []
        while len(stack):
            bl = stack.pop()
            l.append(bl)
            if bl == a:
                break
        while len(l):
            bl = l.pop()
            self.blocks[r2].append(bl)
            self.loc[bl] = (r2, len(self.blocks[r2]) - 1)


    def _get_handler(self, cmd):
        return getattr(self, '_' + '_'.join(cmd))


    def execute(self, raw):
        if raw == 'quit':
            return self.blocks
        cmd, a, b = self._parse_cmd(raw)
        handler = self._get_handler(cmd)
        handler(a, b)
        return self.blocks



if __name__ == '__main__':
    n = int(stdin.readline().strip())
    bw = BlockWorld(n)
    i = 0
    while True:
        cmd = stdin.readline().strip()
        if cmd == 'quit':
            break
        i+=1
        blks = bw.execute(cmd)

    blks = bw.blocks
    for i in range(len(blks)):
        if len(blks[i]) == 0:
            print(f'{i}:')
        else:
            print(f'{i}:', ' '.join([str(x) for x in blks[i]]))

