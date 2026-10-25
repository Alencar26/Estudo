import os
import shutil

extentions = ('png', 'jpg', 'jpeg', 'mp4', 'KVG', 'gif', 'MOV', 'MKV', 'WMV', 'BMP', 'pdf')
path = "C:/Users/andre/OneDrive/Área de Trabalho/celular_apagar_depois"
destination = "D:/pictures/celular-quebrado"

def getFiles(path):
    for root, subFolder, files in os.walk(path):
        for file in files:
            if file.lower().endswith(extentions):
                moveFile(destination ,file, root)  

def moveFile(destination, file, root):
    shutil.move(os.path.join(root,file), destination + '/' + file) 

getFiles(path)