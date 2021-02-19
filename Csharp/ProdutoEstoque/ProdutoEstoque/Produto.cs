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
            Console.WriteLine($"Dados atualizados: {Nome}, {Preco.ToString("C")}, {Quantidade} unidades, Total: {ValorEmEstoque().ToString("C")} ");
        }

        public void RemoverProdutos(int quantidade)
        {
            Quantidade -= quantidade;
            Console.WriteLine($"Dados atualizados: {Nome}, {Preco.ToString("C")}, {Quantidade} unidades, Total: {ValorEmEstoque().ToString("C")} ");
        }

        public string InfoProduto()
        {
            return $"Dados do produto: {Nome}, {Preco.ToString("C")}, {Quantidade} unidades, Total: {ValorEmEstoque().ToString("C")}";
        }
    }
}
