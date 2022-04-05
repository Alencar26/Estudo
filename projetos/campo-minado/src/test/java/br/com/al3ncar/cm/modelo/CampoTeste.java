package br.com.al3ncar.cm.modelo;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

public class CampoTeste {

    private Campo campo;

    @Before
    public void iniciarCampo() {
        campo  = new Campo(3,3);
    }

    @Test
    public void testeVizinhoDistancia1Esquerda() {
        Campo vizinho = new Campo(3,2);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia1Direita() {
        Campo vizinho = new Campo(3,4);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia1Cima() {
        Campo vizinho = new Campo(2,3);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia1Baixo() {
        Campo vizinho = new Campo(4,3);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia2DiagonarCimaEsquerda() {
        Campo vizinho = new Campo(2,2);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia2DiagonarCimaDireita() {
        Campo vizinho = new Campo(2,4);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia2DiagonarBaixoEsquerda() {
        Campo vizinho = new Campo(4,2);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeVizinhoDistancia2DiagonarBaixoDireita() {
        Campo vizinho = new Campo(4,4);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertTrue(resultado);
    }

    @Test
    public void testeNaoVizinhoDistancia3OuMais() {
        Campo vizinho = new Campo(1,2);

        boolean resultado = campo.adicionarVizinho(vizinho);
        Assert.assertFalse(resultado);
    }

}
