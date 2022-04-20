package generics.resolvendoTipaNaInstancia;

import generics.resolvendoTipaNaInstancia.Caixa;

public class CaixaTeste {

    public static void main(String[] args) {

        Caixa<String> caixaA = new Caixa<>();
        caixaA.guardar("Aqora só aceita String"); // agora ele não aceita outro tipo senão String

        String retornoSemCastParaString = caixaA.abrir();
        System.out.println(retornoSemCastParaString);
    }
}
