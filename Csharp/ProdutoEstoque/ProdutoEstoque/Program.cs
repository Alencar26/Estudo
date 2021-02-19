using System;
using System.Globalization;

namespace ProdutoEstoque
{
    public class Program
    {
        static void Main(string[] args)
        {
            Produto x;

            x = new Produto();


            Console.WriteLine("========== BEM VINDO AO PROGRAMA DE PRODUTO ESTOQUE ==========");

            Console.WriteLine("Entre com os dados do produto:");
            Console.Write("Nome: ");
            x.Nome = Console.ReadLine();
            Console.Write("Preço: ");
            x.Preco = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
            Console.Write("Quantidade: ");
            x.Quantidade = int.Parse(Console.ReadLine());

            Console.WriteLine(x.InfoProduto());
            
            bool continua = true;
            do
            {
                Console.WriteLine("\n\nMENU");
                Console.WriteLine("( 1 ) - Adicionar produto");
                Console.WriteLine("( 2 ) - Remover produto");
                Console.WriteLine("( 3 ) - Sair\n");
                Console.Write("Sua escolha: ");
                string opcao = Console.ReadLine();

                
                switch (opcao)
                {
                    case "1":
                        Console.WriteLine("Informe a quantidade que deseja adicionar:");
                        Console.Write("Informe: ");
                        x.AdicionarProdutos(int.Parse(Console.ReadLine()));
                        break;
                    case "2":
                        Console.WriteLine("Informe a quantidade que deseja remover:");
                        Console.Write("Informe: ");
                        x.RemoverProdutos(int.Parse(Console.ReadLine()));
                        break;
                    case "3":
                        continua = false;
                        break;
                }

            } while (continua);
        }
    }
}
