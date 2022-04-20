package generics.resolvendoTipoComRestricao;

public class CaixaNumeroTeste {

    public static void main(String[] args) {

        //Agora só pode usar o tipo

        CaixaNumero<Double> caixaA = new CaixaNumero<>();
        caixaA.guardar(2.3);

        System.out.println(caixaA.abrir());
    }
}
