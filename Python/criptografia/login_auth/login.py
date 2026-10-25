import hashlib
import os
import time
import json
from md5_bruteforce import brute_attack

class Usuario:
    
    def __init__(self, nome:str, email=None, senha=None) -> None:
        self.__nome: str = nome
        self.__email: str = email
        self.__senha: str = senha
        
    @property
    def nome(self) -> str:
         return self.__nome   
        
    @property
    def senha(self) -> str:
        return self.__senha
    
    @senha.setter
    def senha(self, senha: str) -> None:
        self.__senha = senha
        
    @property
    def email(self) -> str:
        return self.__email
    
    @email.setter
    def email(self, email: str) -> None:
        self.__email = email
        
    def to_dict(self) -> dict[str,str]:
        return {
            'nome': self.__nome,
            'email': self.__email,
            'senha': self.__senha
        }
    
    @classmethod
    def from_dict(cls, json:dict) -> 'Usuario':
        return cls(json['nome'], json['email'], json['senha'])
         
    def __str__(self) -> str:
        return f"NOME:{self.__nome}\nEMAIL:{self.__email}\nSENHA:{self.__senha}"
         
class Login:
    
    @classmethod
    def cadastrar(cls,usuario:Usuario) -> None:
        print("\n============= CADATRO ==============\n")
        email = cls().__obter_email()
        usuario.email = email
        
        senha = cls().__obter_senha()
        senha_hash = cls().__gerarHash(senha)
        usuario.senha = senha_hash
        
        cls().__salvar_em_arquivo(usuario)
        print("\n------ FIM DO CADASTRO ------\n")
    
    @classmethod
    def autenticar(cls,email:str, senha:str) -> None:
        print("\n============= LOGIN ==============\n")
        usuario:Usuario|None = cls().__buscar_usuario('database.txt', email)
        if cls().__validar_senha(usuario, senha):
            print(f"LOGADO COM SUCESSO\n{usuario}")
        else: 
            print("FALHA AO LOGAR (Login ou senha inválido)")
    
    @classmethod
    def bruteforce(cls) -> None:
        print(f"INICIO BRUTEFORCE: {time.strftime('%H:%M:%S')}")
        hash_list = cls().__buscar_senhas_hash("database.txt")
        brute_attack(hash_list)
        print(f"FIM BRUTEFORCE: {time.strftime('%H:%M:%S')}")
    
    def __gerarHash(self, senha:str) -> str:
        return hashlib.md5(senha.encode()).hexdigest()
    
    def __obter_email(self) -> str:
        email: str = input("Email (Até 80 caracteres): ")
        while(len(email) > 80):
            email = input("Email deve ter máximo de 80 caractere: ")
        return email
    
    def __obter_senha(self) -> str:
        senha: str = input("Senha (deve conter 4 caracteres): ")
        while(len(senha) != 4):
            senha = input("Senha (4 caracteres), informe novamente: ")
        return senha
    
    
    def __salvar_em_arquivo(self, usuario:Usuario) -> None:
        if not os.path.exists('database.txt'):
            os.mknod('database.txt')
        print("Salvando novo cadastro...")
        with open('database.txt', 'r+') as db:
            database = json.load(db)
            database.append(usuario.to_dict())
            db.seek(0)
            json.dump(database, db, indent=4)
        print("Cadastro registrado...")
    
    def __buscar_usuario(self, arquivo:str, email:str) -> Usuario | None:
        with open(arquivo, 'r') as db:
            database = json.load(db)
            for usuario in database:
                if usuario['email'] == email:
                    return Usuario.from_dict(usuario)
        return None
    
    def __buscar_senhas_hash(self, arquivo:str) -> list[str]:
        aux = 0
        hashList:list[str] = []
        with open(arquivo, 'r') as db:
            database = json.load(db)
            for usuario in database:
                if aux >= 4:
                    continue
                aux = aux + 1
                hashList.append(usuario['senha'])
        return hashList
    
    def __validar_senha(self, usuario:Usuario|None, senha:str) -> bool:
        if usuario is not None:
            senha_hash = self.__gerarHash(senha)
            return usuario.senha == senha_hash
        return False


rodando = True
while rodando:
    print("\n1 - CADASTRO\n2 - LOGIN\n3 - BRUTEFORCE\n4 - SAIR\n")
    op = int(input("R: "))
    match op:
        case 1:
            user = Usuario(input("\nQual seu nome: "))
            Login.cadastrar(user)
        case 2:
            Login.autenticar(input("Email: "), input("Senha: "))
        case 3:
            Login.bruteforce()
        case _:
            rodando = False
            
    

