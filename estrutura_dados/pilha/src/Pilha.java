public interface Pilha<E> {

    public int size();
    public boolean isEmpty();
    public E top();
    public void push(E element);
    public E pop() throws PilhaNegativaException;
}