using System;
using System.Collections.Generic;
using DemoLINQExample.Entities;
using System.Linq;

namespace DemoLINQExample
{
    class Program
    {
        static void Print<T>(string message, IEnumerable<T> collection)
        {
            Console.WriteLine(message);
            foreach (T obj in collection)
            {
                Console.WriteLine(obj);
            }
            Console.Write("\n");
        }

        static void Main(string[] args)
        {
            Category category1 = new Category(1, "Tools", 2);
            Category category2 = new Category(2, "Computers", 1);
            Category category3 = new Category(3, "Electronics", 1);

            List<Product> products = new List<Product>();

            products.Add(new Product(1, "Computer", 1100.0, category2));
            products.Add(new Product(2, "Hammer", 90, category1));
            products.Add(new Product(3, "TV", 1700, category3));
            products.Add(new Product(4, "Tablet", 700, category3));
            products.Add(new Product(5, "MacBook", 1800, category2));
            products.Add(new Product(6, "Level", 70, category1));
            products.Add(new Product(7, "Saw", 80, category1));


            var r1 = products.Where(p => p.Price < 900 && p.Category.Tier == 1);
            Print("Tier 1 and Price < 900:", r1);

            var r2 = products.Where(p => p.Category.Name == "Tools").Select(p => p.Name);
            Print("Names of Products from Tools", r2);

            // USANDO OBJETO ANONIMO NA EXPRESSÃO LAMBDA
            var r3 = products.Where(p => p.Name.StartsWith("C")).Select(p => new { p.Name, p.Price, AliasCategoryName = p.Category.Name });
            Print("Products starts with C: ", r3);

            // pegando produtosque tenha o tier 1 e ordenando por valor crescente e depois por nome.
            var r4 = products.Where(p => p.Category.Tier == 1).OrderBy(p => p.Price).ThenBy(p => p.Name);
            Print("Tier 1 order by price then by name: ", r4);

            // pule 2 elementos e pegue 4 elementos
            var r5 = r4.Skip(2).Take(4);
            Print("Skip 2 and take 4: ", r5);

            var r6 = products.First();
            Console.WriteLine("first: " + r6 + "\n");

            //retornando apenas um elemento
            var r7 = products.Where(p => p.Id == 3).SingleOrDefault();
            Console.WriteLine("Single or default: " + r7);

            var r8 = products.Where(p => p.Category.Id == 1).Select(p => p.Price).Aggregate(0.0, (x,y) => x + y);
            Console.WriteLine("Category 1 aggregate sum: " + r8 + "\n");


            // group by

            var r9 = products.GroupBy(p => p.Category);
            foreach (IGrouping<Category, Product> group in r9)
            {
                Console.WriteLine("Category " + group.Key.Name + ":");
                foreach (Product p in group)
                {
                    Console.WriteLine(p);
                }
                Console.WriteLine();
            }
        }
    }
}
