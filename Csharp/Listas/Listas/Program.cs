using System;
using System.Collections.Generic;

namespace Listas
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> lista = new List<string>();

            lista.Add("Hello World");
            lista.Add("Hello World");
            lista.Add("Hello World");
            lista.Add("Hello World");
            lista.Add("Hello World");

            // diferença do insert é que ele aceita informar qual será a posição
            // que será inserido os dados.
            lista.Insert(3, "Olá Mundo");

            foreach (var item in lista)
            {
                Console.WriteLine(item);
            }

            Console.WriteLine($"Quantidade de itens na lista: {lista.Count}");

            // ==================== CRIANDO LISTA COM ELEMENTOS =======================

            List<string> lista2 = new List<string>() { "String1", "String2", "String3" };

            // ============================== FIND ==========================================

            string br = lista.Find(x => x[0] == 'O');
            Console.WriteLine($"String em pt-Br: {br}");

            // ============================= REMOVE =======================================

            // removendo pelo nome do elemento!
            lista.Remove("Olá Mundo");

            //removendo pela posição
            lista.RemoveAt(3);

            // remover elementos em uma faixa determinada
            lista.RemoveRange(2, 3); // << a partir da posição 2, quero remover 3 elementos

            foreach (var item in lista)
            {
                Console.WriteLine(item);
            }

        } 
    }
}
