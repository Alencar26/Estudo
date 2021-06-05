package Client;

import java.io.IOException;
import java.io.PrintStream;
import java.net.Socket;
import java.util.Scanner;

public class CliSocket {

    public void execute(){

        Scanner teclado = new Scanner(System.in);
        try{
            Socket client = new Socket("127.0.0.1", 12345);

            PrintStream output = new PrintStream(client.getOutputStream());
            Scanner input = new Scanner(client.getInputStream());

            String mensagem = "";
            while(!mensagem.toUpperCase().equals("SAIR") && !mensagem.toUpperCase().equals("FECHAR")){
                mensagem = teclado.nextLine();
                output.println(mensagem);
                System.out.println(input.nextLine());
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
