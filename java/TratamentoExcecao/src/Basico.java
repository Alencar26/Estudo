public class Basico {

    public static void main(String[] args) {

        Aluno a1 = null;

        try {
            imprimirAluno(a1);
        } catch (Exception e) {
            System.out.println("Ocorreu um erro AQUI: Instância de Aluno é null | " + e.getMessage());
        }

        //Exemplo de como usar try/catch...
    }

    public static class Aluno {
        Aluno() {
            String nome;
        }
    }

    public static void imprimirAluno(Aluno aluno) {
        System.out.println(aluno);
    }
}
