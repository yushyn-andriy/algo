import os
import pickle


import multiprocessing




class ChildProcess(multiprocessing.Process):
    def __init__(self, pipein):
        self.pipein = pipein
        super().__init__()
    
    def run(self):
        self.pipein = os.fdopen(self.pipein, 'wb')
        data = ['string', 777]
        pickle.dump(data, self.pipein)
        # self.pipein.write('My name is ChildProcess')
        self.pipein.close()
    

if __name__ == '__main__':
    pipeout, pipein = os.pipe()

    p = ChildProcess(pipein)
    p.start()
    p.join()

    os.close(pipein)

    pipeout = os.fdopen(pipeout, 'rb')
 
    # pipeContent = pipeout.read()
    pipeContent = pickle.load(pipeout)
    print(f'Pipe sad: {pipeContent}')
