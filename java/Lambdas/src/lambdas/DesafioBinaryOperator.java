package lambdas;


import java.text.DecimalFormat;
import java.util.function.Function;
import java.util.function.UnaryOperator;

public class DesafioBinaryOperator {

    public static void main(String[] args) {

        /*
        * 1. A partir do produto calcular o preço real (com desconto)
        * 2. Imposto Municipal: >=2500 (8,5%)/ < 2500 (Isento)
        * 3. Frete: >= 3000 (100)/ < 3000 (50)
        * 4. Arredondar: Deixar duas casas decimais
        * 5. Formatar: R$1234,56
        * */

        Produto p = new Produto("Ipad", 3235.89, 0.13);
        efetuarCompra(p);

    }

    public static double aplicarDesconto(Produto p){
        Function<Produto, Double> vFinal = prod -> prod.preco * (1 - prod.desconto);
        return vFinal.apply(p);
    }

    public static double aplicarImpostoMunicipal(double v){
        UnaryOperator<Double> impostoMunicipal = valor -> valor >= 2500? valor * 0.085 : valor;
        return v + impostoMunicipal.apply(v);
    }

    public static double calculaFrete(double v){
        UnaryOperator<Double> frete = valor -> valor >= 3000? valor + 100 : valor + 50;
        return frete.apply(v);
    }

    public static String formatarValor(double v){
        DecimalFormat real = new DecimalFormat("0.##");
        return real.format(v);
    }

    public static void efetuarCompra(Produto p){
        String valorProduto = formatarValor(calculaFrete(aplicarImpostoMunicipal(aplicarDesconto(p))));
        System.out.println("Parabéns pela aquisição. O valor à ser pago é: R$" + valorProduto);
    }
}
