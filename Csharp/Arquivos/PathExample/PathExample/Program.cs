using System;
using System.IO;

namespace PathExample
{
    class Program
    {
        static void Main(string[] args)
        {
            string path = @"C:caminho_do_arquivo\Arquivos\pastaDeTeste\File1.txt";

            Console.WriteLine("Caracter de separação: " + Path.DirectorySeparatorChar);
            Console.WriteLine("caracter de separação de pastas: " + Path.PathSeparator);
            Console.WriteLine("Nome do diretório do arquivo: " + Path.GetDirectoryName(path));
            Console.WriteLine("Nome do arquivo: " + Path.GetFileName(path));
            Console.WriteLine("Extenção do arquivo: " + Path.GetExtension(path));
            Console.WriteLine("Nome do arquivo sem a extenção: " + Path.GetFileNameWithoutExtension(path));
            Console.WriteLine("Pasta temporária do sistema: " + Path.GetTempPath());
        }
    }
}
