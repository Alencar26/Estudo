package br.com.al3ncar.cm.modelo;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Predicate;

public class Tabuleiro {

    private int linhas;
    private int colunas;
    private int minas;

    private final List<Campo> campos = new ArrayList<>();

    public Tabuleiro(int linhas, int colunas, int minas) {
        this.linhas = linhas;
        this.colunas = colunas;
        this.minas = minas;

        gerarCampos();
        associarVizinhos();
        sortearMinas();
    }

    private void gerarCampos() {
        for(int linha = 1; linha <= linhas; linha++) {
            for (int coluna = 1; coluna <= colunas; coluna++) {
                campos.add(new Campo(linha, coluna));
            }
        }
    }

    private void associarVizinhos() {
        for(Campo campo : campos) {
            for(Campo candidatoAVizinho : campos) {
                campo.adicionarVizinho(candidatoAVizinho);
            }
        }
    }

    private void sortearMinas() {
        long minasArmadas = 0;
        Predicate<Campo> minado = Campo::isMinado;

        do {
            minasArmadas = campos.stream().filter(minado).count();
                int campoAleatorio = (int) (Math.random() * campos.size());
            campos.get(campoAleatorio).minar();
        } while (minasArmadas < minas);
    }

    public void abrir(int linha, int coluna) {
        campos.parallelStream()
                .filter(campo -> campo.getLinha() == linha)
                .filter(campo -> campo.getColuna() == coluna)
                .findFirst()
                .ifPresent(Campo::abrir);
    }

    public void alternarMarcacao(int linha, int coluna) {
        campos.parallelStream()
                .filter(campo -> campo.getLinha() == linha)
                .filter(campo -> campo.getColuna() == coluna)
                .findFirst()
                .ifPresent(Campo::alternarMarcacao);
    }

    //método que diz se jogador ganhou.
    public boolean objetivoAlcancado() {
        return campos.stream().allMatch(Campo::objetivoAlcancado);
    }

    //reiniciar jogo
    public void reniciar() {
        campos.forEach(Campo::reiniciar);
        sortearMinas();
    }

    public String toString() {
        StringBuilder sb = new StringBuilder();

        int i = 0;
        for (int linha = 1; linha <= linhas; linha++) {
            for (int coluna = 1; coluna <= colunas; coluna++) {
                sb.append(" ");
                sb.append(campos.get(i));
                sb.append(" ");
                i++;
            }
            sb.append("\n");
        }

        return sb.toString();
    }
}
