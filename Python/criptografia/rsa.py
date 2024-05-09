
import rsa
import time

texto: str = "RSA: algoritmo dos professores do MIT: Rivest, Shamir e Adleman"
bits_rsa = [1024, 2048, 4096, 8192]


def cripto_RSA(texto:str, bits:list[int]):
    for bit in bits:
        
        print(30*'-',f'CRIPTOGRAFIA RSA com {bit} bits', 30*'-')
        pubkey, privkey = rsa.newkeys(bit)
        
        start_time = time.time()
        enc_texto = rsa.encrypt(texto.encode(), pubkey) #encriptar texto
        end_time = time.time()
        enc_time = end_time - start_time
 
        start_time = time.time()
        dec_texto = rsa.decrypt(enc_texto, privkey).decode() #decriptar
        end_time = time.time()
        dec_time = end_time - start_time

        print(f"""Mensagem encriptada: {enc_texto}, tempo: {enc_time}s\nMensagem decriptada: {dec_texto}, tempo: {dec_time}s""")
        input("press ENTER para seguir.")

cripto_RSA(texto, bits_rsa)