using System;
using System.Linq;
using System.Collections.Generic;

namespace LinqExample
{
    class Program
    {
        static void Main(string[] args)
        {
            //my data source
            int[] numbers = new int[] { 2, 3, 4, 5 };

            //my query expression
            IEnumerable<int> result = numbers
                .Where(x => x % 2 == 0)
                .Select(x => x * 10);
             
            // execute the query
            foreach (var n in result)
            {
                Console.WriteLine(n);
            }

            //===================== OPERAÇÕES ======================

            //OPERAÇÕES DE FILTRO
            numbers.Where(x => x == 2);
            numbers.OfType<int>();

            //OPERAÇÕES PARA ORDENAR
            numbers.OrderBy(x => x.Equals(x % 2 == 0));
            numbers.OrderByDescending(x => x.Equals(x % 2 == 0));

            numbers.OrderBy(x => x.Equals(x % 2 == 0)).ThenBy(x => x < 0);
            numbers.OrderBy(x => x.Equals(x % 2 == 0)).ThenByDescending(x => x < 0);
            numbers.Reverse();

            //OPERAÇÕES DE CONJUNTO
            numbers.Distinct();

            int[] numbers2 = new int[] { 2, 3, 4 };
            
            numbers.Except(numbers2);
            numbers.Intersect(numbers2);
            numbers.Union(numbers2);

            //OPERAÇÕES DE QUANTIFICAÇÃO (TODOS ELEMENTOS)

            numbers.All(x => x > 2);
            numbers.Any(x => x > 2);
            numbers.Contains(2);

            // OPERAÇÕES DE PROJEÇÃO

            numbers.Select(x => x < 4);

            int[][] arrays = {
                                new[] {1,2,3},
                                new[] {4,5},
                                new[] {6},
                                new[] {7,8,9}
                                };

            numbers = arrays.SelectMany(x => x).ToArray(); // ou
            var resultado = arrays.SelectMany(x => x); // ou
            IEnumerable<int> resultado2 = arrays.SelectMany(x => x);

            //OPERAÇÕES PARTITION (PAGINAÇÃO)

            numbers.Skip(3); // pule tantos elementos
            numbers.Take(5); // pegue tantos elementos

            // OPERAÇÕES JOIN  (JUNÇÃO)

            var innerJoin = numbers.Join(numbers2, n => n, n2 => n2, (n, n2) => n);

            var innerJoin2 = numbers.GroupJoin(numbers2, n => n, n2 => n2, (n, n2) => n);

            //OPERAÇAO DE AGRUPAMENTO

            var result2 = numbers.GroupBy(x => x % 2 == 0);

            //OPERAÇÕES QUE GERA UMA FONTE DE DADOS VAZIA

            var database = Enumerable.Empty<string>();// ou
            var database2 = Enumerable.Empty<string>().ToList();

            //OPERAÇÃO DE IGUALDADE

            numbers.SequenceEqual(numbers2);

            // OPERAÇÕES PARA OBTER ELEMENTOS

            numbers.ElementAt(0); // pegar pelo index
            numbers.First();// pegar o primeiro
            numbers.FirstOrDefault(); // primeiro por padrão
            numbers.Last(); //último
            numbers.LastOrDefault(); // ultimo por padrão
            numbers.Single();
            numbers.SingleOrDefault();

            // OPERAÇÕES DE CONVERÇÃO

            numbers.AsEnumerable();
            numbers.AsQueryable();

            //OPERAÇÃO DE CONCATENAÇÃO

            numbers.Concat(numbers2);

            // OPERAÇÕES DE AGREGAÇÃO

            string[] vect = new string[] { "A", "B", "C", "D" };
            vect.Aggregate((s, s2) => s + ", " + s2);
            // result: A, B, C, D <--- colocou uma virgula no meio.

            var media = numbers.Average(); // pega a média dos valores
            int total = numbers.Count();
            long total2 = numbers.LongCount();
            int maiorValor = numbers.Max();
            int menorValor = numbers.Min();
            int soma = numbers.Sum();
        }
    }
}
