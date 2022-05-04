package br.com.al3ncar.calc.model;

import java.util.ArrayList;
import java.util.List;

public class Memoria {

    private enum TipoComando {
            ZERAR, NUMERO, DIV, MULT, SUB, SOMA, IGUAL, VIRGULA;
    };

    // USANDO PADRÃO DE PROJETO SINGLETON: só posso ter uma instância dessa classe.

    private static final Memoria instancia = new Memoria();
    private final List<MemoriaObserver> observadores = new ArrayList<>();
    private String textoAtual = "";

    private Memoria() {
    }

    public void adicionarObservadores(MemoriaObserver observador) {
        observadores.add(observador);
    }

    public static Memoria getInstancia() {
        return instancia;
    }

    public String getTextoAtual() {
        return textoAtual.isEmpty() ? "0" : textoAtual;
    }

    public void processarComando(String texto) {

        TipoComando tipoComando = detectarTipoComando(texto);
        System.out.println(tipoComando);
        if ("AC".equals(texto)){
            textoAtual = "";
        } else {
            textoAtual += texto;
        }
        observadores.forEach(obs -> obs.valorAlterado(getTextoAtual()));
    }

    private TipoComando detectarTipoComando(String texto) {
        if (textoAtual.isEmpty() && texto.equals("0")) {
            return null;
        }

        try {
            Integer.parseInt(texto);
            return TipoComando.NUMERO;
        } catch (NumberFormatException e) {
            // Quando não for número
            if ("AC".equals(texto)) {
                return TipoComando.ZERAR;
            } else if ("/".equals(texto)) {
                return  TipoComando.DIV;
            } else if ("+".equals(texto)) {
                return  TipoComando.SOMA;
            } else if ("-".equals(texto)) {
                return  TipoComando.SUB;
            } else if ("x".equals(texto)) {
                return  TipoComando.MULT;
            } else if ("=".equals(texto)) {
                return  TipoComando.IGUAL;
            } else if (",".equals(texto)) {
                return  TipoComando.VIRGULA;
            }
        }
        return null;
    }
}
