package Janelas;

import java.awt.BorderLayout;
import java.awt.Image;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;
import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;

public class TesteImagem {

    public static void main(String[] args) {
        JFrame frame = new JFrame("Exemplo JLabel com imagem");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setSize(500, 500);
        frame.setVisible(true);
        
        try {
            BufferedImage imagem = ImageIO.read(new File("./mario.png"));
            ImageIcon icon = new ImageIcon(imagem.getScaledInstance(200, 200, Image.SCALE_SMOOTH));
            JLabel label = new JLabel(icon);
            
            JPanel painel = new JPanel();
            painel.add(label);
            
            frame.getContentPane().add(painel, BorderLayout.CENTER);
        } catch (IOException e) {
            System.out.println("Erro ao carregar imagem: " + e.getMessage());
        }

        frame.pack();
        frame.setVisible(true);
    }
}
