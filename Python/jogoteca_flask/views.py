from flask import render_template, request, redirect, session, url_for, send_from_directory
from flask import flash

from app import app 
from ext.database import db
from models import Jogos, Usuarios
from helpers import recuperar_imagem, salva_imagem
from ext.forms import FormularioJogo, FormularioLogin

@app.route('/')
def index():
    logado = False
    if session['usuario_logado'] is not None:
        logado = True
    lista_jogos = Jogos.query.order_by(Jogos.id).all()
    return render_template('lista.html', titulo='Jogos', jogos=lista_jogos, logado=logado)

@app.route('/jogo/novo')
def novo_jogo():
    if 'usuario_logado' not in session or session['usuario_logado'] is None:
        return redirect(url_for('login', proxima=url_for('novo_jogo')))
    form = FormularioJogo()
    return render_template('novo-jogo.html', titulo='Novo Jogo', form=form)

@app.route('/jogo/editar/<int:id>')
def editar_jogo(id):
    if 'usuario_logado' not in session or session['usuario_logado'] is None:
        return redirect(url_for('login', proxima=url_for('editar_jogo', id=id)))
    
    jogo = Jogos.query.filter_by(id=id).first()
    capa_jogo = recuperar_imagem(id=jogo.id)
    
    form = FormularioJogo()
    form.nome.data = jogo.nome
    form.categoria.data = jogo.categoria
    form.console.data = jogo.console
    
    if jogo:
        return render_template('editar-jogo.html', titulo='Editar Jogo', jogo=jogo, capa_jogo=capa_jogo, form=form)
    else:
        return redirect(url_for('index'))

@app.route('/jogo/atualizar', methods=['POST',])
def atualizar():
    
    form = FormularioJogo()
    
    if form.validate_on_submit(): 
        id = request.form['id']
        jogo = Jogos.query.filter_by(id=id).first()
        jogo.nome = form.nome.data
        jogo.categoria = form.categoria.data
        jogo.console = form.console.data

        db.session.add(jogo)
        db.session.commit()
        
        salva_imagem(jogo=jogo, img=request.files['arquivo'])

    return redirect(url_for('index'))   

@app.route('/jogo/deletar/<int:id>')
def deletar_jogo(id):
    if 'usuario_logado' not in session or session['usuario_logado'] is None:
        return redirect(url_for('login'))
    
    #outra opção
    #Jogos.query.filter_by(id=id).delete()
    #db.session.commit()
    
    jogo = Jogos.query.filter_by(id=id).first()
    if jogo:
        db.session.delete(jogo)
        db.session.commit()
        flash('Jogo deletado com sucesso!')
        return redirect(url_for('index'))
    else:
        flash('Jogo não encontrado!')
#        return redirect(url_for('index')) 

@app.route('/jogo/criar', methods=['POST',])
def criar():
    
    form = FormularioJogo(request.form)
    
    if not form.validate_on_submit():
        return redirect(url_for('novo_jogo'))
    
    nome = form.nome.data
    categoria = form.categoria.data
    console = form.console.data
    
    jogo = Jogos.query.filter_by(nome=nome).first()
    if jogo:
        flash('Jogo já existente!')
        return redirect(url_for('novo_jogo'))
    novo_jogo = Jogos(nome=nome, categoria=categoria, console=console)
    db.session.add(novo_jogo)
    db.session.commit()
    
    jogo = Jogos.query.filter_by(nome=nome).first()
    salva_imagem(jogo=jogo, img=request.files['arquivo'])
    
    return redirect(url_for('index'))

@app.route('/login')
def login():
    form = FormularioLogin()
    proxima = request.args.get('proxima')
    if proxima is None:
        proxima = url_for('index')
    return render_template('login.html', proxima=proxima, form=form)

@app.route('/autenticar', methods=['POST', ])
def auth():
    form = FormularioLogin(request.form)
    usuario = Usuarios.query.filter_by(nickname=form.nickname.data).first() 
    if usuario:
        if form.senha.data == usuario.senha:
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

@app.route('/static/img/<nome_arquivo>')
def imagem(nome_arquivo):
    return send_from_directory('static/img', nome_arquivo)