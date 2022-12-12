import sys


GREATINGS_COUNTRY = {
    'HELLO': 'ENGLISH',
    'HOLA': 'SPANISH',
    'HALLO': 'GERMAN',
    'BONJOUR': 'FRENCH',
    'CIAO': 'ITALIAN',
    'ZDRAVSTVUJTE': 'RUSSIAN',
} 


if __name__ == '__main__':
    case = 1
    while True:
        line = sys.stdin.readline().strip()
        if line == '#':
            break

        print(f'Case {case}: {GREATINGS_COUNTRY.get(line, "UNKNOWN")}')
        case += 1
