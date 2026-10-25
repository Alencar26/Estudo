import java.util.Scanner;

public class MatrizAdjacente {

    public static void main(String[] args) {

        Scanner input = new Scanner(System.in);

        int [][] matriz = {{0,1,1,0,1},{1,0,1,1,0},{1,1,0,1,1},{0,1,1,0,0},{1,0,1,0,0}};
        printMatriz(matriz);

        System.out.println("\nVerificar conexão da cidade:");
        System.out.print("Cidade: ");
        conexoesDaCidade(input.nextInt(), matriz);
    }

    static void printMatriz(int[][] matriz) {
        System.out.print(" ");
        for(int i = 1; i <= matriz.length; i++) {
            System.out.print(" "+ i);
        }
        for (int i = 0; i < matriz.length; i++) {
            System.out.println();
            System.out.print(i+1+" ");
            for (int j = 0; j < matriz.length; j++) {
                System.out.print(matriz[i][j]+" ");
            }
        }
    }

    static void conexoesDaCidade(int cidade, int[][] matriz) {
        String conexoes = "";
        for (int i = 0; i < matriz.length; i++) {
            if (matriz[cidade-1][i] == 1) {
                conexoes += (i+1)+",";
            }
        }
        printConexoes(conexoes);
    }

    static void printConexoes(String conexoes) {
        System.out.println("A cidade informada faz conexão com:");
        String[] arrayConexoes = conexoes.split(",");
        for(String cidade : arrayConexoes) {
            System.out.println("Cidade: " + cidade);
        }
    }
}
