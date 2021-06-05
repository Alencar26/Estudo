package Server;

import java.io.IOException;
import java.io.PrintStream;
import java.net.Socket;
import java.util.Scanner;

public class MultConex implements  Runnable{

    private Socket client;

    public MultConex(Socket client){
        this.client = client;
    }

    @Override
    public void run() {
        int count = 0;
        Scanner input = null;
        try {
            input = new Scanner(client.getInputStream());

            PrintStream output = new PrintStream(client.getOutputStream());

            String recebido = "";
            while (!recebido.toUpperCase().equals("SAIR")){
                recebido = input.nextLine();
                System.out.println(recebido);

                output.println(++count);
            }
            client.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
