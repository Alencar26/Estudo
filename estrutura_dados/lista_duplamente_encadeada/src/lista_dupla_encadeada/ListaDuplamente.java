package lista_dupla_encadeada;

public class ListaDuplamente {

    protected int size;
    protected Node header, trailer;

    public ListaDuplamente() {
        size = 0;
        header = new Node(null, null, null);
        trailer = new Node(null, header, null);
        header.setProximo(trailer);
    }

    public int size() {
        return size;
    }

    public boolean estaVazio() {
        return (size == 0);
    }

    public Node getPrimeiro() throws IllegalAccessError {
        if (estaVazio()) throw new IllegalAccessError("Lista está vazia.");
        return header.getProximo();
    }

    public Node getUltimo() throws IllegalAccessError {
        if (estaVazio()) throw new IllegalAccessError("Lista está vazia.");
        return trailer.getAnterior();
    }

    public Node getAnterior(Node node) throws IllegalAccessError {
        if (node == header) throw new IllegalAccessError("Não é possível voltar alé do cabeçalho da lista.");
        return node.getAnterior();
    }

    public Node getProximo(Node node) throws IllegalAccessError {
        if (node == trailer) throw new IllegalAccessError("Não é possível avaçar além do final da lista.");
        return node.getProximo();
    }

    public void addAntes(Node nodeReferencia, Node novoNode) throws IllegalAccessError {
        Node nodeTemporario = getAnterior(nodeReferencia);

        novoNode.setAnterior(nodeTemporario);
        novoNode.setProximo(nodeReferencia);
        nodeReferencia.setAnterior(novoNode);
        nodeTemporario.setProximo(novoNode);
        size++;
    }

    public void addDepois(Node nodeReferencia, Node novoNode) throws IllegalAccessError {
        Node nodeTemporario = getProximo(nodeReferencia);

        novoNode.setAnterior(nodeReferencia);
        novoNode.setProximo(nodeTemporario);
        nodeTemporario.setAnterior(novoNode);
        nodeReferencia.setProximo(novoNode);
        size++;
    }

    public void addComoPrimeiro(Node novoNode) {
        addDepois(header, novoNode);
    }

    public void addComoUltimo(Node novoNode) {
        addAntes(trailer, novoNode);
    }

    public void remove(Node nodeDesejado) {
        Node nodeAnteriorTemporario = getAnterior(nodeDesejado);
        Node nodePosteriorTemporario = getProximo(nodeDesejado);

        nodePosteriorTemporario.setAnterior(nodeAnteriorTemporario);
        nodeAnteriorTemporario.setProximo(nodePosteriorTemporario);
        nodeDesejado.setAnterior(null);
        nodeDesejado.setProximo(null);
        size--;
    }

    @Override
    public String toString() {
        String retorno = "";
        Node node = this.getPrimeiro();
        String elementoDoNode = node.getElemento().toString();
        retorno = elementoDoNode + "\n";

        while (node != this.trailer) {
            node = node.getProximo();
            if ( node.getElemento() != null) {
                retorno += node.getElemento().toString();
                retorno += "\n";
            }
        }
        return retorno;
    }
}
