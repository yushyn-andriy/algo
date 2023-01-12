import sys
from collections import namedtuple



Person = namedtuple('Person', [
    'title',
    'first_name',
    'last_name', 
    'address',
    'home_phone',
    'work_phone',
    'campus_box',
    'department',
])

stdin = sys.stdin



if __name__ == '__main__':
    n = int(stdin.readline().strip())
    persons = []
    for _ in range(n):
        department = stdin.readline().strip()
        for row in stdin:
            if not row.strip():
                break
            
            data = row.strip().split(',')
            data.append(department)
            p = Person(*data)
            persons.append(p)

    persons = sorted(persons, key=lambda x: x.last_name)
    for p in persons:
        print('-'*40)
        print(f'{p.title} {p.first_name} {p.last_name}')
        print(f'{p.address}')
        print(f'Department: {p.department}')
        print(f'Home Phone: {p.home_phone}')
        print(f'Work Phone: {p.work_phone}')
        print(f'Campus Box: {p.campus_box}')
            