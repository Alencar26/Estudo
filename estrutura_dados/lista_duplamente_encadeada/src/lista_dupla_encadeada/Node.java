package lista_dupla_encadeada;

public class Node<E> {

    private E elemento;
    private Node<E> proximo, anterior;

    Node() {
        this(null, null, null);
    }

    protected Node(E elemento, Node<E> anterior, Node<E> proximo ) {
        this.elemento = elemento;
        this.proximo = proximo;
        this.anterior = anterior;
    }

    public void setElemento(E novoElemento) {
        this.elemento = novoElemento;
    }

    public void setProximo(Node<E> novoProximo) {
        this.proximo = novoProximo;
    }

    public void setAnterior(Node<E> novoAnterior) {
        this.anterior = novoAnterior;
    }

    public E getElemento() {
        return elemento;
    }

    public Node<E> getProximo() {
        return proximo;
    }

    public Node<E> getAnterior() {
        return  anterior;
    }
}
