package br.com.al3ncar.calc.model;

import java.util.ArrayList;
import java.util.List;

public class Memoria {

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

    public void processarComando(String valor) {
        if ("AC".equals(valor)){
            textoAtual = "";
        } else {
            textoAtual += valor;
        }
        observadores.forEach(obs -> obs.valorAlterado(getTextoAtual()));
    }
}
