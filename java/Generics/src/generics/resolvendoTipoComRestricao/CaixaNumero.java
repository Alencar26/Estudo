package generics.resolvendoTipoComRestricao;

public class CaixaNumero<N extends Number> extends Caixa<N>{
    // Classe para resolver o tipo e incluir restrição a mais
    // N obrigatóriamente tem que extender da classe Number
}