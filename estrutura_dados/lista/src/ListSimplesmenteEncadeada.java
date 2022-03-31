import java.util.List;

public class ListSimplesmenteEncadeada<E> {

    protected Node head;
    protected long size;

    public ListSimplesmenteEncadeada() {
        this.head = null;
        this.size = 0;
    }

    // Adicionar novo elemento na lista
    public void addNewNode(E elemento) {
        Node n = new Node();
        n.setElement(elemento);
        n.setNext(this.head);
        this.head = n;
        this.size++;
    }

    // remover o nó da cabeça da lista
    public void removeNode() {
        Node n = this.head;
        this.head = this.head.getNext();
        n = null;
        this.size--;
    }

    public void removeAll() {
        Node n;

        while (this.head != null) {
            this.removeNode();
        }
    }

    @Override
    public String toString() {
        String itensLista = "";
        Node i = this.head;
        while (i != null) {
            itensLista += i.getElement() + "\n";
            i = i.getNext();
        }
        return "ListaSimplesmenteEncadeada {" +
                "head = " + this.head +
                ", size = " + this.size +
                ", \nElementos da lista: \n{\n" + itensLista +
                "}";
    }

    public static void main(String[] args) {

    }
}
