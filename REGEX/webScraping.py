import requests
import re

url = "https://monicahillman.github.io/monibank/"
response = requests.get(url)
html = response.text

# [^>] --> classe Negando do sinal de '>' REGEX não pega esse caracter
# (.*?) --> o sinal de '?' depois do '*' deixa o REGEX 'Lazy', ele só vai pegar a primeiro ocorrência
regex = r'<(h[1-6])[^>]*>(.*?)<\/(h[1-6])>'

titulos = re.findall(regex, html, re.DOTALL)
for titulo_html in titulos:
  print(titulo_html)