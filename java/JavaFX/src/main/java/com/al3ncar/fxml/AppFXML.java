package com.al3ncar.fxml;

import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.scene.layout.GridPane;
import javafx.stage.Stage;

import java.net.URL;
import java.util.Objects;

public class AppFXML extends Application {

    @Override
    public void start(Stage stage) throws Exception {
        String estilosCSS = Objects
                .requireNonNull(
                        getClass().getResource("/main/java/com/al3ncar/fxml/Login.css"))
                .toExternalForm();
        URL loginFxml = getClass().getResource("/main/java/com/al3ncar/fxml/Login.fxml");
        assert loginFxml != null;

        GridPane raiz = FXMLLoader.load(loginFxml);

        Scene cena = new Scene(raiz,350, 350);
        cena.getStylesheets().add(estilosCSS);

        stage.setTitle("Tela de Login");
        stage.setResizable(false); // não deixa o utilizador mudar tamanho da tela.
        stage.setScene(cena);
        stage.show();
    }

    public static void main(String[] args) {
        launch(args);
    }
}
