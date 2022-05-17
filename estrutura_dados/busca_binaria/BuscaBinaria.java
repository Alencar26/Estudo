
public class BuscaBinaria {
  
  public static void main(String[] args) {
    
    int[] vetor = {15, 22, 46, 71, 120, 125, 171, 201, 367, 1243};
    int chave = 171;
    int r = 0;

    System.out.println("Teste de busca binária");
    System.out.println("Busca a chave: " + chave);
    
    r = buscaBinaria(chave, vetor);
    System.out.println("Resultado: " + r);
  }

  public static int buscaBinaria(int chave, int vetor[]) {

    int pontoInicial = 0;
    int tamanhoArray = vetor.length;
    int pontoFinal = tamanhoArray-1;

    int meio=0;
    while (pontoInicial <= pontoFinal) {
      meio = (pontoInicial + pontoFinal) / 2;
      if (chave == vetor[meio]) {
       return meio;
      } else {
        if (chave < vetor[meio]) {
          pontoFinal = meio -1; 
        }
        else {
          pontoInicial = meio +1;
        }
      } // else
    }
    return -1;
  }
}
