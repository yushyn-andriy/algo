from typing import Any
from collections import abc


class FrozenJSON:
    def __init__(self, data):
        self.__data = dict(data)

    
    def __getattr__(self, __name: str) -> Any:
        if hasattr(self.__data, __name):
            # if it is a dict the for example it can has .keys attribute 
            # print(self.__data, __name, hasattr(self.__data, __name))
            return getattr(self.__data, __name)
        return FrozenJSON.build(self.__data[__name])

    @classmethod
    def build(cls, obj):
        if isinstance(obj, abc.Mapping):
            return cls(obj)
        elif isinstance(obj, abc.MutableSequence):
            return [cls.build(o) for o in obj]
        else:
            return obj
        raise TypeError


d = {
    'a': [{}],
    'b': {
        'c': 123
    }
}


# d = FrozenJSON(d)

# print(d.a.pop(), d.b.keys())



class FrozenJSONImpr:
    def __init__(self, data) -> None:
        print('__init__', data)
        self.__data = dict(data)

    def __new__(cls, obj):
        print('__new__', obj)
        # all the arguments provided to __new__ then passed to __init__!!!
        if isinstance(obj, abc.Mapping):
            obj = super().__new__(cls)
            return obj
        elif isinstance(obj, abc.MutableSequence):
            return [cls(o) for o in obj]
        else:
            return obj
    
    def __getattr__(self, __name: str) -> Any:
        if hasattr(self.__data, __name):
            return getattr(self.__data, __name)
        return FrozenJSONImpr(self.__data[__name])

d = FrozenJSONImpr(d)

print(d.a, d.b.keys())