## descrição
'''
Quando nós queremos desenhar uma figura com objetivo de animação,
nós precisamos usar a função draw(desenhar em inglês).
'''
def setup():
    size(400,400)

def draw():
    frameRate(60) ## aqui eu determina quantos frames podem ser exibidos por segundo
    println(frameCount) ## está imprimindo no console os frames que estão passando em tela
## uma taxa legal e frames por segundo em jogos seria de mínimo 30, ideal 60 ou mais
    posicaoAleatoria()

def posicaoAleatoria():
    strokeWeight(10) ## grossura do ponto
    stroke(random(0,255),random(0,255),random(0,255)) ## cores aleatórias para os pontos
    point(random(0,width), random(0,height)) ##função random determina uma posição aleatória de 0 até largura e altura da tela
