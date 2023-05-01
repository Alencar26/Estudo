package Janelas;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.FlowLayout;

import javax.swing.BoxLayout;
import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.border.Border;

public class Janela extends JFrame {
    
    public Janela() {
        setJanela();
        iniciarComponentes();
    }

    private void setJanela() {
        this.setSize(200, 200);
        this.setTitle("Minha Janela");
        this.setDefaultCloseOperation(EXIT_ON_CLOSE); // matar processo quando fechar o programa.
        this.setResizable(true); // não deixa redimencionar a janela.
        this.setMinimumSize(new Dimension(200,200)); // tamanho minimo
        this.setMaximumSize(new Dimension(300,300)); //Tamanoh máximo
        this.getContentPane().setBackground(Color.GREEN); // cor de fundo
        this.setLayout(null); // não deixa usar o layout padrão.
    }

    private void inserirPanelsBorderLayout() {
        JPanel panelCentro = new JPanel();
        panelCentro.setBackground(Color.RED);
        this.getContentPane().add(panelCentro, BorderLayout.CENTER); // adiciona o panel

        JPanel panelNorte = new JPanel();
        panelNorte.setBackground(Color.BLACK);
        this.getContentPane().add(panelNorte, BorderLayout.NORTH); // adiciona o panel

        JPanel panelLeste = new JPanel();
        panelLeste.setBackground(Color.GREEN);
        this.getContentPane().add(panelLeste, BorderLayout.WEST); // adiciona o panel
    }

    private void inserirPanelsFlowLayout() {
        this.setLayout(new FlowLayout());
        JPanel panelCentro = new JPanel();
        panelCentro.setBackground(Color.RED);
        this.getContentPane().add(panelCentro); // adiciona o panel

        JPanel panelNorte = new JPanel();
        panelNorte.setBackground(Color.BLACK);
        this.getContentPane().add(panelNorte); // adiciona o panel

        JPanel panelLeste = new JPanel();
        panelLeste.setBackground(Color.YELLOW);
        this.getContentPane().add(panelLeste); // adiciona o panel
    }

    private void inserirPanelsBoxLayout() {
        this.setLayout(new BoxLayout(this.getContentPane(), BoxLayout.Y_AXIS));
        JPanel panelCentro = new JPanel();
        panelCentro.setBackground(Color.RED);
        this.getContentPane().add(panelCentro); // adiciona o panel

        JPanel panelNorte = new JPanel();
        panelNorte.setBackground(Color.BLACK);
        this.getContentPane().add(panelNorte); // adiciona o panel

        JPanel panelLeste = new JPanel();
        panelLeste.setBackground(Color.YELLOW);
        this.getContentPane().add(panelLeste); // adiciona o panel
    }

    private void inserirPanelsSemLayout() {
        //Para que funcione esse método sem layout, tem que adicionar o layout nulo no JFrame.
        JPanel panelCentro = new JPanel();
        panelCentro.setBounds(0, 0, 100, 100);
        panelCentro.setBackground(Color.RED);
        this.getContentPane().add(panelCentro); // adiciona o panel

        JPanel panelNorte = new JPanel();
        panelNorte.setBounds(100, 100, 100, 100);
        panelNorte.setBackground(Color.BLACK);
        this.getContentPane().add(panelNorte); // adiciona o panel

        JPanel panelLeste = new JPanel();
        panelLeste.setBounds(200, 200, 100, 100);
        panelLeste.setBackground(Color.YELLOW);
        this.getContentPane().add(panelLeste); // adiciona o panel
    }

    
    private void iniciarComponentes() {
        inserirPanelsSemLayout();
        // inserirPanelsBoxLayout();
        // inserirPanelsFlowLayout();
        // inserirPanelsBorderLayout(); 
        // fecharPrograma();
    }

    //método para matar o processo do programa - tem uma forma mais fácil. Estou usando em setJanela();
    private void fecharPrograma() {
        addWindowListener(new java.awt.event.WindowAdapter() {
            @Override
            public void windowClosing(java.awt.event.WindowEvent evento) {
                System.exit(0);
            }
        });
    }

    public static void main(String[] args) {
        new Janela().setVisible(true);
    }
}
