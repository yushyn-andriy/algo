import sys



stdin = sys.stdin
EPSILON = 10**-4


def get_judgment_day(H, U, D, F):
    i = 1
    fatigue = U - U * (1 - F / 100.0)
    distance_climbed = 0
    while True:        
        if U >= 0:
            distance_climbed += U
        if distance_climbed - H > 0:
            return f'success on day {i}'

        distance_climbed -= D
        if distance_climbed < 0:
            return f'failure on day {i}'

        U = U - fatigue
        i += 1


if __name__ == '__main__':
    while True:
        H, U, D, F = [
            float(x)
            for x in stdin.readline().strip().split()
        ]
        if H == 0:
            break
        
        print(get_judgment_day(H, U, D, F))
