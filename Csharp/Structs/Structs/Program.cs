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

            string resultado = (p.X > p.Y ) ? "p.X é maior que p.Y" : "p.Y é maior que p.X";
            Console.WriteLine(resultado);
            Console.WriteLine(p);
        }
    }
}
