package pilhaV2;

public class Pilha<E> implements IPilha<E> {

    protected Node<E> topo;
    protected long size;

    public Pilha() {
        this.topo = null;
        this.size = 0;
    }

    @Override
    public long size() {
        return this.size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public E top() {
        return topo.getElemento();
    }

    @Override
    public void push(E elemento) {
        Node<E> n = new Node<>();
        n.setElemento(elemento);
        n.setPrev(this.topo);
        this.topo = n;
        this.size++;
    }

    @Override
    public E pop() {
        Node<E> n = this.topo;
        E elemnto = n.getElemento();

        this.topo = this.topo.getPrev();
        n = null;
        this.size--;
        return elemnto;
    }

    @Override
    public String toString() {
        String itensPilha = "";
        Node<E> i = this.topo;
        while(i != null) {
          itensPilha += i.getElemento() + "\n";
          i = i.getPrev();
        }

        return "Pilha {" +
                "Topo: " +
                this.topo +
                ", tamanho: " +
                this.size +
                ",\nElementos da pilha: \n{\n" +
                itensPilha +
                "}";
    }
}
