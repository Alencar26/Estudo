using PredicateExample.Entities;
using System;
using System.Collections.Generic;

namespace ActionExample
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

            //list.ForEach(x => { x.Price += x.Price * 0.1; }); //<== incrementando 10% para cada roduto
            list.ForEach(UpdatePrice); // <== Action


            foreach (Product p in list)
            {
                Console.WriteLine(p);
            }
        }


        //Action
        public static void UpdatePrice(Product p)
        {
            p.Price += p.Price * 0.1;
        }
    }
}
