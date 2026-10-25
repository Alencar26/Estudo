package version1;

public interface Fila<E> {

    public int tamanho();
    public boolean estaVazio();
    public E primeiroDaFila();
    public void inserirNaFila(E elemento);
    public E tirarPrimeiroDaFila();
}
