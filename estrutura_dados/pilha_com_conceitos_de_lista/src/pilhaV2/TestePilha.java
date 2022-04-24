package pilhaV2;

public class TestePilha {

    public static void main(String[] args) {
        execute(new Pilha<String>());
    }

    public static void execute(Pilha<String> pilha) {
        populandoPilha(pilha);
        String resultado = !removendoUmAUm(pilha) ? "Pilha vazia." : "Removeu todos.";
        System.out.println(resultado);
    }

    public static void populandoPilha(Pilha<String> pilha) {
        pilha.push("Maria");
        pilha.push("Antonio");
        pilha.push("Unibrasil");
        pilha.push("Marcos");
        pilha.push("João");
        pilha.push("Karine");
        pilha.push("Karina");
        pilha.push("Luís");
        pilha.push("José");
        pilha.push("André Augusto Rolim de Alencar");
    }

    public static boolean removendoUmAUm(Pilha<String> pilha) {
        if (pilha.isEmpty()) return false;
        do {
            System.out.println("Retirando: "+ pilha.pop());
        } while (!pilha.isEmpty());
        return true;
    }
}
