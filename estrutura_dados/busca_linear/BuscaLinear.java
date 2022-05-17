public class BuscaLinear {
  
  public static void main(String[] args) {
    
    int[] vetor = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13};
    int chave = 7;

    //Busca com inteiros
    int result = buscaLinearV1(chave, vetor);
    System.out.println("Resultado: " + result);

    int result2 = buscaLinearV2(chave, vetor);
    System.out.println("Resultado V2: " + result2);

    //Busca em String
    String[] vetorS = {"Ana", "Helo", "Dede", "Caio", "Lola", "Bia"};
    String chaveS = "Bia";

    int result3 = buscaLinearV3(chaveS, vetorS);
    System.out.println("Resultado Busca String: " + result3);
  }

  public static int buscaLinearV1(int chave, int[] valores) {
    int n = valores.length;
    for (int i = 0; i < n; i++) {
      if (chave == valores[i]) {
        return i;
      }
    }
    return -1;
  }

  public static int buscaLinearV2(int chave, int[] valores) {
    int n = 0;
    while (n != valores.length) {
      if (chave == valores[n]) return n;
      n++;
    }
    return -1;
  }

  public static int buscaLinearV3(String chave, String[] vetor) {
   int n = 0;
     while (n != vetor.length) {
     if (chave.compareTo(vetor[n]) == 0) return n;
     n++;
   }
   return -1;
  }
}
