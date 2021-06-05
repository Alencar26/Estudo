package Server;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Locale;
import java.util.Scanner;

public class SrvSocket {

    public void execute(){
        try{
            ServerSocket server = new ServerSocket(12345);

            String recebido = "";
            while(!recebido.toUpperCase().equals("FECHAR")) {
                System.out.println("Aguardando conexão");
                Socket client = server.accept();
                System.out.println("Conectou!");
                Scanner input = new Scanner(client.getInputStream());
                recebido = "";
                while (!recebido.toUpperCase().equals("SAIR") && !recebido.toUpperCase().equals("FECHAR")){
                    recebido = input.nextLine();
                    System.out.println(recebido);
                }
                client.close();
            }
            server.close();
        } catch (IOException e){
            e.printStackTrace();
        }
        System.out.println("Server Down");
    }

}
