package net.datastructures;

import net.datastructures.Node;

/** Lista simplesmente encadeada */
public class ListaSimplesmenteEncadeada<E> {
    // nodo cabeça da lista
    protected Node head;
    protected long size;

    public ListaSimplesmenteEncadeada(){
        head = null;
        size = 0;
    }

    /**
     * Adiciona uma novo elemento na cabeça da lista, deslocando a cabeça atual;
     * @param elemento Elemento a ser inserido
     */
    public void addNewNode( E elemento){
        Node n  = new Node();
        n.setElement(elemento);
        n.setNext(this.head);
        this.head = n;
        this.size++;
    }

    /**
     * Remove o nó da cabeça da lista
     */
    public void removeNode(){
        Node n = this.head;
        this.head = this.head.getNext();
        n = null;
        this.size--;
    }

    public void removeAll(){
        Node n;

        while (this.head != null){
            this.removeNode();
        }

    }


    @Override
    public String toString() {
        String itemsLista = "";
        Node i = this.head;
        while (i != null) {
            itemsLista += i.getElement() + " \n";
            i = i.getNext();
        }
        return "ListaSimplesmenteEncadeada {" +
                "head=" + head +
                ", size=" + size +
                ",\nElementos da lista:  \n{\n" + itemsLista +
                '}';


    }



    public static void main(String[] args ){

    }

}




