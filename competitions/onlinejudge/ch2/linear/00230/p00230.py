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


AUTHOR = 1
TITLE = 0
BORROW = 'BORROW'
RETURN = 'RETURN'
SHELVE = 'SHELVE'


if __name__ == '__main__':
    stock = read_stock()
    commands = read_commands()

    sorted_books = sorted(stock.items(), key=lambda x: (x[AUTHOR], x[TITLE]))
    books_position = {
        title: i
        for i, (title, author) in enumerate(sorted_books)
    }
    borrowed = [False] * len(stock)
    returned = [False] * len(stock)

    for command, title in commands:
        if command == BORROW:
            position = books_position[title]
            borrowed[position] = True
            returned[position] = False
        elif command == RETURN:
            position = books_position[title]
            returned[position] = True
        elif command == SHELVE:
            previous = -1
            for i, (title, author) in enumerate(sorted_books):
                if not borrowed[i]:
                    previous = i
                elif returned[i]:
                    if previous == -1:
                        print(f'Put {title} first')
                    else:
                        print(f'Put {title} after {sorted_books[previous][TITLE]}')

                    returned[i] = False
                    borrowed[i] = False
                    previous = i
            print('END')


