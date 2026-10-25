using System;
using System.Collections.Generic;
using System.Text;

namespace SystemEmployee.Entities
{
    class OutsourcedEmployee : Employee
    {
        public double AdditionalCharge { get; set; }

        public OutsourcedEmployee()  { }

        public OutsourcedEmployee(string name, int hours, double valuePerHour, double additionalCharge)
            : base(name, hours, valuePerHour)
        {
            AdditionalCharge = additionalCharge;
        }

        public sealed override double Payment()
        {
            double salary = base.Payment();
            return salary += AdditionalCharge * 1.1;
        }
    }
}
