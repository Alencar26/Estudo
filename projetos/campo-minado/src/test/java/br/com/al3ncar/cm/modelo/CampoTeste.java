package br.com.al3ncar.cm.modelo;

import br.com.al3ncar.cm.excecao.ExplosaoException;
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

    @Test
    public void testeValorPadraoAtributoMarcado() {
        Assert.assertFalse(campo.isMarcado());
    }

    @Test
    public void testeAlternarMarcacao() {
        campo.alternarMarcacao();
        Assert.assertTrue(campo.isMarcado());
    }

    @Test
    public void testeAlternarMarcacaoDuasChamadas() {
        campo.alternarMarcacao();
        campo.alternarMarcacao();
        Assert.assertFalse(campo.isMarcado());
    }

    @Test
    public void testeAbrirCampoNaoMinadoNaoMarcado() {
        Assert.assertTrue(campo.abrir());
    }

    @Test
    public void testeAbrirCampoNaoMinadoMarcado() {
        campo.alternarMarcacao();
        Assert.assertFalse(campo.abrir());
    }

    @Test
    public void testeAbrirCampoMinadoMarcado() {
        campo.minar();
        campo.alternarMarcacao();
        Assert.assertFalse(campo.abrir());
    }

    @Test
    public void testeAbrirCampoMinadoNaoMarcado() {
        campo.minar();
        Assert.assertThrows(ExplosaoException.class, () -> {
            campo.abrir();
        });
    }

    @Test
    public void testeMinarCampo() {
        Assert.assertTrue(campo.minar());
    }

    @Test
    public void testeMinarCampoJaMinado() {
        campo.minar();
        Assert.assertFalse(campo.minar());
    }

    @Test
    public void testeAbrirCampoComVizinhos() {
        Campo vizinho1 = new Campo(2,2);
        Campo vizinho2 = new Campo(2,3);
        Campo vizinho3 = new Campo(2,4);

        campo.adicionarVizinho(vizinho1);
        campo.adicionarVizinho(vizinho2);
        campo.adicionarVizinho(vizinho3);

        Assert.assertTrue(campo.abrir());
    }

    @Test
    public void testeCamposVizinhosAbertos() {
        Campo vizinho1 = new Campo(2,2);
        Campo vizinho2 = new Campo(2,3);
        Campo vizinho3 = new Campo(2,4);

        campo.adicionarVizinho(vizinho1);
        campo.adicionarVizinho(vizinho2);
        campo.adicionarVizinho(vizinho3);
        campo.abrir();

        Assert.assertTrue(vizinho1.isAberto() && vizinho2.isAberto() && vizinho3.isAberto());
    }

    @Test
    public void testeAbrirVizinhoDoVizinhoDoCampoQueFoiAberto() {
        Campo vizinhoDoMeuVizinho = new Campo(1,1);

        Campo meuVizinho = new Campo(2,2);
        meuVizinho.adicionarVizinho(vizinhoDoMeuVizinho);

        campo.adicionarVizinho(meuVizinho);
        campo.abrir();

        Assert.assertTrue(meuVizinho.isAberto() && vizinhoDoMeuVizinho.isAberto());
    }

    @Test
    public void testePararPropagacaoDeCamposComVizinhoDoMeuVizinhoMinado() {
        Campo vizinhoDoMeuVizinho = new Campo(1,1);
        vizinhoDoMeuVizinho.minar();

        Campo meuVizinho = new Campo(2,2);
        meuVizinho.adicionarVizinho(vizinhoDoMeuVizinho);

        campo.adicionarVizinho(meuVizinho);
        campo.abrir();

        Assert.assertTrue(meuVizinho.isAberto() && vizinhoDoMeuVizinho.isFechado());
    }

}
