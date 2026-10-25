public class OrdenacaoPorInsercao {
  
  public static void main(String[] args) {
    
   int[] vetor = {2, 12, 3, 8, 16, 11, 1, 60};
  ordenacaoPorInsercao(vetor, vetor.length);

   for (int itemDoVetor : vetor) {
     System.out.println(itemDoVetor);
   }
  }

  static void insere(int itemAtual, int[] vetor, int posicaoAtual) {
    while (posicaoAtual > 0 && itemAtual < vetor[posicaoAtual - 1]) {
      vetor[posicaoAtual] = vetor[posicaoAtual - 1];
      posicaoAtual--;
    }
    vetor[posicaoAtual] = itemAtual;
  }

  static void ordenacaoPorInsercao(int[] vetor, int tamanho) {
    for (int i = 0; i < tamanho; i++) {
      insere(vetor[i], vetor, i);
    }
  }
}
