def for_else_ok():
    for i in range(3):
        print('for_else_ok:', i)
    else:
        print('for_else_ok: else branch')

def for_else_early_break():
    for i in range(3):
        print('for_else_early_break:', i)
        break
    else:
        print('for_else_early_break: else branch')



def while_else_ok():
    i = 0
    while i<3:
        print('while_else_ok:', i)
        i+=1
    else:
        print('while_else_ok: else branch')


def while_else_early_break():
    i = 0
    while i < 3:
        print('while_else_early_break:', i)
        i += 1
        break
    else:
        print('while_else_early_break: else branch')


def try_else_ok():
    try:
        print('try_else_ok: try')
    except:
        print('try_else_ok: except')
    else:
        print('try_else_ok: else')


def try_else_exception():
    try:
        print('try_else_ok: try')
        raise Exception
    except:
        print('try_else_ok: except')
    else:
        print('try_else_ok: else')



for_else_ok()
print('+'*20)
for_else_early_break()
print('+'*20)
print()
while_else_ok()
print('+'*20)
while_else_early_break()

print('+'*20)
print()
try_else_ok()
print('+'*20)
try_else_exception()



from io import StringIO, BytesIO
import os



text = '''
Hello world!
1
2
3
'''
rbytes = b'''
Hello world
1
2
3
'''

sio = StringIO(
    text,
    os.linesep,
)
print(sio.readlines())

bio = BytesIO(rbytes)
for line in bio:
    print('pos=', bio.tell())
    # bio.write(b'a')
    print(line.decode('utf-8').strip())


class ContextManager:
    def __enter__(self):
        print('enter')
        import sys
        self.original_write = sys.stdout.write
        sys.stdout.write = self.back_write
        return self

    def back_write(self, s):
        self.original_write(s[::-1])
        

    def __exit__(self, exc_type, exc_value, traceback):
        print('exit', exc_type, exc_value, traceback)
        import sys
        sys.stdout.write = self.original_write
        if exc_type is ZeroDivisionError:
            print('division by nil do not do things like this')
            return True


with ContextManager() as a:
    print(a)
    print('hello world!!!')
    raise ZeroDivisionError

print('hello world!!!')
