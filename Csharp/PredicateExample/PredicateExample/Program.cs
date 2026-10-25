using PredicateExample.Entities;
using System;
using System.Collections.Generic;

namespace PredicateExample
{
    class Program
    {
        static void Main(string[] args)
        {

            List<Product> list = new List<Product>();
            list.Add(new Product("Tv", 900.00));
            list.Add(new Product("Mouse", 50.00));
            list.Add(new Product("Tablet", 350.50));
            list.Add(new Product("HD Case", 80.90));

            //list.RemoveAll(x => x.Price <= 100); // <== esse funciona perfeitamente.

            list.RemoveAll(ProductTest); // predicate << sempre temessa assinatura (recebe um obj e retorna um bool)

            foreach (Product p in list)
            {
                Console.WriteLine(p);
            }
        }

        public static bool ProductTest(Product p)
        {
            return p.Price <= 100.0;
        }
    }
}
