package com.al3ncar.layout;

import javafx.geometry.Pos;
import javafx.scene.layout.TilePane;

import java.util.ArrayList;
import java.util.List;

public class TesteTilePane extends TilePane {

    public TesteTilePane() {

        //Mesmo que os elementos mudem de tamanho os "containers" que cada elementos
        //fica são do mesmo tamanho. (os "tijolinhos" são do mesmo tamanho)
        List<Quadrado> quadrados = new ArrayList<>();

        for (int i = 1; i <= 10; i++) {
            quadrados.add(new Quadrado(i * 10));
        }

        //setando a posição dos elementos dentro dos tiles.
        setTileAlignment(Pos.BOTTOM_RIGHT);

        getChildren().addAll(quadrados);
    }
}
