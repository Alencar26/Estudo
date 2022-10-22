import java.util.*;

public class OrdenacaoNatural {

    public static void main(String[] args) {
        Set<Serie> seriesDesordenadas = new HashSet<>(){{
            add(new Serie("Spy Family", "Anime", 20));
            add(new Serie("Aneis de Poder", "Aventura", 120));
            add(new Serie("B99", "Comedia", 25));
        }};
        System.out.println("---Lista Desordenada---");
        printList(seriesDesordenadas);

        Set<Serie> seriesOrdenadas = new TreeSet<>(new ComparatorNomeGeneroTempoDuracao());
        seriesOrdenadas.addAll(seriesDesordenadas);

        System.out.println("\n---Lista Ordenada com estilo---");
        printList(seriesOrdenadas);
    }

    protected static void printList(Set<Serie> listaSeries) {
        for(Serie serie : listaSeries)
            System.out.println(
                    "Série: " + serie.getNome() +
                    " - " + serie.getGenero() +
                    " - " + serie.getTempoDuracao()+"|min/ep");
    }
}

class ComparatorNomeGeneroTempoDuracao implements Comparator<Serie> {

    @Override
    public int compare(Serie serie1, Serie serie2) {
        int nome = serie1.getNome().compareTo(serie2.getNome());
        if(nome != 0) return nome;

        int genero = serie1.getGenero().compareTo(serie2.getGenero());
        if(genero != 0) return nome;

        return Integer.compare(serie1.getTempoDuracao(), serie2.getTempoDuracao());
    }
}

class Serie {
    private String nome;
    private String genero;
    private int tempoDuracao;

    public Serie(String nome, String genero, int tempoDuracao) {
        this.nome = nome;
        this.genero = genero;
        this.tempoDuracao = tempoDuracao;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public String getGenero() {
        return genero;
    }

    public void setGenero(String genero) {
        this.genero = genero;
    }

    public int getTempoDuracao() {
        return tempoDuracao;
    }

    public void setTempoDuracao(int tempoDuracao) {
        this.tempoDuracao = tempoDuracao;
    }
}
