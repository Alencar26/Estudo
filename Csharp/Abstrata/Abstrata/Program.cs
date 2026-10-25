using System;
using Abstrata.Entities;
using System.Globalization;
using Abstrata.Entities.Enums;
using System.Collections.Generic;


namespace Abstrata
{
    class Program
    {
        static void Main(string[] args)
        {
            List<Shape> list = new List<Shape>();

            Console.Write("Enter the number of shapes: ");
            int n = int.Parse(Console.ReadLine());

            for (int i = 1; i <= n; i++)
            {
                Console.WriteLine($"Shape #{i} data: ");
                Console.Write("Rectangle or Circle (r/c)? ");
                char c = char.Parse(Console.ReadLine());
                Console.Write("Color (Black, Blue, Red): ");
                Color s = Enum.Parse<Color>(Console.ReadLine());

                switch (c)
                {
                    case 'r':
                        Console.Write("Width: ");
                        double w = double.Parse(Console.ReadLine());
                        Console.Write("Heigth: ");
                        double h = double.Parse(Console.ReadLine());

                        list.Add(new Ractangle(s, w, h));
                        break;
                    case 'c':
                        Console.Write("Radius: ");
                        double r = double.Parse(Console.ReadLine());

                        list.Add(new Circle(s, r));
                        break;
                }
            }

            Console.WriteLine("SAHPE AREAS:");

            foreach (Shape shape in list)
            {
                Console.WriteLine(shape.Area().ToString("F2", CultureInfo.InvariantCulture));
            }
        }
    }
}
