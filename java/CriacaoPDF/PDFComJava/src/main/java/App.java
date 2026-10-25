
import io.al3ncar.PDF.Vendas.Produto;
import io.al3ncar.PDF.Vendas.Venda;
import io.al3ncar.PDF.relatorios.Relatorio;
import io.al3ncar.PDF.relatorios.RelatorioSimples;

import java.util.ArrayList;
import java.util.List;



public class App {

    public static void main(String[] args) {
    
        List<Produto> produtos = new ArrayList<>();
        Venda venda = new Venda("André", produtos);

        venda.addProdutoAoCarrinho(new Produto("pinga", 2, 0.10));
        venda.addProdutoAoCarrinho(new Produto("pinga", 2, 0.10));
        venda.addProdutoAoCarrinho(new Produto("pinga", 2, 0.10));

        Relatorio relat = new RelatorioSimples(venda);
        relat.gerarCabecalho();
        relat.gerarCorpo();
        relat.gerarRodape();
        relat.imprimir();
    }
}
