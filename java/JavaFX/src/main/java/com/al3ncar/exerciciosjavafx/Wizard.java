package com.al3ncar.exerciciosjavafx;

import javafx.application.Application;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.HBox;
import javafx.stage.Stage;

public class Wizard extends Application {

    private Stage janela;
    private Scene passo1;
    private Scene passo2;
    private Scene passo3;

    @Override
    public void start(Stage stage) throws Exception {
        janela = stage;

        criarPasso1();
        criarPasso2();
        criarPasso3();

        janela.setScene(passo1);
        janela.setTitle("Wizard :: Passo 1");
        janela.show();
    }

    private void criarPasso1() {
        Button btnProximoPasso = new Button("Ir para passo 2 >>");

        btnProximoPasso.setOnAction(e -> {
            janela.setScene(passo2);
            janela.setTitle("Wizard :: Passo 2");
        });

        HBox box = new HBox();
        box.setAlignment(Pos.CENTER);
        box.getChildren().add(btnProximoPasso);

        passo1 = new Scene(box, 400, 400);
    }

    private void criarPasso2() {
        Button btnProximoPasso = new Button("Ir para passo 3 >>");

        btnProximoPasso.setOnAction(e -> {
            janela.setScene(passo3);
            janela.setTitle("Wizard :: Passo 3");
        });

        Button btnPassoAnterior = new Button("<< Voltar para passo 1");

        btnPassoAnterior.setOnAction(e -> {
            janela.setScene(passo1);
            janela.setTitle("Wizard :: Passo 1");
        });
        HBox box = new HBox();
        box.setAlignment(Pos.CENTER);
        box.getChildren().add(btnPassoAnterior);
        box.getChildren().add(btnProximoPasso);

        passo2 = new Scene(box, 400, 400);
    }

    private void criarPasso3() {

        Button btnPassoAnterior = new Button("<< Voltar para passo 2");

        btnPassoAnterior.setOnAction(e -> {
            janela.setScene(passo2);
            janela.setTitle("Wizard :: Passo 2");
        });

        Button btnSair = new Button("Sair");
        btnSair.setOnAction(e -> {
            // zero significa que está concluíndo a aplicação sem erro.
            System.exit(0);
        });

        HBox box = new HBox();
        box.setAlignment(Pos.CENTER);
        box.getChildren().add(btnPassoAnterior);
        box.getChildren().add(btnSair);

        passo3 = new Scene(box, 400, 400);
    }

    public static void main(String[] args) {
        launch(args);
    }
}
