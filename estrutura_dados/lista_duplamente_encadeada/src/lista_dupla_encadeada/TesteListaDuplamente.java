package lista_dupla_encadeada;


public class TesteListaDuplamente {

    public static void main(String[] args) {
        execute();
    }

    public static void execute() {
        ListaDuplamente lista = new ListaDuplamente();
        Node node, node1, node2;

        node1 = new Node();
        node1.setElemento("José");

        lista.addComoPrimeiro(node1);

        node2 = new Node();
        node2.setElemento("João Carlos");
        lista.addDepois(node1, node2);

        node = new Node();
        node.setElemento("Mais um nome incluído");
        lista.addDepois(node2, node);

        System.out.println("+-----------------------------------+ ");
        System.out.println("Lista : ");
        System.out.println(lista);
        System.out.println("+-----------------------------------+ ");

        lista.remove(node2);

        System.out.println("+-----------------------------------+ ");
        System.out.println("Lista : ");
        System.out.println(lista);
        System.out.println("+-----------------------------------+ ");
    }
}
