import smtplib
import ssl
import email.message

msg = email.message.Message()
msg["Subject"] = "Assunto (Título)"

nome = "André"

body = f"""
<h1>Aqui começa o corpo do Email</h1>
<p>Olá, {nome}. Tudo bem?</p>
<p>Segue em anexo o documento solicitado</p>
"""

msg["From"] = "meu-email-exemplo@gmail.com"
msg["To"] = "email-destino@email.com"
senha = "senha"
msg.add_header("Content-Type", "text/html")
msg.set_payload(body)

contexto = ssl.create_default_context() #criação de contexto de segurança
with smtplib.SMTP('smtp.gmail.com', 587) as conexao: #conexão com o servidor de email (abre e fecha)
    conexao.ehlo() #verificar conexão
    conexao.starttls(context=contexto) #iniciar conexão segura
    conexao.login(msg["From"], senha)
    conexao.sendmail(msg["From"], msg["To"], msg.as_string().encode("utf-8")) #enviar email
