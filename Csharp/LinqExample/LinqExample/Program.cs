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
        }
    }
}
