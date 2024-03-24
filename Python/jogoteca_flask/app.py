from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from ext import forms

app: Flask = Flask(__name__)
app.config.from_pyfile('config.py')

#instância banco de dados
db: SQLAlchemy = SQLAlchemy(app)
forms.init_app(app)

from views import *

if __name__ == '__main__': 
    app.run(host='0.0.0.0', port=5000, debug=True)