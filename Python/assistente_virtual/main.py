import os
import re
import speech_recognition as sr
import gtts
from playsound import playsound

__path__ = os.path.dirname(os.path.realpath(__file__))

microfone = sr.Recognizer()
while True:
    with sr.Microphone() as source:
        
        microfone.adjust_for_ambient_noise(source)
        print("Diga alguma coisa: ")
        audio = microfone.listen(source)
        try:
            frase = microfone.recognize_google(audio, language='pt-BR')
            
            with open(f'{__path__}/frase.txt', 'w') as arquivo:
                arquivo.write(format(frase))

            with open(f'{__path__}/frase.txt', 'w') as arquivo:
                for linha in arquivo:
                    frase = gtts.gTTS(linha, lang='pt-br')
                    frase.save('frase.mp3')
                    playsound('frase.mp3')
            
            if re.search(r'\b'+"ajudar"+r'\b',format(frase)):
                print("Voce falou algo sobre Ajudar?")
            elif re.search(r'\b'+"Meu nome é"+r'\b', format(frase)):
                temp = re.search('Meu nome é (.*)', format(frase))
                assert temp is not None #garantindo que a expressão não é nula
                print("Olá, " + temp.group(1))
            print("Você disse: " + frase)
        except sr.UnknownValueError:
            print("Não entendi")
            
            