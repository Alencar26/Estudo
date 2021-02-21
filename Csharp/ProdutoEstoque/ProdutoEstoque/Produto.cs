using System;

namespace ProdutoEstoque
{
    class Produto
    {
        public string Nome { get; set; }
        public double Preco { get; set; }
        public int Quantidade { get; set; }


        public Produto(string nome, double preco, int quantidade) : this(nome, preco)
        {
            Quantidade = quantidade;
        }

        public Produto(string nome, double preco)
        {
            Nome = nome;
            Preco = preco;  
        }

        public double ValorEmEstoque()
        {
            return Quantidade * Preco;
        }

        public void AdicionarProdutos(int quantidade)
        {
            Quantidade += quantidade;

            Console.WriteLine($"Dados atualizados: {ToString()}");

        }

        public void RemoverProdutos(int quantidade)
        {
            Quantidade -= quantidade;

            Console.WriteLine($"Dados atualizados: {ToString()}");

        }

        public string InfoProduto()
        {
            return $"Dados do produto: {ToString()}";
        }

        public override string ToString()
        {
            return $"{Nome}, {Preco.ToString("C")}, {Quantidade} unidades, Total: {ValorEmEstoque().ToString("C")}";
        }
    }
}
