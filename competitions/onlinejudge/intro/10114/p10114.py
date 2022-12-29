import sys



def get_monthes(months, down_payment, loan, deprecation_records):
    EPSILON = 10**-4

    deprications = [0]*(1 + int(months))
    for _ in range(int(deprecation_records)):
        m, d = [float(i) for i in sys.stdin.readline().strip().split()]
        deprications[int(m)] = d
    for i in range(int(months)):
        if abs(deprications[i] - 0) < EPSILON: 
            deprications[i] = deprications[i - 1]


    worth = (loan + down_payment) * (1 - deprications[0])
    owes = loan
    month_payment = loan / months

    current_month = 0
    index = 1
    while owes - worth > EPSILON:
        current_deprication = deprications[index]
        owes -= month_payment
        worth = worth * (1 - current_deprication)
        current_month += 1
        index += 1
    
    return current_month



if __name__ == '__main__':
    while True:
        args = [float(i) for i in sys.stdin.readline().strip().split()]
        if args[0] < 0:
            break

        months = get_monthes(*args)
        if months > 1 or months == 0:
            result = f'{months} months'
        else:
            result = f'{months} month'
        print(result)
