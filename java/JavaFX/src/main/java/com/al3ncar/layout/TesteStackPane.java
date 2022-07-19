package com.al3ncar.layout;

import javafx.scene.layout.StackPane;

public class TesteStackPane extends StackPane {

    public TesteStackPane() {

        Caixa c1 = new Caixa().comTexto("1");
        Caixa c2 = new Caixa().comTexto("2");
        Caixa c3 = new Caixa().comTexto("3");
        Caixa c4 = new Caixa().comTexto("4");
        Caixa c5 = new Caixa().comTexto("5");
        Caixa c6 = new Caixa().comTexto("6");

        // adiciona as caixas em pilha, onde a última adicionada é a primeira da pilha
        getChildren().addAll(c2,c3,c4,c5,c6,c1);

        setOnMouseClicked(e -> {
            //verificando se o click está ocorrendo em um ponto maior que a metade da tela.
            // click está ocorrendo do lado direito.
            if (e.getSceneX() > (getScene().getWidth() / 2)) {
                //pega o elemento de indice 0 e manda para frente da pilha de elementos
                getChildren().get(0).toFront();
            } else {
                //pega o elemento de indice 5 (que está em exibição) e manda para o final da pilha
                getChildren().get(5).toBack();
            }
        });


    }
}
