import requests
import re

url = "https://monicahillman.github.io/monibank/"
response = requests.get(url)
html_content = response.text

# Verifique se a requisição foi bem-sucedida
if response.status_code == 200:
    # Conteúdo da página
    pagina = response.text

    # \b (word boundary) --> cria um limite naocorrência e busca isoladamente
    #  e só vai buscar a palavra que estiver dentro de \bpalavra\b
    
    # \B (Non-Word Boundary/not-a-word boundary)
    
    # Padrao regex para encontrar linhas com a palavra "moni"
    padrao = r'\bbank\b'

    # Divida o conteúdo da página em linhas
    linhas = pagina.split('\n')

    for indice, linha in enumerate(linhas, start=1):
        if re.search(padrao, linha, re.IGNORECASE):
            print(f"Linha {indice}: {linha}")
else:
    print(f"A requisição GET para {url} falhou com o código {response.status_code}")