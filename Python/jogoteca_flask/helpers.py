import os
from app import app

__img_path = app.config['UPLOAD_PATH']

def recuperar_imagem(id):
    for nome_arquivo in os.listdir(__img_path):
        if f'capa{id}.jpg' == nome_arquivo: 
            return nome_arquivo
    return 'default.png'

def salva_imagem(jogo, img):
    if img.filename != '':
        img.save(f'{__img_path}/capa{jogo.id}.jpg')
