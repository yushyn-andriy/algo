import sys



stdin = sys.stdin


class Signal:
    RED = 'RED'
    GREEN = 'GREEN'
    YELLOW = 'YELLOW'

    def __init__(self, cycle=10):
        self.cycle = cycle
        self.red = self.cycle
        self.yellow = 5
        self.green = self.red - self.yellow
        self.current_time = 0
        self.current_state = self.RED
        self.red_last = self.red
        self.yellow_last = self.yellow
        self.green_last = self.green

    def _reset(self):
        self.red_last = self.red
        self.yellow_last = self.yellow
        self.green_last = self.green

    def next_second(self):
        self.current_time += 1
        if self.current_state == self.RED:
            self.red_last -= 1
            if self.red_last == 0:
                self.current_state = self.GREEN
                self._reset()
        elif self.current_state == self.GREEN:
            self.green_last -= 1
            if self.green_last == 0:
                self.current_state = self.YELLOW
                self._reset()
        else:
            self.yellow_last -= 1
            if self.yellow_last == 0:
                self.current_state = self.RED
                self._reset()

        return self.current_state

    def __repr__(self):
        return self.current_state


def readlist(converter):
    return [converter(x) for x in stdin.readline().strip().split()]


if __name__ == '__main__':
    idx = 1
    while True:
        l = readlist(int)
        if not l:
            break
        signals = [Signal(cycle) for cycle in l]


        simula = {}
        HOUR = 60 * 60
        faced_green = False
        faced_yellow = False
        first_faced_green_time = 0
        first_faced_yellow_time = 0
        for i in range(HOUR):
            #print(f'time:{i+1} {signals[0].current_state}')
            if not faced_green:
                if all([signal.current_state == signal.GREEN for signal in signals]):
                    faced_green = True
                    first_faced_green_time = i
                    #print('faced green', i)
            elif not faced_yellow and any([signal.current_state == signal.YELLOW for signal in signals]):
                faced_yellow = True
                first_faced_yellow_time = i
                #print('faced yellow', i)
            elif faced_green and faced_yellow:
                if all([signal.current_state == signal.GREEN for signal in signals]):
                    # print(signals)
                    print(idx, i - first_faced_green_time)# , i - first_faced_yellow_time)
                    break

            for signal in signals:
                signal.next_second()

        idx += 1

