using System;
using DelegateExample.Service;

namespace DelegateExample
{
    delegate double Operation(double x, double y);

    class Program
    {
        static void Main(string[] args)
        {
            double x = 10;
            double y = 5;
            double result;

            Operation opSum = CalculationService.Sum;
            Operation opMax = CalculationService.Max;

            result = opSum(x, y);

            Console.WriteLine(result);
        }
    }
}
