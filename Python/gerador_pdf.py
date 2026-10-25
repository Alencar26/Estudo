from reportlab.lib.pagesizes import letter
from reportlab.pdfgen import canvas

pdf_files = "gerador_pdf.pdf"
c = canvas.Canvas(pdf_files, pagesize=letter)
c.drawString(100, 100, "Hello World")
c.save()
print("PDF criado pelo reportlab")
