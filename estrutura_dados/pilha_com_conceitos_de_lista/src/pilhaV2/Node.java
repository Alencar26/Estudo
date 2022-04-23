package pilhaV2;

public class Node<E> {

    private E elemento;
    private Node<E> prev;

    public Node() {
        this(null, null);
    }

    public Node(E elemento, Node<E> prev) {
        this.elemento = elemento;
        this.prev = prev;
    }

    public E getElemento() {
        return elemento;
    }

    public void setElemento(E elemento) {
        this.elemento = elemento;
    }

    public Node<E> getPrev() {
        return prev;
    }

    public void setPrev(Node<E> prev) {
        this.prev = prev;
    }

    @Override
    public String toString() {
        return " { " + this.elemento + " }";
    }
}
