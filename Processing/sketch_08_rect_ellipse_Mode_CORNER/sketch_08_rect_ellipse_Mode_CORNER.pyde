size(400,400)
stroke(41,113,125)
strokeWeight(5)
line(width/2, 0, width/2, height)
line(0, height/2, width, height/2)

fill(247,215,85)
rectMode(CORNER) ## centraliza no canto superior esquerdo
rect(width/2, height/2, 100, 100) ## por padrão o retanguro/quadrado começa das pontas

## criando um circulo
fill(101,182,196)
ellipseMode(CORNER) ## muda a forma de exibição padrão para ccentralizar do canto superior esquerdo
ellipse(width/2, height/2, 50,50) ## x,y,largura, altura
