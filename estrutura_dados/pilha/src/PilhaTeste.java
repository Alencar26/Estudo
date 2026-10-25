public class PilhaTeste {

    public static void main(String[] args){

        PilhaArray<String> array = new PilhaArray<String>();
        System.out.println("PilhaTeste");

        execute(array);
    }

    public static void execute(PilhaArray<String> array) {
        array.push("João");
        array.push("Maria");
        array.push("Antônio");
        array.push("Unibrasil");
        array.push("Marcos");
        array.push("Eugênio");
        array.push("Leão");
        array.push("Joaquim");
        array.push("José");
        array.push("André Augusto Rolim de Alencar");

        System.out.println("Tamanho da pilha: " + array.size());

        System.out.println("Mostra a pilha: " + array);

        try {
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
            array.pop();
            System.out.println("Mostra a pilha: " + array);
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
