from flask_sqlalchemy import SQLAlchemy

db: SQLAlchemy = SQLAlchemy()

def init_app(app):
    db.init_app(app)