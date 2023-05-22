import matplotlib.pyplot as plt
import numpy as np

print("\n========================================================================")
print("Olá, sejá bem-vindo(a) ao Circo dos Palhaços Pirotécnicos do Canhão!")
print("""Nossa próxima atração será um maravilhoso tiro de canhão.
Onde nosso querido palhaço Kama-e-Kase será lançado aos ares
\nSua 'função', entendeu 'função' rsrs, será informar a alturam máxima e 
a distância máxima que nosso palhacinho será lançado. Preparado?!""")
print("=========================================================================\n")

a = -2
b = 119
c = 0

# y = ax^2 + bx + c
def calcular_parabola(a, b, c, x):
    y = a * x**2 + b * x + c
    return y

def calcular_delta(a,b,c):
    delta = b**2 - 4 * a * c
    return delta

def calcular_distancia_max(b, delta, a):
    return (-b + delta ** 0.5) / (2 * a)

delta = calcular_delta(a,b,c);
altura_max = calcular_distancia_max(b, delta, a)
distancia_max = -b / (2 * a)

# distancia do eixo X
x = np.linspace(0, 80, 10)
y = calcular_parabola(a, b, c, x)

fig, ax = plt.subplots()
ax.plot(x, y)
plt.show()