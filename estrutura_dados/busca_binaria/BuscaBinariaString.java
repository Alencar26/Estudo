
public class BuscaBinariaString {
  
  public static void main(String[] args) {
    
    String[] vetor = {"André", "Helô", "Ana", "Lucas", "Bia", "Dora"};
    String chave = "Lucas";
    int retornoPosicao = 0;

    System.out.println("Teste de busca binária");
    System.out.println("Busca a chave: " + chave);
    
    retornoPosicao = buscaBinaria(chave, vetor);
    System.out.println("Resultado: " + retornoPosicao);
  }

  public static int buscaBinaria(String chave, String vetor[]) {

    int pontoInicial = 0;
    int tamanhoArray = vetor.length;
    int pontoFinal = tamanhoArray-1;

    int meio=0;
    while (pontoInicial <= pontoFinal) {
      meio = (pontoInicial + pontoFinal) / 2;
      if (chave.compareTo(vetor[meio]) == 0) {
       return meio;
      } else {
        if (chave.compareTo(vetor[meio]) < 0) {
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
