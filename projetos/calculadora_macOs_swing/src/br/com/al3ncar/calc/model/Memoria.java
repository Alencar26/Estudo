package br.com.al3ncar.calc.model;

public class Memoria {

    // USANDO PADRÃO DE PROJETO SINGLETON: só posso ter uma instância dessa classe.

    private static final Memoria instancia = new Memoria();
    private String textoAtual = "";

    private Memoria() {
    }

    public static Memoria getInstancia() {
        return instancia;
    }

    public String getTextoAtual() {
        return textoAtual.isEmpty() ? "0" : textoAtual;
    }
}
