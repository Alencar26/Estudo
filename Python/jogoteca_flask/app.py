from flask import Flask, render_template, request, redirect, session, url_for
from flask import flash
from flask_sqlalchemy import SQLAlchemy

class Jogo:
    def __init__(self, nome, categoria, console):
        self.nome = nome
        self.categoria = categoria
        self.console = console

jogo1 = Jogo('Super Mario', 'Ação', 'SNES')
jogo2 = Jogo('Pokemon Gold', 'RPG', 'GBA')
jogo3 = Jogo('Mortal Kombat', 'Luta', 'SNES')
lista_jogos = [jogo1,jogo2,jogo3]

class Usuario:
    def __init__(self, nome, nickname, senha):
        self.nome = nome
        self.nickname = nickname
        self.senha = senha

usuario1 = Usuario('Fernando', 'fernando', '123456')
usuario2 = Usuario('Maria', 'maria', 'abcdef')
usuario3 = Usuario('João', 'joao', 'qwerty')

usuarios = {
            usuario1.nickname: usuario1,
            usuario2.nickname: usuario2,
            usuario3.nickname: usuario3
            }

app = Flask(__name__)
app.secret_key = 'fanpsf39-11-fjefjewefnwpfweEF#%!#%!1'

app.config['SQLALCHEMY_DATABASE_URI'] = \
    '{SGBD}://{USER}:{PASSWORD}@{SERVER}:{PORT}/{DATABASE}'.format(
        SGBD='mysql+mysqlconnector',
        USER='root',
        PASSWORD='mysql',
        SERVER='localhost',
        PORT='3307',
        DATABASE='jogoteca'
    )

#instância banco de dados
db = SQLAlchemy(app)

class Jogos(db.Model):
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    nome = db.Column(db.String(50), nullable=False)
    categoria = db.Column(db.String(40), nullable=False)
    console = db.Column(db.String(20), nullable=False)
    
    def __repr__(self):
        return '<Name %r>' % self.name
    
class Usuarios(db.Model):
    nickname = db.Column(db.String(8), primary_key=True)
    nome = db.Column(db.String(20), nullable=False)
    senha = db.Column(db.String(100), nullable=False)
    
    def __repr__(self):
        return '<Name %r>' % self.name

@app.route('/')
def index():
    return render_template('lista.html', titulo='Jogos', jogos=lista_jogos)

@app.route('/novo-jogo')
def novo_jogo():
    if 'usuario_logado' not in session or session['usuario_logado'] is None:
        return redirect(url_for('login', proxima=url_for('novo_jogo')))
    return render_template('novo-jogo.html', titulo='Novo Jogo')

@app.route('/criar', methods=['POST',])
def criar():
    nome = request.form['nome']
    categoria = request.form['categoria']
    console = request.form['console']
    jogo = Jogo(nome, categoria, console)
    
    lista_jogos.append(jogo)
    return redirect(url_for('index'))

@app.route('/login')
def login():
    proxima = request.args.get('proxima')
    if proxima is None:
        proxima = url_for('index')
    return render_template('login.html', proxima=proxima)

@app.route('/autenticar', methods=['POST', ])
def auth():
    
    if request.form['usuario'] in usuarios:
        usuario = usuarios[request.form['usuario']]
        if request.form['senha'] == usuario.senha:
            session['usuario_logado'] = usuario.nickname
            flash(usuario.nickname + ' logado com sucesso!')
            proxima_pagina = request.form['redirect']
            return redirect(proxima_pagina)
    else:
        flash('Usuario ou senha invalidos!')
        return redirect(url_for('login', proxima=request.form['redirect']))
    
@app.route('/logout')
def logout():
    session['usuario_logado'] = None
    flash('Logout efetuado com sucesso!')
    return redirect(url_for('login'))

app.run(host='0.0.0.0', port=5000, debug=True)