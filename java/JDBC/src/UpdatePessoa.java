import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class UpdatePessoa {

    public static void main(String[] args) throws SQLException {

        Scanner sc = new Scanner(System.in);

        final String sqlConsulta = "SELECT * FROM pessoas WHERE codigo = ?;";
        final String sqlUpdate = "UPDATE pessoas SET nome = ? WHERE codigo = ?;";

        System.out.print("Informe o id da pessoa que deseja alterar o nome: ");
        int codigoPessoa = sc.nextInt();

        Connection conexao = FabricaConexao.getConexao("curso_java");

        PreparedStatement stmt = conexao.prepareStatement(sqlConsulta);
        stmt.setInt(1, codigoPessoa);

        ResultSet result = stmt.executeQuery();

        if (result.next()) {
            Pessoa p = new Pessoa(result.getInt(1), result.getString(2));
            System.out.println("Pessoa selecionada: " + p.getNome());

            System.out.println("Informe o novo nome: ");
            String novoNome = sc.nextLine();

            stmt.close();
            stmt = conexao.prepareStatement(sqlUpdate);
            stmt.setString(1, novoNome);
            stmt.setInt(2, codigoPessoa);
            stmt.execute();
        }

        sc.close();
        stmt.close();
        conexao.close();
    }
}
