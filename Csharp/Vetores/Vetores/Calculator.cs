using System;
using System.Collections.Generic;
using System.Text;

namespace Vetores
{
    class Calculator
    {
        // ao adicionar a palavra "params" não precisa mais instanciar um array ao chamar o método.
        public static int Sum(params int[] numbers)
        {
            int sum = 0;
            for (int i = 0; i < numbers.Length; i++)
            {
                sum += numbers[i];
            }
            return sum;
        }
        public static void preencher_matriz(int[,] mat)
        {
            for (int i = 0; i < mat.GetLength(0); i++)
            {
                for (int j = 0; j < mat.GetLength(1); j++)
                {
                    Console.WriteLine("Entre com os valores da matriz:");
                    Console.Write($"Posição {i}.{j}: ");
                    mat[i, j] = int.Parse(Console.ReadLine());
                }
            }
        }
        public static void diagonal(int[,] mat, int tamanho)
        {
            int negativos = 0;
            int[] arr = new int[tamanho];

            for (int i = 0; i < tamanho; i++)
            {
                for (int j = 0; j < tamanho; j++)
                {
                    if (mat[i, j] < 0)
                        negativos += 1;

                    if (i == j)
                        arr[i] = mat[i, j];
                }
            }

            Console.WriteLine("Diagonal: ");
            foreach (var num in arr)
            {
                Console.Write(num + " ");
            }

            Console.WriteLine($"\nNúmeros negativos: {negativos}");
        }
    }
}
