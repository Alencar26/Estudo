using System;
using System.Collections.Generic;
using System.Text;

namespace DelegateExample.Service
{
    class CalculatioShowService
    {
        public static void ShowMax(double x, double y)
        {
            Console.WriteLine((x > y) ? x : y);
        }

        public static void ShowSum(double x, double y)
        {
            Console.WriteLine(x + y);
        }
    }
}
