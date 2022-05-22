import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

public class CriarBanco {

    public static void main(String[] args) throws SQLException {

        final String url = "jdbc:mysql://localhost:3307?verifyServerCertificate=false&useSSL=true";
        final String user = "root";
        final String password = "mysql";

        Connection conexao = DriverManager.getConnection(url, user, password);

        Statement stmt = conexao.createStatement();
        stmt.execute("CREATE DATABASE IF NOT EXISTS curso_java");

        conexao.close();
    }
}
