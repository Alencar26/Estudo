using System;
using System.Globalization;

namespace Vetores
{
    class Program
    {
        static void Main(string[] args)
        {
            Quarto[] quartos = new Quarto[10];

            Console.WriteLine("Informe quantos quartos serão alugados: ");
            int nQ = int.Parse(Console.ReadLine());
            int qntQuartos = 10;
            for (int i = 0; i < nQ; i++)
            {
                
                Console.WriteLine($"Temos {qntQuartos} quartos disponíveis:");
                Console.WriteLine("Cadastro:");
                Console.Write("Nome :");
                string nome = Console.ReadLine();
                Console.Write("email: ");
                string email = Console.ReadLine();
                Console.Write("Numero do quarto: ");
                int q = int.Parse(Console.ReadLine());

                quartos[q] = new Quarto { Nome = nome, Email = email };
                qntQuartos -= 1;
            }

            for (int i = 0; i < quartos.Length; i++)
            {
                if (quartos[i] != null)
                    Console.WriteLine($"\n{i}: {quartos[i].Nome}, {quartos[i].Email}");
            }

            // =================================================================================
            int result = Calculator.Sum(new int[] { 2,3,5,2 });
            Console.WriteLine(result);
        }
    }
}
