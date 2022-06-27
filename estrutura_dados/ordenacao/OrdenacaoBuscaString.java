public class OrdenacaoBuscaString {

    public static void main(String[] args) { executarTeste(); }

    //ordenação por intercalação ==========================================================
    static void ordenacaoPorIntercalacao(String[] vetor, int primeiro, int ultimo) {

        if (primeiro != ultimo) {

            int meio = (primeiro + ultimo) / 2;
            ordenacaoPorIntercalacao(vetor, primeiro, meio);
            ordenacaoPorIntercalacao(vetor, meio + 1, ultimo);

            //INTERCALAÇÃO

            String[] arrayComTamanhoDaMetade = new String[ultimo - primeiro + 1];
            int i = primeiro;
            int j = meio + 1;
            int k = 0;

            while (i <= meio && j <= ultimo) {
                if (vetor[i].compareTo(vetor[j]) < 0) {
                    arrayComTamanhoDaMetade[k] = vetor[i];
                    i++;
                } else {
                    arrayComTamanhoDaMetade[k] = vetor[j];
                    j++;
                }
                k++;
            }

            //sobrou só primera metade
            while (i <= meio) {
                arrayComTamanhoDaMetade[k] = vetor[i];
                k++;
                i++;
            }

            //sobrou só segunda metade
            while (j <= ultimo) {
                arrayComTamanhoDaMetade[k] = vetor[j];
                k++;
                j++;
            }

            for (k = 0; k <= ultimo - primeiro; k++) {
                vetor[primeiro + k] = arrayComTamanhoDaMetade[k];
            }
        }
    }

    //BUSCA BINÁRIA ===================================================================
    static int buscaBinaria(String chave, String[] vetor) {

        int pontoInicial = 0;
        int tamanhoArray = vetor.length;
        int pontoFinal = tamanhoArray - 1;

        int meio = 0;
        while (pontoInicial <= pontoFinal) {
            meio = (pontoInicial + pontoFinal) / 2;

            if (chave.compareTo(vetor[meio]) == 0) {
                return meio;
            } else {
                if (chave.compareTo(vetor[meio]) < 0) {
                    pontoFinal = meio - 1;
                } else {
                    pontoInicial = meio + 1;
                }
            }
        }
        return -1;
    }

    // UTILITÁRIOS =====================================================================
    static void printVetor(String[] v) {

        for (int i = 0; i < v.length; i++) {
            if(i == 0) System.out.print("[ ");
            if(i == (v.length - 1)) {
                System.out.print(v[i] + " ]\n");
                break;
            }
            System.out.print(v[i] + ", ");
        }
    }

    static void resultadoBuscaBinaria(int result, String[] vetor) {
        if (result == -1) {
            System.out.println("\nResultado: Não encontrado.");
        } else {
            System.out.println("\nResultado: Encontrado.");
            System.out.println("Valor: " + vetor[result] + " - Na posição: " + result);
        }
    }

    static void executarTeste() {
        // criando vetor desordenado.
        String[] v = {"str5", "str2", "str7", "str1", "str3", "str4", "str8", "str6", "str0", "str9"};

        // mostrando vetor desordenado.
        System.out.print("Vetor original: ");
        printVetor(v);

        // ordenando e mostrando vetor.
        ordenacaoPorIntercalacao(v, 0, v.length - 1);
        System.out.print("Vetor ordenado: ");
        printVetor(v);

        // efetuando busca por valor pré-definido no vetor.
        resultadoBuscaBinaria(buscaBinaria("str7", v), v);

        // efetuando busca por valor inexistente no vetor.
        resultadoBuscaBinaria(buscaBinaria("str11", v), v);
    }
}
