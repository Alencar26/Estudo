public class PilhaArray<E> implements  Pilha<E> {

    protected int capacidade;
    public static final int CAPACIDADE = 1000;
    protected E array[];
    protected int top = -1;

    public PilhaArray() {
        this(CAPACIDADE);
    }

    public PilhaArray(int capacidade) {
        this.capacidade = capacidade;
        array = (E[]) new Object[capacidade];
    }

    @Override
    public int size() {
        return (top + 1);
    }

    @Override
    public boolean isEmpty() {
        return (top < 0);
    }

    @Override
    public E top() {
        return array[top];
    }

    @Override
    public void push(E element) {
        array[++top] = element;
    }

    @Override
    public E pop() throws PilhaNegativaException {
        E element;

        element = array[top];
        array[top--] = null;

        if (this.size() == 0)
            throw new PilhaNegativaException("A Pilha está VAZIA.");

        return element;
    }

    public void status(String op, Object element) {
        System.out.print("------> " + op);
        System.out.println(", returns " + element);
        System.out.print("result: size = " + size() + ", isEmpty = " +
                isEmpty());
        System.out.println(", stack: " + this);
        System.out.println("\n");
    }

    @Override
    public String toString() {
        StringBuilder s;
        s = new StringBuilder("[");
        if (size() > 0) s.append(array[0]);
        if (size() > 1)
            for (int i = 1; i <= size() - 1; i++) {
                s.append(", ").append(array[i]);
            }
        return s + "]";
    }

    public static void main(String[] args) {

        Object obj;
        PilhaArray<Integer> A = new PilhaArray<Integer>();
        try {
            A.status("new ArrayStack<Integer> A", null);
            A.push(7);
            A.status("A.push(7)", null);
            obj = A.pop();
            A.status("A.pop()", obj);
            A.push(9);
            A.status("A.push(9)", null);
            obj = A.pop();
            A.status("A.pop()", obj);
            PilhaArray<String> B = new PilhaArray<String>();
            B.status("new ArrayStack<String> B", null);
            B.push("João");
            B.status("B.push(\"João\")", null);
            B.push("Carlos");
            B.status("B.push(\"Carlos\")", null);
            obj = B.pop();
            B.status("B.pop()", obj);
            B.push("Sérgio");
            B.status("B.push(\"Sérgio\")", null);
        } catch (PilhaNegativaException e) {
            System.out.println(e.getMessage());
        }
    }
}