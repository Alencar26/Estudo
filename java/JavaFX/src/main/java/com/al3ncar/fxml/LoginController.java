package com.al3ncar.fxml;

import javafx.fxml.FXML;
import javafx.geometry.Pos;
import javafx.scene.control.PasswordField;
import javafx.scene.control.TextField;
import org.controlsfx.control.Notifications;

public class LoginController {

    @FXML
    private TextField emailField;
    @FXML
    private PasswordField passwordField;

    public void entrar() {
        boolean emailValido = emailField.getText().equals("admin");
        boolean senhaValida = passwordField.getText().equals("123456");

        if (emailValido && senhaValida) {
            Notifications.create()
                    .position(Pos.TOP_RIGHT)
                    .title("Login FXML")
                    .text("Login Efetuado com sucesso").showInformation();
        } else {
            Notifications.create()
                    .position(Pos.TOP_RIGHT)
                    .title("Login FXML")
                    .text("Login falhou").showError();
        }
    }
}
