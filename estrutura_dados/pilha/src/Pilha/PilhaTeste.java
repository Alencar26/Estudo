package Pilha;

public class PilhaTeste {

    public static void main(String[] args){

        PilhaArray<String> array = new PilhaArray<String>();
        System.out.println("PilhaTeste");

        array.push("José da Silva");
        array.push("Mariano dos Santos");
        System.out.println("Tamanho da pilha: " + array.size());

        System.out.println("Mostra a pilha: " + array);

        try {
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
        } catch (PilhaNegativaException e) {
            System.out.println("Erro AQUI: " + e.getMessage());
        }
    }
}
