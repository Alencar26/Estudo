public class OrdenacaoPorSelecaoString {
  
  public static void main(String[] args) {
    
    String[] v = {"Willian", "Pedro", "Ana", "Helo", "Zeze", "Felipe", "Maria"};
    ordenarPorSelecao(v);
    System.out.println("Ordenação por seleção:");
    for (int i = 0; i < v.length; i++) {
     System.out.println(v[i]); 
    }
  }

  public static int seleciona(String[] v, int n) {
   int  i = 0;
   for (int j = 1; j < n; j++) {
     if (v[i].compareTo(v[j]) < 0) {
       i = j;
     }
  }
  return i;
}

  public static void ordenarPorSelecao(String[] v) {
    int n = v.length;
    int i;
    String temporario;
    while (n > 1) {
      i = seleciona(v, n);
      temporario = v[n - 1];
      v[n - 1] = v[i];
      v[i] = temporario;
      n--;
    }
  }
}
