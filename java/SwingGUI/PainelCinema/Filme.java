package PainelCinema;

public class Filme {
    
    private String titulo;
    private String genero;
    private String imagem;

    public Filme(String titulo, String genero, String imagem) {
        this.titulo = titulo;
        this.genero = genero;
        this.imagem = imagem;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getGenero() {
        return genero;
    }

    public String getImagem() {
        return imagem;
    }
}
