package com.al3ncar.layout;

import javafx.geometry.Insets;
import javafx.geometry.Orientation;
import javafx.geometry.Pos;
import javafx.scene.layout.FlowPane;

public class TesteFlowPane extends FlowPane {

    public TesteFlowPane() {

        /*
            FlowPane organiza os elementos no fluxo em que são
            escritas no código, uma ao lado a outra e se na mesma
            linha não há espaço ele quebra para linha de baixo.
            Pode-se fazer a analogia à escrita.
        */

        Quadrado q1 = new Quadrado();
        Quadrado q2 = new Quadrado();
        Quadrado q3 = new Quadrado();
        Quadrado q4 = new Quadrado();
        Quadrado q5 = new Quadrado();

        setHgap(10); //espaço na horizontal entre os elementos
        setVgap(10); //espaço na vertical entre os elementos
        setPadding(new Insets(10)); //padding igual web

        setOrientation(Orientation.VERTICAL); //setando a orientação dos elementos
        setAlignment(Pos.CENTER); //alinhar os elementos

        getChildren().addAll(q1, q2, q3, q4, q5);
    }
}
