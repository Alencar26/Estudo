from flask import Flask, render_template, request, redirect, url_for, session, make_response

app = Flask(__name__)


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

app.secret_key = 'nfnfo32i208hgindepi2301fjpf13-3'

@app.route('/')
def index():
    return render_template('index.html', content = ['banana', 'maça', 'pera'])

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