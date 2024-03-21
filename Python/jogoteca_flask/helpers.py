import os
from app import app

def recuperar_imagem(nome_game):
    for nome_arquivo in os.listdir(app.config['UPLOAD_PATH']):
        if f'{nome_game}.jpg' == nome_arquivo: 
            return nome_arquivo
    return 'default.png'