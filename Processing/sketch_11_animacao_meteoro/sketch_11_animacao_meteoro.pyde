def setup():
    size(500,500)
    noStroke()
    

def draw():
    frameRate(60)
    fill(47,95,211) ## cor do céu
    rect(0,0,width, height) ## retando para fazer o céu
    
    fill(255,191,0) ## cor do meteoro
    ellipse(frameCount, frameCount, 30,30) ## meteoro
    
    fill(250,233,73) ##cor das estrelas
    ellipse(random(0,width), random(0,height), 5,5) ## estrelas
    
    fill(57,53,20) ## cor do chão
    rect(0, 420, width, height) ## criando o chão
    
