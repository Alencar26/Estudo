package Server;

import java.io.IOException;
import java.io.PrintStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Locale;
import java.util.Scanner;

public class SrvSocket {

    public void execute(){
        try{
            ServerSocket server = new ServerSocket(12345);

            String recebido = "";
            while(true) {


                System.out.println("Aguardando conexão");
                Socket client = server.accept();
                System.out.println("Conectou!");

                new Thread(new MultConex(client)).start();
            }
           // server.close();
        } catch (IOException e){
            e.printStackTrace();
        }
        System.out.println("Server Down");
    }

}
