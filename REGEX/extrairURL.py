import re

# Texto de exemplo que contém URLs
texto_html = """..
  Visite o nosso site em https://www.exemplo.com.
  Para mais informações, acesse http://www.outroexemplo.com/pagina.
  Conheça nosso patrocinador oficial https://www.sitedasuaempresa.com.br.
"""

# Expressão Regular para encontrar URLs
padrao = r'https?://[^\s]+'

# Usando a função findall() para extrair as URLs
urls_encontradas = re.findall(padrao, texto_html)

# Imprime as URLs encontradas
if urls_encontradas:
    print("URLs encontradas:")
    for index, url in enumerate(urls_encontradas, start=1):
        print(f"{index}: {url}")
else:
    print("Nenhuma URL encontrada.")
