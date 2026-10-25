using System;

namespace ExtentionMethodsExample
{
    class Program
    {
        static void Main(string[] args)
        {
            DateTime dt = new DateTime(2021, 03, 05, 15, 36, 59);
            Console.WriteLine(dt.ElapsedTime());
        }
    }
}
