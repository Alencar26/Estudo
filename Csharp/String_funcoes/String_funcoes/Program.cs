using System;

namespace String_funcoes
{
    class Program
    {
        static void Main(string[] args)
        {
            string myString = "abcde FGHIJ ABC abc DEFG     ";

            //converter todos para maiúsculo.
            string maiusculo = myString.ToUpper();

            //convertendo para minuscula.
            string minusculo = myString.ToLower();

            //apaga espaços em branco antes e depois da string
            string trim = myString.Trim();

            //função de prucura primeira ocorrência da string
            int n = myString.IndexOf("FG");

            //método de procura última ocorrcia da string
            int n2 = myString.LastIndexOf("FG");


            //recortar string
            string recorte = myString.Substring(5); // recorta a string a partir da posição 5
            string recorte2 = myString.Substring(5, 4); // começa na posição 5 e recorta 4 caracteres


            //trocar elementos da string
            string trocada = myString.Replace('a', 'X'); // trocando todo caracter 'a' por 'X'
            string trocada2 = myString.Replace("abc", "wxyz"); // trocando strings "abc" por "wxyz"

            // verificar se a string é vazia ou nula
            string empty = "";
            bool b = string.IsNullOrEmpty(empty);
            bool b2 = string.IsNullOrWhiteSpace(myString); // verifica tbm os espaço em brancos

            string vazia = b == true ? "String vazia" : "String não está vazia";
            string vazia2 = b2 == true ? "String vazia" : "String não está vazia";



            //========================== PRINTS ===========================

            Console.WriteLine("String sem modificação: " + myString + "-");
            Console.WriteLine("Maiúscula: " + maiusculo + "-");
            Console.WriteLine("Minúsculo: " + minusculo + "-");
            Console.WriteLine("Trim: " + trim + "-");
            Console.WriteLine("Index onde está o início de FG: " + n);
            Console.WriteLine("Index onde está o término de FG: " + n2);
            Console.WriteLine("string recortada: " + recorte);
            Console.WriteLine("string recortada com 4 caracteres: " + recorte2);
            Console.WriteLine("string troccada: " + trocada);
            Console.WriteLine("string troccada: " + trocada2);
            Console.WriteLine(vazia);
            Console.WriteLine(vazia2);
        }
    }
}
