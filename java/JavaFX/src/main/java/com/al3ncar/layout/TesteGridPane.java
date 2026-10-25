package com.al3ncar.layout;

import javafx.scene.layout.ColumnConstraints;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.RowConstraints;

public class TesteGridPane extends GridPane {

    public TesteGridPane() {

        Caixa c1 = new Caixa().comTexto("1");
        Caixa c2 = new Caixa().comTexto("2");
        Caixa c3 = new Caixa().comTexto("3");
        Caixa c4 = new Caixa().comTexto("4");
        Caixa c5 = new Caixa().comTexto("5");
        Caixa c6 = new Caixa().comTexto("6");

        setGridLinesVisible(true); // deixar as linhas do grid visíveis

        // criando 5 colunas
        getColumnConstraints().addAll(
                columnConstraints(),
                columnConstraints(),
                columnConstraints(),
                columnConstraints(),
                columnConstraints()
        );

        // criando 5 linhas
        getRowConstraints().addAll(
                rowConstraints(),
                rowConstraints(),
                rowConstraints(),
                rowConstraints(),
                rowConstraints()
        );

        //espaçamento estre os itens
        setVgap(10);
        setHgap(10);

        //add(caixa, coluna, linha)
        //add(caixa, coluna, linha, qntsColunasOcupar, qntsLinhasOcupar)
        add(c1, 0, 0, 2, 2);
        add(c2, 1, 1, 2, 2);
        add(c3, 4, 2, 1, 3);
        add(c4, 3, 1);
        add(c5, 0, 4, 2, 1);
        add(c6, 3, 3);
    }

    private ColumnConstraints columnConstraints() {
        ColumnConstraints cc = new ColumnConstraints();
        cc.setPercentWidth(20); // percentual do tamanho da largura
        cc.setFillWidth(true); //preencher a largura
        return cc;
    }

    private RowConstraints rowConstraints() {
        RowConstraints rc = new RowConstraints();
        rc.setPercentHeight(20); // percentual do tamanho da altura
        rc.setFillHeight(true); // preencher altura
        return rc;
    }
}
