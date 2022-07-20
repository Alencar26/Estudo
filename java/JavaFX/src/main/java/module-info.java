    module com.al3ncar.exerciciosjavafx {
    requires javafx.controls;
    requires javafx.fxml;
    requires org.controlsfx.controls;


    opens com.al3ncar.exerciciosjavafx;
    opens  com.al3ncar.layout;
    opens com.al3ncar.fxml
            to javafx.fxml;
    exports com.al3ncar.fxml;
    }