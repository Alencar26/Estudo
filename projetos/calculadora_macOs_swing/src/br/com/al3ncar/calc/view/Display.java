package br.com.al3ncar.calc.view;

import br.com.al3ncar.calc.model.Memoria;

import javax.swing.*;
import java.awt.*;

public class Display extends JPanel {

    private final JLabel label;

    public Display() {
        setBackground(new Color(33,37,43));
        label  = new JLabel(Memoria.getInstancia().getTextoAtual());
        label.setForeground(Color.WHITE);
        label.setFont(new Font("courier", Font.PLAIN, 30));

        setLayout(new FlowLayout(FlowLayout.RIGHT, 10, 25 ));

        add(label);
    }
}
