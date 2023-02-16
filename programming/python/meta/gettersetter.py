class LineItem:
    def __init__(self, name, weight, price):
        self.name = name
        self.weight = weight
        self.__price = price

    
    @property
    def weight(self):
        return self.__weight


    @weight.setter
    def weight(self, value):
        if value < 0:
            raise ValueError('Value should be greeter or equal zero')
        else:
            self.__weight = value
    

    def subtotal(self):
        return self.__weight * self.__price



apples = LineItem('apples', 1.32, 20.0)
apples.weight = 1.56
print(apples.subtotal())
apples.weight = -1.56

