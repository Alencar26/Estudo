package com.al3ncar.exerciciosjavafx;

import javafx.application.Application;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.util.Objects;

public class Contador extends Application {

    private int  contador = 0;

    private void atualizaLblNumero(Label lblNumero) {

        lblNumero.setText(Integer.toString(contador));

        //removendo quaisquer estilos que já estejam aplicados
        lblNumero.getStyleClass().remove("verde");
        lblNumero.getStyleClass().remove("vermelho");

        if (contador > 0) {
            lblNumero.getStyleClass().add("verde");
        } else if (contador < 0) {
            lblNumero.getStyleClass().add("vermelho");
        }
    }

    @Override
    public void start(Stage stage) throws Exception {

        Label lblTitulo = new Label("Contador");
        lblTitulo.getStyleClass().add("titulo");
        Label lblNumero = new Label("0");
        lblNumero.getStyleClass().add("numero");

        Button btnDecremento = new Button("-");
        btnDecremento.getStyleClass().add("botoes");
        btnDecremento.setOnAction(e -> {
            contador--;
            atualizaLblNumero(lblNumero);
        });

        Button btnIncremento = new Button("+");
        btnIncremento.getStyleClass().add("botoes");
        btnIncremento.setOnAction(e -> {
            contador++;
            atualizaLblNumero(lblNumero);
        });

        HBox boxBotoes = new HBox();
        boxBotoes.setAlignment(Pos.CENTER);
        boxBotoes.setSpacing(10);
        boxBotoes.getChildren().add(btnDecremento);
        boxBotoes.getChildren().add(btnIncremento);

        VBox boxPrincipal = new VBox();
        boxPrincipal.getStyleClass().add("conteudo");
        boxPrincipal.setAlignment(Pos.CENTER);
        boxPrincipal.setSpacing(10);
        boxPrincipal.getChildren().add(lblTitulo);
        boxPrincipal.getChildren().add(lblNumero);
        boxPrincipal.getChildren().add(boxBotoes);

        String PathCSS = Objects
                .requireNonNull(
                    getClass()
                            .getResource("/main/java/com/al3ncar/exerciciosjavafx/Contador.css"))
                .toExternalForm();
        Scene cenaPrincipal = new Scene(boxPrincipal, 400 ,400);
        cenaPrincipal.getStylesheets().add(PathCSS);
        cenaPrincipal.getStylesheets().add("https://fonts.googleapis.com/css2?family=Oswald:wght@400;700&display=swap");

        stage.setScene(cenaPrincipal);
        stage.show();
    }

    public static void main(String[] args) {
        launch(args);
    }
}
