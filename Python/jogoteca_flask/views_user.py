from flask import render_template, request, redirect, session, url_for
from flask import flash

from app import app 
from models import Usuarios
from ext.forms import FormularioLogin

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
