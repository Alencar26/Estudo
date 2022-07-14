    module com.al3ncar.exerciciosjavafx {
    requires javafx.controls;
    requires javafx.fxml;


    opens com.al3ncar.exerciciosjavafx
            to javafx.fxml;
    exports com.al3ncar.exerciciosjavafx;

    opens  com.al3ncar.layout
            to javafx.fxml;
    exports com.al3ncar.layout;
    }