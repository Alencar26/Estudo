

class DatabaseConnection():
    _instance = None
    conn = None

    def __new__(cls):
        if not cls._instance:
            cls._instance = super().__new__(cls)
            
            cls._instance.conn = "conectado!"
        return cls._instance

    def get_connection(self):
        return self.conn

conexao1 = DatabaseConnection()
print(conexao1.get_connection())

conexao2 = DatabaseConnection()
print(conexao2.get_connection())

print(conexao1 is conexao2)
