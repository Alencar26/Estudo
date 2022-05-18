public class OrdenacaoPorSelecaoDecrescente {
  
  public static void main(String[] args) {
    
    int[] v = {2, 12, 96, 1, 99, 1006, 3, 17, 0, 100};
    ordenarPorSelecao(v);
    System.out.println("Ordenação por seleção:");
    for (int i = 0; i < v.length; i++) {
     System.out.println(v[i]); 
    }
  }

  public static int seleciona(int[] v, int n) {
   int  i = 0;
   for (int j = 1; j < n; j++) {
     if (v[i] > v[j]) { //Só mudar o sinal para deixar decrescente ou crescente.
       i = j;
     }
  }
  return i;
}

  public static void ordenarPorSelecao(int[] v) {
    int n = v.length;
    int i;
    int temporario;
    while (n > 1) {
      i = seleciona(v, n);
      temporario = v[n - 1];
      v[n - 1] = v[i];
      v[i] = temporario;
      n--;
    }
  }
}
