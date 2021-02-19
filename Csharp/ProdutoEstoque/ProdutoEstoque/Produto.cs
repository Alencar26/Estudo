using System;

namespace ProdutoEstoque
{
    class Produto
    {
        public string Nome { get; set; }
        public double Preco { get; set; }
        public int Quantidade { get; set; }

        public double ValorEmEstoque()
        {
            return Quantidade * Preco;
        }

        public void AdicionarProdutos(int quantidade)
        {
            Quantidade += quantidade;
<<<<<<< HEAD
            Console.WriteLine($"Dados atualizados: {ToString()}");
=======
            Console.WriteLine($"Dados atualizados: {ToString()} ");
>>>>>>> dccf48fa9681015ea7015020acc14099773f2030
        }

        public void RemoverProdutos(int quantidade)
        {
            Quantidade -= quantidade;
<<<<<<< HEAD
            Console.WriteLine($"Dados atualizados: {ToString()}");
=======
            Console.WriteLine($"Dados atualizados: {ToString()} ");
>>>>>>> dccf48fa9681015ea7015020acc14099773f2030
        }

        public string InfoProduto()
        {
            return $"Dados do produto: {ToString()}";
        }
<<<<<<< HEAD

=======
        
>>>>>>> dccf48fa9681015ea7015020acc14099773f2030
        public override string ToString()
        {
            return $"{Nome}, {Preco.ToString("C")}, {Quantidade} unidades, Total: {ValorEmEstoque().ToString("C")}";
        }
    }
}
