import sys
from collections import namedtuple


Person =  namedtuple('Person', 'name day month year')


stdin = sys.stdin


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    persons = []
    for _ in range(c):
        row = stdin.readline().strip()
        sd = row.split()
        name = sd[0]
        persons.append(Person(name, *[int(x) for x in sd[1:]])) 


    persons = sorted(persons, key=lambda x: (x.year, x.month, x.day))
    print(persons[-1].name)
    print(persons[0].name)
