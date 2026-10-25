package pilhaV2;

public interface IPilha<E> {

    public long size();
    public boolean isEmpty();
    public E top();
    public void push(E elemento);
    public E pop();
}
