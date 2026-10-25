import os
from app import app
from time import time
from fnmatch import filter

__img_path = app.config['UPLOAD_PATH']

def recuperar_imagem(id):
    for nome_arquivo in os.listdir(__img_path):
        if f'capa{id}' in nome_arquivo: 
            return nome_arquivo
    return 'default.png'

def salva_imagem(jogo, img):
    if img.filename != '':
        if recuperar_imagem(jogo.id) == 'default.png':
            img.save(f'{__img_path}/capa{jogo.id}-{time()}.jpg')
        else:
            capas = filter(os.listdir(__img_path), f'capa{jogo.id}*')
            for capa in capas:
                os.remove(f'{__img_path}/{capa}')
            img.save(f'{__img_path}/capa{jogo.id}-{time()}.jpg')