package generics.resolvendoTipoComCasting;

import generics.resolvendoTipoComCasting.CaixaObjeto;

public class CaixaObjetoTeste {

    public static void main(String[] args) {

        CaixaObjeto caixaA = new CaixaObjeto();
        caixaA.guardar(2.3); //double -> Double

        // o codigo abaixo vai dar problema pois o java espera um número inteiro para ser atribuido na variável
        Double coisaA = (Double) caixaA.abrir(); //casting
        System.out.println(coisaA);

        CaixaObjeto caixaB = new CaixaObjeto();
        caixaB.guardar("Olá, sou uma String");

        String coisaB = (String) caixaB.abrir(); //casting
        System.out.println(coisaB);
    }
}
