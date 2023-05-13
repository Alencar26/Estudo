package Janelas;

import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JLabel;

public class ImageLabel extends JFrame{
    
    ImageIcon iagem = new ImageIcon(getClass().getResource("mario.png"));
    JLabel label = new JLabel(iagem);

    public ImageLabel() {

        add(label);

        setSize(500, 500);
        setDefaultCloseOperation(EXIT_ON_CLOSE);
        setLocationRelativeTo(this);
        setVisible(true);
    }

    public static void main(String[] args) {
        new ImageLabel();
    }
}
