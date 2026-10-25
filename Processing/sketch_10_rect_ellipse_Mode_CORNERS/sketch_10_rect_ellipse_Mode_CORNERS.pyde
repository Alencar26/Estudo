size(400,400)
background(255,255,255) ##corde de fundo da tela
stroke(41,113,125)
strokeWeight(5)

fill(247,215,85)
rectMode(CORNERS) ## Esse modo é totalmente diferente do demais 
rect(0,0,400,400) ## ao inves de ser (x,y,largura,altura)
                                  ## vai ser (começa do X, começa do y, vai até X, vai até y)
                                  ## a firua se adapta ao tamanho informado.

fill(101,182,196)
ellipseMode(CORNERS)
ellipse(width/2, 0, width, height) ## o mesmo vale para a ellipse
