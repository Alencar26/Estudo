public class Sorting {
  
  public static void main(String[] args) {
    
    int[] v = {2, 45, 23, 1, 166, 3, 55, 1003, 30669, 18, 8, 319};
    ordenar(v);
    System.out.println("Ordenação Crescente:");
    for (int i = 0; i < v.length; i++) {
      System.out.println(v[i]);
    }

    ordenarDesc(v);
    System.out.println("\nOrdenação Decrescente:");
    for (int i = 0; i < v.length; i++) {
      System.out.println(v[i]);
    }

  }

  public static void ordenar(int[] vetor) {
    
    int n = vetor.length;
    int temporario;
    for (int i = 0; i < n - 1; i++) {
      for (int j = 0; j < n -1; j++) {
        if (vetor[j] > vetor[j + 1]) {
         temporario = vetor[j + 1];
         vetor[j + 1] = vetor[j];
         vetor[j] = temporario;
        }
      }
    }
  }

  public static void ordenarDesc(int[] vetor) {
    
    int n = vetor.length;
    int temporario;
    for (int i = 0; i < n - 1; i++) {
      for (int j = 0; j < n -1; j++) {
        if (vetor[j] < vetor[j + 1]) {
         temporario = vetor[j + 1];
         vetor[j + 1] = vetor[j];
         vetor[j] = temporario;
        }
      }
    }
  }
}

