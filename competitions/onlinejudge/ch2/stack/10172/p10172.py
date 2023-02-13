import sys
from collections import OrderedDict, deque




cin = sys.stdin


class CargoStation:
    def __init__(self, _id, b_max_cargoes,  A=None, B=None):
        self._id = _id
        self.b_max_cargoes = b_max_cargoes
        self.A = deque([]) if not A else deque(A)
        self.B = deque([]) if not B else deque(B)


    def is_full(self):
        return len(self.B) >= self.b_max_cargoes

    def is_empty(self):
        return len(self.B) == 0


    def __repr__(self) -> str:
        return f'CargoStation(_id={self._id})'


class Carrier:
    def __init__(self, stations, max_cargos, total_containers, unloading_time=1, loading_time=1, transition_time=2):
        self.stations = stations
        self.total_containers = total_containers
        self.max_cargos = max_cargos

        self.cargo = deque([])

        self.unloading_time = unloading_time
        self.loading_time = loading_time
        self.transition_time = transition_time

        self.total_time = 0

    def is_empty(self):
        return len(self.cargo) == 0

    def is_full(self):
        return len(self.cargo) >= self.max_cargos


    def carrying_time(self):
        while True:
            current_arranged_containers = 0
            for i, station in self.stations.items():
                while not self.is_empty():
                    container = self.cargo.pop()
                    if container == station._id:
                        station.A.append(container)
                    else:
                        if station.is_full():
                            self.cargo.append(container)
                            break
                        station.B.append(container)

                    # print("unloading station:", station, station.A, station.B)
                    # print("unloading cargo:", self.cargo)   
                    # print()


                    self.total_time += self.unloading_time
                

                current_arranged_containers += len(station.A)

                while not station.is_empty() and not self.is_full():
                    self.cargo.append(station.B.pop())
                    self.total_time += self.loading_time

                    # print("uploading station:", station, station.A, station.B)
                    # print("uploading cargo:", self.cargo)   
                    # print()




                if current_arranged_containers == self.total_containers:
                    break

                self.total_time += self.transition_time

            if current_arranged_containers == self.total_containers:
                    break

        return self.total_time

import io


cin = io.StringIO('''2
5 2 3
3 4 5 2
2 1 3
0
3 3 5 1
1 4
5 2 3
3 4 5 2
2 1 3
0
3 3 5 1
1 4''')

if __name__ == '__main__':  
    cases = int(cin.readline().strip())
    for _ in range(cases):
        N, S, Q = [int(x) for x in cin.readline().strip().split()]


        stations = OrderedDict()
        total_containers = 0
        for i in range(N):
            args = [int(x) for x in cin.readline().strip().split()]
            total_containers+=args[0]
            cargos = args[1:]
            stations[i+1] = CargoStation(i+1, Q, B=cargos)

        print('Max Cargo', S)
        car = Carrier(stations, S, total_containers)
        print(car.carrying_time())
