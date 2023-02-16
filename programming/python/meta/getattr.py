class Foo:
    def __getattr__(self, attr):
        print(attr)

    def __setattr__(self, __name, __value) -> None:
        print(__name, __value)


f = Foo()

f.a
f.b
f.xyz
f.zxc = '123asd'
f.xc = 'qwerty'
