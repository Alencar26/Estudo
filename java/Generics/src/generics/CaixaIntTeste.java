package generics;

public class CaixaIntTeste {

    public static void main(String[] args) {

        CaixaInt caixaA = new CaixaInt();
        caixaA.guardar(123); // resolveu a tipagem atravez de herança

        int coisaA = caixaA.abrir();
        System.out.println(coisaA);
    }
}
