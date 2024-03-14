from flask import render_template, request, redirect, session, url_for
from flask import flash

from app import app, db
from models import Jogos, Usuarios

@app.route('/')
def index():
    lista_jogos = Jogos.query.order_by(Jogos.id).all()
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
    
    jogo = Jogos.query.filter_by(nome=nome).first()
    if jogo:
        flash('Jogo já existente!')
        return redirect(url_for('novo_jogo'))
    novo_jogo = Jogos(nome=nome, categoria=categoria, console=console)
    db.session.add(novo_jogo)
    db.session.commit()
    
    return redirect(url_for('index'))

@app.route('/login')
def login():
    proxima = request.args.get('proxima')
    if proxima is None:
        proxima = url_for('index')
    return render_template('login.html', proxima=proxima)

@app.route('/autenticar', methods=['POST', ])
def auth():
    usuario = Usuarios.query.filter_by(nickname=request.form['usuario']).first() 
    if usuario:
        if request.form['senha'] == usuario.senha:
            session['usuario_logado'] = usuario.nickname
            flash(usuario.nickname + ' logado com sucesso!')
            proxima_pagina = request.form['redirect']
            return redirect(proxima_pagina)
        else:
            flash('Usuario ou senha invalidos!')
            return redirect(url_for('login', proxima=request.form['redirect']))
    else:
        flash('Usuario ou senha invalidos!')
        return redirect(url_for('login', proxima=request.form['redirect']))
    
@app.route('/logout')
def logout():
    session['usuario_logado'] = None
    flash('Logout efetuado com sucesso!')
    return redirect(url_for('login'))