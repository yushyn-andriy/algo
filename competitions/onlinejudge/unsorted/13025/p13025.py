import sys
from datetime import datetime


stdin = sys.stdin


if __name__ == '__main__':
    t = datetime.strptime('29/05/2013', '%d/%m/%Y')
    print(datetime.strftime(t, '%B %d, %Y %A'))

