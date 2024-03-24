from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField, validators
from flask_wtf.csrf import CSRFProtect

def init_app(app):
    CSRFProtect(app)

class FormularioJogo(FlaskForm):
    nome = StringField('Nome do Jogo', [validators.DataRequired(), validators.Length(min=1, max=50)])
    categoria = StringField('Categoria', [validators.DataRequired(), validators.Length(min=1, max=40)])
    console = StringField('Console', [validators.DataRequired(), validators.Length(min=1, max=20)])
    salvar = SubmitField('Salvar')