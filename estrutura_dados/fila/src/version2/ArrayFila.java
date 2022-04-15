package version2;

import java.util.Arrays;

public class ArrayFila<E> implements Fila<E> {

    protected int capacidade;
    public static final int CAPACIDADE = 1000; //capacidade default
    protected E fila[];
    protected int posicaoPrimeiroDaFila;
    protected int posicaoLugarLivreNoFimDaFila;

    public ArrayFila() {
        this(CAPACIDADE);
    }

    public ArrayFila(int capacidade) {
        this.capacidade = capacidade;
        fila = (E[]) new Object[capacidade];
        posicaoPrimeiroDaFila = 0;
        posicaoLugarLivreNoFimDaFila = 0;
    }

    @Override
    public int tamanho() {
        int f = posicaoPrimeiroDaFila;
        int r = posicaoLugarLivreNoFimDaFila;

        return (this.capacidade - f + r) % this.capacidade;
    }

    @Override
    public boolean estaVazio() {
        return (posicaoPrimeiroDaFila == posicaoLugarLivreNoFimDaFila);
    }

    @Override
    public E primeiroDaFila() {
        return fila[posicaoPrimeiroDaFila];
    }

    @Override
    public void inserirNaFila(E elemento) {
        fila[posicaoLugarLivreNoFimDaFila] = elemento;
        posicaoLugarLivreNoFimDaFila = (posicaoLugarLivreNoFimDaFila + 1) % capacidade;
    }

    @Override
    public E tirarPrimeiroDaFila() {
        E temporario;
        temporario = fila[posicaoPrimeiroDaFila];
        fila[posicaoPrimeiroDaFila] = null;
        posicaoPrimeiroDaFila = (posicaoPrimeiroDaFila + 1) % capacidade;
        return temporario;
    }

    @Override
    public String toString() {
        return "ArrayFila{" +
                "capacidade=" + capacidade + "\n" +
                ", fila=" + Arrays.toString(fila) + "\n" +
                ", posicaoPrimeiroDaFila=" + posicaoPrimeiroDaFila + "\n" +
                ", posicaoLugarLivreNoFimDaFila=" + posicaoLugarLivreNoFimDaFila + "\n" +
                '}' + "\n";
    }
}
