import time

print(30*'-')
print("Iniciou...")
time.sleep(5)
print("terminou")

print(30*'-')
agora = time.localtime()
print(agora)
print(time.strftime("%H:%M:%S", agora))
print(type(agora))

print(30*'-')
time_texto = '21 June, 2024'
time_obj = time.strptime(time_texto, '%d %B, %Y')
print(time_obj)
