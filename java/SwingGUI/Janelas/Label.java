package Janelas;

import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.Image;

import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.SwingConstants;

public class Label extends JFrame {
    
    public Label() {
        setJanela();
        iniciarComponentes();
        setVisible(true); // torna a janela visível.
    }

    private void setJanela() {
        this.setSize(500, 500);
        this.setTitle("Minha Janela");
        this.setDefaultCloseOperation(EXIT_ON_CLOSE); // matar processo quando fechar o programa.
        this.setResizable(true); // não deixa redimencionar a janela.
        this.setMinimumSize(new Dimension(200,200)); // tamanho minimo
        this.setMaximumSize(new Dimension(300,300)); //Tamanoh máximo
        this.getContentPane().setBackground(Color.WHITE); // cor de fundo
        this.setLayout(null); // não deixa usar o layout padrão.
        setLocationRelativeTo(this); //centralizar janela
    }

    private void iniciarComponentes() {
        inserirPanel();
        // inserirPanelExemplo();
    }

    private void inserirPanel() {
        JPanel panel = new JPanel();
        panel.setLayout(null); // não deixa usar o layout padrão.
        panel.setBackground(Color.ORANGE);
        panel.setBounds(0, 0, 1000, 1000);
        this.getContentPane().add(panel);

        JLabel etiqueta = new JLabel("Assim também da para escrever", SwingConstants.CENTER);
        etiqueta.setFont(new Font("Arial", Font.ITALIC, 20)); // define a fonte.
        etiqueta.setBounds(0, 0, 500, 500); //precisa desabilitar o layout padrão (posição da label)

        //coloração do label
        etiqueta.setForeground(Color.BLUE);
        panel.add(etiqueta);

        //IMAGEM
        ImageIcon mario = new ImageIcon(getClass().getResource("mario.png"));
        JLabel imagem = new JLabel(mario);
        imagem.setBounds( 100, 150, 250, 250);
        imagem.setIcon(new ImageIcon(mario.getImage().getScaledInstance(imagem.getWidth(), imagem.getHeight(), Image.SCALE_SMOOTH)));
        panel.add(imagem);
    }


    private void inserirPanelExemplo() {
        JPanel panel = new JPanel();
        panel.setLayout(null); // não deixa usar o layout padrão.
        panel.setBackground(Color.ORANGE);
        panel.setBounds(0, 0, 300, 300);
        this.getContentPane().add(panel);

        JLabel etiqueta = new JLabel();
        etiqueta.setBounds(0, 0, 300, 300); //precisa desabilitar o layout padrão
        etiqueta.setText("Olá mundo!");

        //coloração do label
        etiqueta.setForeground(Color.BLUE);
        etiqueta.setBackground(Color.RED);
        etiqueta.setOpaque(true); // torna a cor opaca.
        etiqueta.setBorder(javax.swing.BorderFactory.createLineBorder(Color.BLACK)); // adiciona uma borda.
        etiqueta.setHorizontalAlignment(javax.swing.SwingConstants.CENTER); // alinhamento horizontal
        etiqueta.setVerticalAlignment(javax.swing.SwingConstants.CENTER); // alinhamento vertical
        etiqueta.setFont(new java.awt.Font("Arial", 0, 20)); // define a fonte.
        etiqueta.setToolTipText("Este é um texto de exemplo"); // adiciona um texto de ajuda.
        etiqueta.setCursor(new java.awt.Cursor(java.awt.Cursor.HAND_CURSOR)); // adiciona um cursor.
        etiqueta.setHorizontalTextPosition(javax.swing.SwingConstants.CENTER); // alinhamento horizontal do texto.;


        panel.add(etiqueta);
    }

    public static void main(String[] args) {
        new Label();
    }
}
