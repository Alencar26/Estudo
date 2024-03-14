from typing import Final

SECRET_KEY: Final = 'fanpsf39-11-fjefjewefnwpfweEF#%!#%!1'

SQLALCHEMY_DATABASE_URI: Final = \
    '{SGBD}://{USER}:{PASSWORD}@{SERVER}:{PORT}/{DATABASE}'.format(
        SGBD='mysql+mysqlconnector',
        USER='root',
        PASSWORD='mysql',
        SERVER='localhost',
        PORT='3307',
        DATABASE='jogoteca'
    )
