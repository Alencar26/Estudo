using System;
using System.IO;

namespace StreamWriterExample
{
    class Program
    {
        static void Main(string[] args)
        {
            string sourcePath = @"C:caminho_do_arquivo\Arquivos\File1.txt";
            string targetPath = @"C:caminho_do_arquivo\Arquivos\File2.txt";

            try
            {
                using (StreamReader sr = File.OpenText(sourcePath))
                {
                    while (!sr.EndOfStream)
                    {
                        string line = sr.ReadLine();
                        using (StreamWriter sw = File.AppendText(targetPath))
                        {
                            sw.WriteLine(line.ToUpper());
                        }
                    }
                }


                // outra fomra de resolver mesmo problema:

                string[] lines = File.ReadAllLines(sourcePath);

                using (StreamWriter sw = File.AppendText(targetPath))
                {
                    foreach (string line in lines)
                    {
                        sw.WriteLine(line.ToUpper());
                    }
                }
            }
            catch (IOException e)
            {
                Console.WriteLine("houve um erro:");
                Console.WriteLine(e.Message);
            }
        }
    }
}
