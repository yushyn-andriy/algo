import sys


stdin = sys.stdin


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(n):
        n_instructions = int(stdin.readline().strip())

        x_pos = 0
        log = {}
        for j in range(1, n_instructions + 1):
            instruction = stdin.readline().strip().split()    
            if len(instruction) == 1:
                op = instruction[0]
            else:
                op = log[int(instruction[2])]

            log[j] = op
            if op == 'LEFT':
                x_pos -= 1
            elif op == 'RIGHT':
                x_pos += 1
        print(x_pos)
