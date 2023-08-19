# property()


from typing import Any


class A:
    def __init__(self, weight) -> None:
        self.__weight = weight


    def get_weight(self):
        print('get_weigth')
        return self.__weight
    
    def set_weight(self, value):
        print('set_weight')
        self.__weight = value
    

    weight = property(get_weight, set_weight, None, 'property weight')


a = A(10)

a.weight = 2
print(a.weight)


class A:
    data = 'QWERTY'

    @property
    def prop_val(self):
        return 'some strange value'
    

a = A()
print('a.data =', a.data)
a.data = 'dasd'
print('vars(a) =', vars(a))
print('a.data = ', a.data)
print('A.data =', A.data)


print('a.prop_val =', a.prop_val)
try:
    a.pro_val = 123
except AttributeError:
    print('exception a.prop_val = 123')

a.__dict__['prop_val'] = 54321
print('a.prop_val =', a.prop_val)
print('A.prop_val = ', A.prop_val)
A.prop_val = 123
# del A.prop_val
print('a.prop_val =', a.prop_val)


# search begins from obj.__class__


def quantity(storage_name):
    def getter(instance):
        print('getter')
        return instance.__dict__[storage_name]
    
    def setter(instance, value):
        print('setter')
        instance.__dict__[storage_name] = value
    
    return property(getter, setter, None, f'property {storage_name}')


class Product:
    price = quantity('price')
    weight = quantity('weight')

    def __init__(self, price, weight) -> None:
         self.price = price
         self.weight = weight


p = Product(1, 2)
p.price = 2
p.price
p.weight = 4
print(p.weight)


class TestGetattr:
    abc = 'qwerty'

    def __getattr__(self, attr):
        # calls only when no attribute in 
        print('gettattr', attr)


tg  = TestGetattr()
tg.abc = 123
print(tg.abc)

del tg.abc
del TestGetattr.abc
print(tg.abc)