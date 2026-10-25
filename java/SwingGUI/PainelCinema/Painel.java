package PainelCinema;

import java.awt.Color;
import java.awt.Dimension;
import java.awt.Image;

import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;

public class Painel extends JFrame {
    
    private JPanel panelContainer, panelFilme, panelImagemFilme, panelInfoFilme;
    private Filme filme;
    private JLabel imagem, titulo, genero;

    public Painel() {
        super("Cinema");
        this.setSize(480, 640);
        this.setDefaultCloseOperation(EXIT_ON_CLOSE);
        this.setLocationRelativeTo(null);

        initComponents();
    }

    private void initComponents() {
        setPanelContainer();
        setPanelFilme();
        setPanelImagemFilme();
        setPanelInfoFilme();

        createFilme();
        setLabelImagem();
    }

    private void setLabelImagem() {
        imagem = new JLabel();
        ImageIcon imgIcon = new ImageIcon(filme.getImagem());
        Icon icon = new ImageIcon(imgIcon.getImage().getScaledInstance(180, 260, Image.SCALE_DEFAULT));
        
        imagem.setIcon(icon);
        panelImagemFilme.add(imagem);
    }

    private void createFilme() {
        filme = new Filme("Vingadores", "Ação", "vingadores.jpg");
    }

    private void setPanelInfoFilme() {
        panelInfoFilme = new JPanel();
        panelInfoFilme.setBackground(Color.BLACK);
        panelFilme.add(panelInfoFilme);
    }

    private void setPanelImagemFilme() {
        panelImagemFilme = new JPanel();
        panelImagemFilme.setBackground(Color.GREEN);
        panelImagemFilme.setPreferredSize(new Dimension(180, 260));
        panelFilme.add(panelImagemFilme);
    }

    private void setPanelFilme() {
        panelFilme = new JPanel();
        panelFilme.setMaximumSize(new Dimension(440, 220));
        panelFilme.setBackground(Color.CYAN);
        panelContainer.add(panelFilme);
    }

    private void setPanelContainer() {
        panelContainer = new JPanel();
        panelContainer.setBackground(Color.DARK_GRAY);
        this.add(panelContainer);
    }

    public static void main(String[] args) {
        new Painel().setVisible(true);
    }
}
