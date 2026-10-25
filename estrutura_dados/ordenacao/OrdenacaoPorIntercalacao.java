public class OrdenacaoPorIntercalacao {
  
  public static void main(String[] args) {
    
    // Ordenação por Intercalação é muito mais rápido que ordenação por seleção ou outros mais simples.
     int[] vetor = {65, 23, 11, 8, 2,133};
     // é passado por parâmetro o vetor, a posição do primeiro elemento e a posição do último elemento.
     ordenarPorIntercalacao(vetor, 0, vetor.length - 1);
     for (int i = 0; i < vetor.length; i++) {
       System.out.println(vetor[i]);
     }
  }

  private static void intercala(int[] vetor, int primeiro, int meio, int ultimo) {
    
    int[] arrayComTamanhoDaMetade = new int[ultimo - primeiro + 1];
    int i = primeiro;
    int j = meio + 1;
    int k = 0;

    while (i <= meio && j <= ultimo) {
      if (vetor[i] < vetor[j]) {
        arrayComTamanhoDaMetade[k] = vetor[i];
        i++;
      } else {
        arrayComTamanhoDaMetade[k] = vetor[j];
        j++;
      }
      k++;
    }

    while (i <= meio) {
      arrayComTamanhoDaMetade[k] = vetor[i];
      k++;
      i++;
    }

    while (j <= ultimo) {
      arrayComTamanhoDaMetade[k] = vetor[j];
      k++;
      j++;
    }
    for (k = 0; k <= ultimo - primeiro; k++) {
      vetor[primeiro + k] = arrayComTamanhoDaMetade[k];
    }
  }

  static void ordenarPorIntercalacao(int[] vetor, int primeiro, int ultimo) {
    if (primeiro != ultimo) {
      int meio = (primeiro + ultimo) / 2;
      ordenarPorIntercalacao(vetor, primeiro, meio);
      ordenarPorIntercalacao(vetor, meio + 1, ultimo);
      intercala(vetor, primeiro, meio, ultimo);
    }
  }
}
