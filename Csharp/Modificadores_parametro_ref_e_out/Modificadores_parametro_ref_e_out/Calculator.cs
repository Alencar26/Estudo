using System;
using System.Collections.Generic;
using System.Text;

namespace Modificadores_parametro_ref_e_out
{
    public class Calculator
    {
        public static void Multiplica(ref int x)
        {
            x = x * 2;
        }

        public static void Triplica(int entrada, out int result)
        {
            result = entrada * 3;
        }
    }
}
