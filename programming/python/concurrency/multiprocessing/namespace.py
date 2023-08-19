import multiprocessing


def modify(nn):
    nn.x = -678
    print(f'pid: {multiprocessing.current_process().pid}')




if __name__ == '__main__':
    manager = multiprocessing.Manager()
    ns = manager.Namespace()
    ns.x = 10

    print('ns.x', ns.x)
    p = multiprocessing.Process(target=modify, args=(ns, ))
    p.start()
    p.join()

    print('ns.x', ns.x)
