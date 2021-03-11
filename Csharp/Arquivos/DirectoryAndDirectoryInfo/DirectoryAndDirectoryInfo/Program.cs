using System;
using System.IO;
using System.Collections.Generic;

namespace DirectoryAndDirectoryInfo
{
    class Program
    {
        static void Main(string[] args)
        {
            string path = @"C:\Users\andre\www\ws_csarp\Exercicios_POO\Csharp\Arquivos\pastaDeTeste";

            try
            {
                // listando o caminho de todas as pastar(inclusive aninhadas)
                IEnumerable<string> folders = Directory.EnumerateDirectories(path, "*.*", SearchOption.AllDirectories);
                Console.WriteLine("FOLDERS: ");
                foreach (string f in folders)
                {
                    Console.WriteLine(f);
                }

                // listando o caminho de todos os arquivos
                var files = Directory.EnumerateFiles(path, "*.*", SearchOption.AllDirectories);
                Console.WriteLine("FILES: ");
                foreach (string f in files)
                {
                    Console.WriteLine(f);
                }

                //criando uma pasta
                Directory.CreateDirectory($@"{path}\newFolder");
            }
            catch (IOException e)
            {
                Console.WriteLine("houve um erro:");
                Console.WriteLine(e.Message);
            }
        }
    }
}
