package Server;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Scanner;

public class SrvSocket {

    public void execute(){
        try{
            ServerSocket server = new ServerSocket(12345);
            System.out.println("Aguardando conexão");
            Socket client = server.accept();
            System.out.println("conectado!");

            Scanner input = new Scanner(client.getInputStream());
            while(input.hasNextLine()) {
                String recebido = input.nextLine();
                System.out.println(recebido);
            }

            client.close();
            server.close();
        } catch (IOException e){
            e.printStackTrace();
        }
    }

}
