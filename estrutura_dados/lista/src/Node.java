public class Node<E> {

   private E element;
   private Node<E> next;

   public Node() {
       this(null, null);
   }

   public Node(E element, Node<E> next) {
       this.element = element;
       this.next = next;
   }

   public E getElement() {
       return this.element;
   }

   public Node<E> getNext() {
       return this.next;
   }

   public void setElement(E newElement) {
       this.element = newElement;
   }

   public void setNext(Node<E> newNext) {
       this.next = newNext;
   }

    @Override
    public String toString() {
        return "Node { " + this.element + "}";
    }
}
