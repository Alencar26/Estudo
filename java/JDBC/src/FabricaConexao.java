import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class FabricaConexao {

    public static Connection getConexao() {

        try {
            final String url = "jdbc:mysql://localhost:3307/?verifyServerCertificate=false&useSSL=true";
            final String user = "root";
            final String password = "mysql";

            return DriverManager.getConnection(url, user, password);
        } catch (SQLException e) {
            // lançando exceção e tranformando em um runtimeException.
            throw new RuntimeException(e);
        }
    }

    public static Connection getConexao(String nomeDoBanco) {

        try {
            final String url = "jdbc:mysql://localhost:3307/"+nomeDoBanco+"?verifyServerCertificate=false&useSSL=true";
            final String user = "root";
            final String password = "mysql";

            return DriverManager.getConnection(url, user, password);
        } catch (SQLException e) {
            // lançando exceção e tranformando em um runtimeException.
            throw new RuntimeException(e);
        }
    }
}
