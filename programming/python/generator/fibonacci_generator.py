def fibonacci(n):
    a, b = 0, 1
    yield a
    yield b
    for _ in range(n):
        a, b = b, a + b
        yield b

# 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144.

print(list(fibonacci(12)))
