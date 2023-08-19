class Quantity:
    __counter = 0

    def __init__(self) -> None:
        cls = self.__class__
        prefix = cls.__name__
        postfix = str(self.__counter)
        self.storage_name = f'{prefix}_{postfix}'
        cls.__counter += 1

    def __set__(self, instance, value):
        print('set')
        instance.__dict__[self.storage_name] = value
    
    def __get__(self, instance, cls):
        print('get', instance, cls)
        return instance.__dict__[self.storage_name]


class Product:
    price = Quantity()
    price1 = Quantity()
    price2 = Quantity()
    price3 = Quantity()
    weight = Quantity()


p = Product()

p.price = 2
p.weight = 100
print(p.price)
print(vars(p))
