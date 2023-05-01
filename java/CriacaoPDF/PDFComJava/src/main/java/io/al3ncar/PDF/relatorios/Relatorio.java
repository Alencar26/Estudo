package io.al3ncar.PDF.relatorios;

public interface Relatorio {
    public void gerarCabecalho();
    public void gerarCorpo();
    public void gerarRodape();
    public void imprimir();
}
