import random
import subprocess
import sys
import difflib as dl
import os


if __name__ == '__main__':
    iteration = 0
    for i in range(10000):
        iteration += 1
        if iteration % 100 == 0:
            print('iteration', iteration)

        n = 15
        m = random.randint(2, 15)
        numbers = []
        while len(numbers) != n:
            rn = random.randint(-1000000, 1000000)
            if rn == 0:
                continue
            else:
                numbers.append(rn)
        

        with open('input.txt', 'w') as f:
            print(n, m, file=f)
            for rn in numbers:
                print(rn, file=f)
            print(0, 0, file=f)


        res1, res2 = [], []
        with open('input.txt', mode='r') as stdin:
            cmd = ['./a.exe']
            process = subprocess.Popen(cmd, stdin=stdin, stdout=subprocess.PIPE)
            for line in process.stdout:
                res1.append(line)


        with open('input.txt', mode='r') as stdin:
            cmd = ['python', './p11321.py']
            process = subprocess.Popen(cmd, stdin=stdin, stdout=subprocess.PIPE)
            for line in process.stdout:
                res2.append(line)
    


        if res1 != res2:

            print(res1)
            print(res2)
            print('different output has been found...', iteration)
            print('terminating programm...')
            sys.exit(0)
