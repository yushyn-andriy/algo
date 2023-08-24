import sys


stdin = sys.stdin



class Engraver:
    def __init__(self, h, l, X):
        self.h = h
        self.l = l
        self.X = X
        self.surface = []
        self.counter = 0
        self.on = False

        # self.build(self.h, self.l)

    def build(self, h, l):
        self.surface = [
            list('x'*l)
            for _ in range(h)
        ]

    def turn_on(self):
        if self.on is False:
            self.on = True
            self.counter += 1


    def turn_off(self):
        self.on = False


    # def engrave(self):
    #     for y in range(self.h):
    #         self.turn_off()
    #         for x in range(self.l):
    #             current_height = self.h - y
    #             # print(current_height, self.X[x], x)
    #             if current_height > self.X[x] and self.surface[y][x] == 'x':
    #                 # print('turn on', self.on)
    #                 self.turn_on()
    #                 self.surface[y][x] = 'o'
    #             elif self.surface[y][x] == 'x':
    #                 self.turn_off()
    #     return self.surface

    def engrave(self):
        for y in range(self.h):
            self.turn_off()
            for x in range(self.l):
                current_height = self.h - y
                # print(current_height, self.X[x], x)
                if current_height > self.X[x]:
                    # print('turn on', self.on)
                    self.turn_on()
                    # self.surface[y][x] = 'o'
                else:
                    self.turn_off()
        return self.surface


    def __str__(self) -> str:
        s = ''
        for row in self.surface:
            s += ''.join(row)
            s += '\n'
        return s


if __name__ == '__main__':
    # for row in stdin:
    #     a, c = list(map(int, row.strip().split()))
    #     if a == 0 and c == 0:
    #         break
     
    #     X = list(map(int, stdin.readline().strip().split()))
        
    #     engraver = Engraver(a, c, X)
    #     engraver.engrave()
    #     print(engraver.counter)
    #     # print(str(engraver))
    #     # print('='*10)

    while True:
        A, C = map(int, input().split())
        if A == 0 and C == 0:
            break
        
        heights = list(map(int, input().split()))
        
        total_turns = 0
        current_height = A
        
        for h in heights:
            if h < current_height:
                total_turns += current_height - h
            current_height = h
        
        print(total_turns)

