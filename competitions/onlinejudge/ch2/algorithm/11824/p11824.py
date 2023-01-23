import sys


stdin = sys.stdin


# O(N) time | O(N) space
def min_price(land_prices, total_amount=5000000):
    land_prices = sorted(land_prices, reverse=True)
    curr_price = 0
    for i, price in enumerate(land_prices):
        curr_price += int(2*pow(price, i+1))
        if curr_price > total_amount:
            return -1
    return curr_price 




if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for _ in range(c):
        land_prices = []
        for row in stdin:
            row = row.strip()
            price = int(row)
            if price == 0:
                break
            
            land_prices.append(price)
        mp = min_price(land_prices)
        if mp == -1:
            print('Too expensive')
        else:
            print(mp)
        