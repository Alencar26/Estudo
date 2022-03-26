public class ChecadaVsNaoChecada {

    public static void main(String[] args) {
        try {
            geraErro1();
        } catch(RuntimeException e) {
            System.out.println(e.getMessage());
        }

        try {
            geraErro2();
        } catch(Exception e) {
            System.out.println(e.getMessage());
        }

        System.out.println("FIM.");
    }
    /*
     * Não chegada (não verificada) <- você escolhe se quer ou não, tratar esse erro.
     * Qualquer classe que herdar de RuntimeException não tem a obrigatoriedade de tratar a exceção.
     */
    // Não verificada ou não checada
    static void geraErro1(){
        throw new RuntimeException("Ocorreu um erro AQUI: #01!");
    }

    //Checada ou verificada <- você é obrigado a tratar essa exceção
    static void geraErro2() throws Exception {
      throw new Exception("Ocorreu um erro AQUI: #02!");
    }
}
