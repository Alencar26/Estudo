import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;

public class ConsultarPessoas1 {

    public static void main(String[] args) throws SQLException {

        Connection conexao = FabricaConexao.getConexao("curso_java");
        String sql = "SELECT * FROM pessoas;";

        Statement stmt = conexao.createStatement();
        ResultSet result = stmt.executeQuery(sql);

        List<Pessoa> pessoas = new ArrayList<>();

        while (result.next()) {
            int codigo = result.getInt("codigo");
            String nome = result.getString("nome");
            pessoas.add(new Pessoa(codigo, nome));
        }

        for (Pessoa p : pessoas) {
            System.out.println(p);
        }

        stmt.close();
        conexao.close();
    }
}
