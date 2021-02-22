using System;

namespace Structs
{
    class Program
    {
        static void Main(string[] args)
        {
            //structs são criados diretamente na stack da memória.
            Point p;
            p.X = 10;
            p.Y = 20;

            Console.WriteLine(p);
        }
    }
}
