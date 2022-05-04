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

    private TipoComando ultimaOperação = null;
    private boolean substituir = false;
    private String textoAtual = "";
    private String textoBuffer = "";

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

        if (tipoComando == null) {
            return;
        } else if (tipoComando == TipoComando.ZERAR) {
            textoAtual = "";
            textoBuffer = "";
            substituir = false;
            ultimaOperação = null;
        } else if (tipoComando == TipoComando.NUMERO
                || tipoComando == TipoComando.VIRGULA) {
            textoAtual = substituir ? texto : textoAtual + texto;
            substituir = false;
        } else {
            //operações
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
            } else if (",".equals(texto) && !textoAtual.contains(",")) {
                return  TipoComando.VIRGULA;
            }
        }
        return null;
    }
}
