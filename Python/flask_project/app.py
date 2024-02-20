import pymysql
from typing import Callable, Final
from flask import Flask, render_template, request, redirect, url_for, session, make_response


app = Flask(__name__)

db = pymysql.connect(
    host='172.17.0.2',
    user='root',
    password='root',
    database='flask_python')

pessoa_var: dict = {
    'nome': 'João',
    'idade': 20,
    'cidade': 'São Paulo',
    'profissao': 'Programador'
}

conteudo_pre_definido: str = """"
<!DOCTYPE html>
<html>
<head>
	<title>
		How to specify the media type
		of the media resource in HTML5?
	</title>
</head>

<body style="text-align: center;">
	<h1 style="color:green;">
		GeeksforGeeks
	</h1>
	<h3>
		How to specify the media type
		of the media resource in HTML5?
	</h3>
	<video width="400" height="350" controls>
		<source src=
"https://media.geeksforgeeks.org/wp-content/uploads/output-1.mp4"
				type="video/mp4">
	</video>
	</br></br>
	<audio controls>
		<source src=
"https://media.geeksforgeeks.org/wp-content/uploads/20190625153922/frog.mp3"
				type="audio/mp3">
	</audio>
</body>
</html>
"""

#session
app.secret_key = 'nfnfo32i208hgindepi2301fjpf13-3'

@app.route('/')
def index(): 
    cursor = db.cursor()
    sql = "SELECT * FROM clientes"
    cursor.execute(sql)
    result = cursor.fetchall()
    
    for i in range(len(result)):
        (id, nome, idade) = result[i]
        print(f'{id} - {nome} - {idade}')
        
    return render_template('index.html', content = result)

@app.route('/deletar', methods=['POST', 'GET'])
def deletar() -> Callable:
    id = request.args.get('id')
    cursor = db.cursor()
    SQL: Final = "DELETE FROM clientes WHERE id = %s"
    cursor.execute(SQL, id)
    db.commit()
    return redirect(url_for('index'))

@app.route('/', methods=['POST', 'GET'])
def update() -> Callable:
    if request.method == 'POST':
        id = request.form.get('id')
        nome = request.form.get('nome')
        idade = request.form.get('idade')
        print(id, nome, idade)
        cursor = db.cursor()
        SQL: Final = "UPDATE clientes SET nome = %s, idade = %s WHERE id = %s"
        cursor.execute(SQL, (nome, idade, id))
        db.commit()
    return redirect(url_for('index'))

@app.route('/pessoa')
def pessoa():
    if request.method == 'GET':
        if request.cookies.get('nome'):
            resp = make_response(render_template('pessoa.html', content = pessoa_var.items()))
        else:
            resp = make_response(render_template('pessoa.html', content = pessoa_var.items()))
            resp.set_cookie('nome', 'Andre')
        return resp
        
    else: 
        return '404 - NOT FOUND'
@app.route('/form', methods=['GET', 'POST'])
def form() -> str:
    if request.method == 'GET':
        if 'usuario' in session:
            usuario = session['usuario']
            return usuario
        else:
            session['usuario'] = 'André'
            return session['usuario']
        #return render_template('form.html', content = conteudo_pre_definido)
    else:
        return request.form['conteudo']

@app.route('/noticia/<slug>')
def noticia(slug: str) -> str:
    return f'<h1>Noticia {slug}</h1>'