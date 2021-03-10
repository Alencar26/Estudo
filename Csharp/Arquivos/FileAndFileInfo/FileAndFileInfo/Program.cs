using System;
using System.IO;

namespace FileAndFileInfo
{
    class Program
    {
        static void Main(string[] args)
        {
            // caminho do arquivo de origem
            string sourcePath = @"C:\caminho_do_arquivo\Csharp\Arquivos\File1.txt";
            // caminho do arquivo de destino
            string targetPath = @"C:\caminho_do_arquivo\Csharp\Arquivos\File2.txt";

            try
            {
                // copiando conteúdo do arquivo para outro arquivo!
                FileInfo fileInfo = new FileInfo(sourcePath);
                fileInfo.CopyTo(targetPath);

                // lendo todas as linhas do  arquivo e jogando para um vetor
                string[] lines = File.ReadAllLines(sourcePath);
                foreach (string line in lines)
                {
                    Console.WriteLine(line);
                }
            }
            catch (IOException e)
            {
                Console.WriteLine("Erro inesperado!");
                Console.WriteLine($"Informações do erro: {e.Message}");
            }
        }
    }
}
