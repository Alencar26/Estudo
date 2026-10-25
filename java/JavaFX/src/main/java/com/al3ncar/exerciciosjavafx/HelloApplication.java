package com.al3ncar.exerciciosjavafx;

import javafx.application.Application;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.HBox;
import javafx.stage.Stage;

import java.io.IOException;

public class HelloApplication extends Application {
    @Override
    public void start(Stage     stage) throws IOException {

        Button b1 = new Button("A");
        Button b2 = new Button("B");
        Button b3 = new Button("C");

        HBox box = new HBox();
        box.setAlignment(Pos.CENTER);
        box.setSpacing(10);
        box.getChildren().add(b1);
        box.getChildren().add(b2);
        box.getChildren().add(b3);

        Scene cena = new Scene(box, 150, 100);

        stage.setScene(cena);
        stage.show();
    }

    public static void main(String[] args) {
        launch();
    }
}