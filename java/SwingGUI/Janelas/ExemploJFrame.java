package Janelas;

import javax.swing.*;
import java.awt.*;

public class ExemploJFrame {
    
    public static void main(String[] args) {
        JFrame frame = new JFrame("Exemplo JFrame");
        frame.setSize(500, 500);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        // Obter o painel principal do JFrame
        Container contentPane = frame.getContentPane();
        
        // Criar um painel para adicionar na parte inferior do JFrame
        JPanel painelInferior = new JPanel();
        JLabel label = new JLabel("Sou um JLabel");
        JTextField textField = new JTextField(20);
        JButton botao = new JButton("Clique aqui");
        painelInferior.add(label);
        painelInferior.add(textField);
        painelInferior.add(botao);
        
        // Definir o layout do JFrame
        contentPane.setLayout(new BorderLayout());
        
        // Adicionar o painel principal e o painel inferior no JFrame
        contentPane.add(new JScrollPane(new JList<>(new String[]{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"})), BorderLayout.WEST);
        contentPane.add(painelInferior, BorderLayout.SOUTH);

        // Exibir o JFrame
        frame.setVisible(true);
    }
}
