import sys


stdin = sys.stdin



def read_stock():
    stock = {}
    while True:
        line = stdin.readline().strip()
        if line == 'END':
            break
        title, author = [
                x.strip()
                for x in
                line.split('by')
            ]
        stock[title] = author
    return stock


def read_commands():
    commands = []
    while True:
        line = stdin.readline().strip()
        if line == 'END':
            break
        command = line.split(' "')
        if len(command) == 2:
            commands.append((command[0], '"' + command[1]))
        else:
            commands.append((command[0], None))
    return commands


if __name__ == '__main__':
    stock = read_stock()
    commands = read_commands()


    shelve = stock.copy()
    returned = {}
    for command, arg in commands:
        if command == 'BORROW':
            del shelve[arg]
            if arg in returned:
                del returned[arg]
        elif command == 'RETURN':
            shelve[arg] = stock[arg]
            returned[arg] = stock[arg]
        elif command == 'SHELVE':
            # sort by authors and by book name
            shelve_list = sorted(shelve.items(), key=lambda x: (x[1], x[0]))
            for i, (name, author) in enumerate(shelve_list):
                if i == 0 and name in returned:
                    print(f'Put {name} first')
                elif name in returned:
                    print(f'Put {name} after {shelve_list[i-1][0]}')

            # empty after each SHELVE command
            returned = {}
            print('END')


