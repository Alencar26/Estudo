using System;
using System.Text;
using System.Globalization;
using System.Collections.Generic;
using SystemEmployee.Entities;

namespace SystemEmployee
{
    class Program
    {
        static void Main(string[] args)
        {
            List<Employee> employeesList = new List<Employee>();

            string name;
            int hours;
            double valuePerHour;
            double additionalCharge;

            Console.Write("Enter the number of employees:");
            int numEmployee = int.Parse(Console.ReadLine());

            for (int i = 1; i <= numEmployee; i++)
            {
                Console.WriteLine($"Employee #{i} data: ");
                Console.Write("Outsourced (y/n)? ");
                string outsourced = Console.ReadLine();

                switch (outsourced)
                {
                    case "y":
                        Console.Write("Name: ");
                        name = Console.ReadLine();
                        Console.Write("Hours: ");
                        hours = int.Parse(Console.ReadLine());
                        Console.Write("Value per hour: ");
                        valuePerHour = double.Parse(Console.ReadLine());
                        Console.Write("Additional charge: ");
                        additionalCharge = double.Parse(Console.ReadLine());
                        
                        OutsourcedEmployee outsourcedEmployee = new OutsourcedEmployee(name, hours, valuePerHour, additionalCharge);
                        employeesList.Add(outsourcedEmployee);
                        break;
                    case "n":
                        Console.Write("Name: ");
                        name = Console.ReadLine();
                        Console.Write("Hours: ");
                        hours = int.Parse(Console.ReadLine());
                        Console.Write("Value per hour: ");
                        valuePerHour = double.Parse(Console.ReadLine());

                        Employee employee = new Employee(name, hours, valuePerHour);
                        employeesList.Add(employee);
                        break;
                }
            }

            Console.WriteLine("PAYMENTS:");
            foreach (Employee employee in employeesList)
            {
                StringBuilder sb = new StringBuilder();
                sb.Append(employee.Name);
                sb.Append($" - $ {employee.Payment().ToString("F2", CultureInfo.InvariantCulture)}");
                Console.WriteLine(sb.ToString());
            }
        }
    }
}
