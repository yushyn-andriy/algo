import multiprocessing
import sys
import random


def get_mem_stat(pid):
    with open(f'/proc/{pid}/status', 'r') as f:
        lines = f.readlines()
        for line in lines:
            if 'VmSize' in line:
                break
        return int(line.split()[1])


def find_prime_numbers(n):
    d = 2
    factors = []
    while d*d<=n:
        while (n % d) == 0:
            factors.append(d)
            n //= d
        d += 1
    if n > 1:
        factors.append(n)
    return factors


def execute_process(start, stop):
    for i in range(start, stop+1):
        print(i, find_prime_numbers(i))


def main():
    processes: list[multiprocessing.Process] = []
    for (start, stop) in [
        (1, 125_000),
        (125_000, 250_000),
        (250_000, 375_000),
        (375_000, 500_000),
        (500_000, 625_000),
        (625_000, 750_000),
        (750_000, 875_000),
        (875_000, 1000_000),
    ]:
        process = multiprocessing.Process(target=execute_process, args=(start, stop))
        processes.append(process)
        process.start()

    with open('memory_usage.txt', 'a') as f:
        for p in processes:
            print(f'pid: {p.pid}, VmSize: {get_mem_stat(p.pid)}', file=f)


    p: multiprocessing.Process = random.choice(processes)
    with open('log.txt', 'w') as f:
        print('terminating: ', p.pid, file=f)
    p.terminate()




    for i in processes:
        i.join()

    return 0


if __name__ == '__main__':
    sys.exit(main())
