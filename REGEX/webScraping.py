import requests
import re

url = "https://monicahillman.github.io/monibank/"
response = requests.get(url)
html = response.text

# [^>] --> classe Negando tudo que vem depois do <(h[1-6])
# (.*?) --> o sinal de '?' depois do '*' deixa o REGEX 'Lazy', ele só vai pegar a primeiro ocorrência
regex = r'<(h[1-6])[^>]*>(.*?)<\/(h[1-6])>'

# \1 ---> "backreference", para não repetir o pardrão como a regex de cima
# podemos usar o backreference para "espelhar" o primeiro padrão no final da regex
regex2 = r'<(h[1|2])[^>]*>(.*?)<\/\1>'

titulos = re.findall(regex2, html, re.DOTALL)
for titulo_html in titulos:
  print(titulo_html)