import turtle



def generate_sequance(n_iter):
    sequance = list('LLR')
    tmp_sequance = [] 
    for _ in range(n_iter):
        lr = list('LR') * ((len(sequance) + 1) // 2)
        i, j = 0, 0
        while i < len(sequance):
            tmp_sequance.append(lr[j]) 
            tmp_sequance.append(sequance[j])
            i+=1
            j+=1
        else:
            tmp_sequance.append(lr[j])
            sequance = tmp_sequance
            tmp_sequance = []
    return sequance



def draw_dragons_curve(n):
    skk = turtle.Turtle()
    skk.speed(2000)
    sequance = generate_sequance(n)
    skk.pensize(1)
    skk.forward(10)

    for op in sequance:
        if op == 'L':
            skk.left(90)
        elif op == 'R':
            skk.right(90)

        skk.forward(10)

    turtle.done()


if __name__ == '__main__':
    draw_dragons_curve(15)
