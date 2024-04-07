import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

#np.random.seed(7)
dados = np.random.randint(low=1,high=1500,size=10)
print(dados)
plt.plot(dados)
plt.plot(color="red",marker="o", ms = 5, mec="black", makerfacecolor='w' ,linestyle="-.",linewidth=2,alpha=0.5)
plt.title("Dados gerados")
plt.xlabel("Numero")
plt.ylabel("Valor")
plt.grid(axis="y", color="gray", linestyle="--", linewidth=0.5, alpha=0.5)
plt.show()