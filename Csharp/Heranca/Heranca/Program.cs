using Heranca.Entities;
using System;
using System.Collections.Generic;
using System.Globalization;

namespace Heranca
{
    class Program
    {
        static void Main(string[] args)
        {
            /*  PARTINDO PARA UMA NOVA REVISÃO EM CLASSES ABSTRATAS (logo após os comentários) */

            /*
            Account acc = new Account(2001, "Ana", 500.0);
            BusinessAccount bacc = new BusinessAccount(2002, "Alencar", 100.0, 500.0);

            //CONVERSÕES
            // (conversão da subclasse para a super classe)
            // UPCASTING
            Account acc1 = bacc; // isso é válido! já que o  business Account é um Account
            Account acc2 = new BusinessAccount(2002, "Alencar", 100.0, 500.0);
            Account acc3 = new SavingsAccount(2003, "Helô", 10000, 0.05);


            //DOWNCASTING

            // isso não funciona
            //BusinessAccount acc4 = acc2;

            BusinessAccount acc4 = (BusinessAccount) acc2; // tem que fazer um casting (conversão)

            //NÃO PODE FAZER ISSO!! nÃO SÃO COMPATÍVEIS
            // BusinessAccount acc5 = (BusinessAccount)acc3; // <<< VAI OCORRER UM ERRO NA HORA DE EXECUTAR O PROGRAMA.

            //verificando se a instância  da variável é do tipo BusinessAccount 
            if (acc3 is BusinessAccount)
            {
                BusinessAccount acc6 = (BusinessAccount)acc3;
            }

            // sintaxe alternativa para converter (casting) o tipo da variável.
            SavingsAccount acc7 = acc3 as SavingsAccount;


            //=======================================================================================================

            Account acc8 = new Account(2001, "Ana", 500.0);
            Account acc9 = new SavingsAccount(2021, "Ana P.", 500.0, 0.01);

            acc8.Withdraw(10);
            acc9.Withdraw(10);

            Console.WriteLine(acc8.Balance);
            Console.WriteLine(acc9.Balance);
            */

            List<Account> list = new List<Account>();

            list.Add(new SavingsAccount(002, "Ana", 230.50, 0.02));
            list.Add(new BusinessAccount(003, "Heloísa", 210.23, 400.00));
            list.Add(new SavingsAccount(004, "Marcos", 0.00, 0.02));
            list.Add(new BusinessAccount(005, "Maria", 1500.22, 652.00));

            double soma = 0;
            foreach (Account account in list)
            {

                soma += account.Balance;
            }

            Console.WriteLine($"Total das Contas: {soma.ToString("F2", CultureInfo.InvariantCulture)}");
        }
    }
}
  