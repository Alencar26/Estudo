package io.al3ncar.PDF.relatorios;

import java.io.FileOutputStream;

import com.lowagie.text.Chunk;
import com.lowagie.text.Document;
import com.lowagie.text.Element;
import com.lowagie.text.Font;
import com.lowagie.text.Paragraph;
import com.lowagie.text.pdf.PdfWriter;

import io.al3ncar.PDF.Vendas.Produto;
import io.al3ncar.PDF.Vendas.Venda;

public class RelatorioSimples implements Relatorio{

    private Venda venda;
    private Document documentPDF;
    private String caminhoRelatorio = "relatorioPDF.pdf";

    public RelatorioSimples(Venda venda) {
        this.venda = venda;
        this.documentPDF = new Document();
        try {
            PdfWriter.getInstance(this.documentPDF, new FileOutputStream(caminhoRelatorio));
            this.documentPDF.open();
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }

    @Override
    public void gerarCabecalho() {
        Paragraph paragrafoTitulo = new Paragraph();
        paragrafoTitulo.setAlignment(Element.ALIGN_CENTER);
        //vamos colocar um bloco de texto dentro do parágrafo
        //com fonte x e tamanho x
        paragrafoTitulo.add(
            new Chunk(
                "Relatório de vendas simples", 
                new Font(Font.HELVETICA, 24)
        ));

        this.documentPDF.add(paragrafoTitulo);
        this.documentPDF.add(new Paragraph(" ")); //pula linha

        Paragraph paragrafoData = new Paragraph();
        paragrafoData.setAlignment(Element.ALIGN_CENTER);
        paragrafoData.add(new Chunk(this.venda.getDataVenda().toString()));

        this.documentPDF.add(paragrafoData);
        this.documentPDF.add(new Paragraph(" ")); //pula linha
        this.documentPDF.add(new Paragraph(" ")); //pula linha

        Paragraph paragrafoCliente = new Paragraph();
        paragrafoCliente.setAlignment(Element.ALIGN_CENTER);
        paragrafoCliente.add(
            new Chunk(
                "Cliente: " + this.venda.getNomeCliente(),
                new Font(Font.BOLD, 16)
            )
        );
        this.documentPDF.add(paragrafoCliente);
        
        Paragraph paragrafroSessao = new Paragraph("---------------------------------------------------------------");
        paragrafroSessao.setAlignment(Element.ALIGN_CENTER);
        this.documentPDF.add(paragrafroSessao);
        this.documentPDF.add(new Paragraph(" ")); //pula linha
    }

    @Override
    public void gerarCorpo() {
        Paragraph paragItensVendidos = new Paragraph();
        paragItensVendidos.setAlignment(Element.ALIGN_CENTER);
        paragItensVendidos.add(new Chunk( "ITENS VENDIDOS: ", new Font(Font.TIMES_ROMAN, 16))); 
        this.documentPDF.add(paragItensVendidos);

        for (Produto produto : this.venda.getProdutosVendidos()) {
            Paragraph pNomeProduto = new Paragraph();
            pNomeProduto.add(new Chunk(produto.getNome(), new Font(Font.COURIER, 14)));
            Paragraph pDadosProduto = new Paragraph();
            pDadosProduto.add(new Chunk("Quantidade vendida: " + produto.getQuantidade() + 
                                        "Valor unitário: " + produto.getValor() + 
                                        "Total: R$ " + produto.calcularPreco()));

            this.documentPDF.add(pNomeProduto);
            this.documentPDF.add(pDadosProduto);
            this.documentPDF.add(new Paragraph("----------------")); //pula linha
        }
        

        Paragraph pTotal = new Paragraph();
        pTotal.setAlignment(Element.ALIGN_RIGHT);
        pTotal.add(new Chunk("Total: R$ " + this.venda.calcularValorTotalCarrinho(), new Font(Font.TIMES_ROMAN, 20)));
        this.documentPDF.add(pTotal);
    }

    @Override
    public void gerarRodape() {
        Paragraph pulaLinParagraph = new Paragraph();
        pulaLinParagraph.add(new Chunk(" "));
        this.documentPDF.add(pulaLinParagraph);
        Paragraph paragrafoRodape = new Paragraph();
        paragrafoRodape.setAlignment(Element.ALIGN_CENTER);
        paragrafoRodape.add(new Chunk("------------------------------"));
        this.documentPDF.add(paragrafoRodape);
        Paragraph pRodape = new Paragraph();
        pRodape.setAlignment(Element.ALIGN_CENTER);
        pRodape.add(new Chunk("www.al3ncar.com.br", new Font(Font.TIMES_ROMAN, 14)));
        this.documentPDF.add(pRodape);
    }

    @Override
    public void imprimir() {
        if(this.documentPDF != null && this.documentPDF.isOpen()) {
            this.documentPDF.close();
        }
    }
}
