
#SEM USAR THREADS
def tarefa1() -> None:
    x = 0
    while x < 10:
        print("Tarefa 1")
        x += 1

def tarefa2() -> None:
    x = 0 
    while x < 10:
        print("Tarefa 2")
        x += 1
print(15*'-','SEM USAR THREADS',15*'-')
tarefa1()
tarefa2()


print(15*'-','COM THREADS',15*'-')
import threading

threading.Thread(target=tarefa1).start()
tarefa2()

