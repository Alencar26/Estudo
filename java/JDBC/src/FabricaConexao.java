import java.io.IOException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.Properties;

public class FabricaConexao {

    public static Connection getConexao() {

        try {
            Properties prop = getProperties();
            final String url = prop.getProperty("banco.url") + prop.getProperty("banco.params");
            final String user = prop.getProperty("banco.user");
            final String password = prop.getProperty("banco.password");

            return DriverManager.getConnection(url, user, password);
        } catch (SQLException | IOException e) {
            // lançando exceção e tranformando em um runtimeException.
            throw new RuntimeException(e);
        }
    }

    public static Connection getConexao(String nomeDoBanco) {

        try {
            Properties prop = getProperties();
            final String url = prop.getProperty("banco.url") + nomeDoBanco + prop.getProperty("banco.params");
            final String user = prop.getProperty("banco.user");
            final String password = prop.getProperty("banco.password");

            return DriverManager.getConnection(url, user, password);
        } catch (SQLException | IOException e) {
            // lançando exceção e tranformando em um runtimeException.
            throw new RuntimeException(e);
        }
    }

    private static Properties getProperties() throws IOException {
        Properties prop = new Properties();
        final String PATH = "/conexao.properties";
        prop.load(FabricaConexao.class.getResourceAsStream(PATH));
        return prop;
    }
}

