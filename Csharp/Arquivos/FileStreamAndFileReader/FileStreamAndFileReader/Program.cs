using System;
using System.IO;

namespace FileStreamAndFileReader
{
    class Program
    {
        static void Main(string[] args)
        {
            string path = @"C:\Users\andre\www\ws_csarp\Exercicios_POO\Csharp\Arquivos\File1.txt";


            // SEM USER O BLOCO USING
            /*
            FileStream fs = null;
            StreamReader sr = null;
            StreamReader srResumido = null;


            try
            {
                //instanciando o FileStream passando por argumento o caminho e o modo como eu quero lidar com o arquivo
                fs = new FileStream(path, FileMode.Open);
                sr = new StreamReader(fs);

                string line = sr.ReadLine();
                Console.WriteLine(line);

                // forma resumida de abrir um arquivo.
                srResumido = File.OpenText(path);

                while (!srResumido.EndOfStream)
                {
                    string linha = srResumido.ReadLine();
                    Console.WriteLine(linha);
                }

            }
            catch (IOException e)
            {
                Console.WriteLine($"Houve um erro: ");
                Console.WriteLine(e.Message);
            }
            finally
            {
                if (sr != null) sr.Close();
                if (fs != null) fs.Close();
                if (srResumido != null) srResumido.Close();
            }
            */


            // COM O BLOCO USING

            try
            {
                using (StreamReader fs = File.OpenText(path))
                {
                    while (!fs.EndOfStream)
                    {
                        string line = fs.ReadLine();
                        Console.WriteLine(line);
                    }
                }
            }
            catch (IOException e)
            {
                Console.WriteLine("Houve um erro: ");
                Console.WriteLine(e.Message);
            }
            

        }
    }
}
