def object_maker(the_class, args):
    # new_object = the_class.__new__(the_class, args) -> object.__new__(the_class, args)
    new_object = object.__new__(the_class, args)
    the_class.__init__(new_object)
    return new_object




class Bar:
    def __init__(self):
        pass
    
    def say_hi(self):
        print('Hi')

o =object_maker(Bar, None)        
o.say_hi()
print(o.__class__)