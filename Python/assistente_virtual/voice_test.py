import pyttsx3

# Iniciando a biblioteca na variável engine
engine = pyttsx3.init()
engine.setProperty('voice', "brazil")
engine.setProperty('rate', 200)
engine.setProperty('valume', 1)
engine.say("Olá Ana paula... Minha voz é muito ruim. Nem parece português")

engine.runAndWait()