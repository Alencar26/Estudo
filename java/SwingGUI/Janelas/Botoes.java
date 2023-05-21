package Janelas;

import java.awt.Color;
import java.awt.Font;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.awt.event.MouseMotionListener;
import java.awt.event.MouseWheelEvent;
import java.awt.event.MouseWheelListener;

import javax.swing.BorderFactory;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.SwingConstants;
import javax.swing.border.BevelBorder;

public class Botoes extends JFrame {
    
    private JPanel painel;
    private JButton botao;
    private Color 
                    original    = Color.red,
                    precionado  = Color.green,
                    posicionado = Color.blue;
    
    public Botoes() {
        loadComponents();
    }

    private void initPanel() {
        if (this.painel == null) {
            this.painel = new JPanel();
        }
    }

    private void setPanel() {
        this.painel.setLayout(null); // não deixa usar o layout padrão.
        this.painel.setBackground(Color.GREEN);
        this.painel.setBounds(0, 0, 1000, 1000);
        this.getContentPane().add(this.painel);
    }

    private void setLabel() {
        JLabel etiqueta = new JLabel("Minha Label aqui", SwingConstants.CENTER);
        etiqueta.setFont(new Font("Arial", Font.ITALIC, 20)); // define a fonte.
        etiqueta.setBounds(0, 0, 500, 500); //precisa desabilitar o layout padrão (posição da label)
        //coloração do label
        etiqueta.setForeground(Color.BLUE);
        this.painel.add(etiqueta);
    }

    private void setJanela() {
        this.setSize(500, 500);
        this.setTitle("Minha Janela");
        this.setDefaultCloseOperation(EXIT_ON_CLOSE);
        this.setVisible(true);
        setLocationRelativeTo(this); //centralizar janela
    }


    private void setButton() {
        botao = new JButton("Clique aqui");
        botao.setFont(new Font("Courier", Font.BOLD, 20));
        botao.setBackground(original);
        botao.setForeground(Color.BLACK);
        botao.setBounds(10, 10, 200, 40);
        botao.setOpaque(true); //indica se o botão é tranparente ou não
        botao.setBorder(BorderFactory.createBevelBorder(BevelBorder.RAISED)); //efeito de sombras
        this.painel.add(botao);

        // interacaoButtonClicks();
        // interacaoBotaoMovimentos();
        interacaoBotaoMouseScroll(); //rolagem do mouse
    }

    //interacao com Mouse Listener. Só lida com os clicks do mouse.
    private void interacaoButtonClicks() {
        botao.addMouseListener(new MouseListener() {
            @Override
            public void mouseClicked(MouseEvent e) {
                System.out.println("Botao clicado");
            }

            @Override
            public void mousePressed(MouseEvent e) {
                botao.setBackground(precionado);
                // System.out.println("Botao pressionado");
            }

            @Override
            public void mouseReleased(MouseEvent e) {
                botao.setBackground(original);
                // System.out.println("Botao liberado");
            }

            @Override
            public void mouseEntered(MouseEvent e) {
                botao.setBackground(posicionado);
                // System.out.println("Botao entrou");
            }

            @Override
            public void mouseExited(MouseEvent e) {
                botao.setBackground(original);
            }
        });
    }

    //movimentação e movimentação com click
    private void interacaoBotaoMovimentos() {
        botao.addMouseMotionListener(new MouseMotionListener() {

            int contador = 0;

            //CLICAR E MOVER MOUSE
            @Override
            public void mouseDragged(MouseEvent e) {
                contador--;
                System.out.println("Botao CLICADO E MOVIDO: " + contador);
            }

            //MOVER MOUSE
            @Override
            public void mouseMoved(MouseEvent e) {
                contador++;
                System.out.println("Botao MOVIDO: " + contador);
            }
            
        });
    }

    //moralgem do mouse
    private void interacaoBotaoMouseScroll() {
        botao.addMouseWheelListener(new MouseWheelListener() {

            int contador = 0;
            @Override
            public void mouseWheelMoved(MouseWheelEvent e) {
                if (e.getPreciseWheelRotation() == -1.0) {
                    contador--;
                    System.out.println("PARA CIMA: " + contador);
                } else {
                    contador++;
                    System.out.println("PARA BAIXO:" + contador);
                }        
            }
        });
    }

    private void loadComponents() {
        setJanela();
        initPanel();
        setPanel();
        setButton();
    }

    public static void main(String[] args) {
        new Botoes();
    }
}
