import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class TesteConexao {

    public static void main(String[] args) throws SQLException {

        final String url = "jdbc:mysql://localhost:3307?verifyServerCertificate=false&useSSL=true";
        final String user = "root";
        final String password = "mysql";

        //Se passar por esse método significa que o banco se conectou com sucesso.
        Connection conexao = DriverManager.getConnection(url, user, password);
        System.out.print("Conexão efetuada com sucesso.");

        //fechar conexão
        conexao.close();
    }
}
