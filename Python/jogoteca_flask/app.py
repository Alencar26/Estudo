from flask import Flask
from ext import database
from ext import forms

app: Flask = Flask(__name__)
app.config.from_pyfile('config.py')

database.init_app(app)
forms.init_app(app)

from views_game import *
from views_user import *

if __name__ == '__main__': 
    app.run(host='0.0.0.0', port=5000, debug=True)