import java.sql.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class ConsultaPessoas2 {

    public static void main(String[] args) throws SQLException {

        Scanner sc = new Scanner(System.in);
        Connection conexao = FabricaConexao.getConexao("curso_java");
        String sql = "SELECT * FROM pessoas WHERE nome like ?;";

        System.out.print("Informe parte do nome que deseja procurar: ");
        String parteDoNome = sc.nextLine();

        PreparedStatement stmt = conexao.prepareStatement(sql);
        stmt.setString(1, "%"+parteDoNome+"%");
        ResultSet result = stmt.executeQuery();

        List<Pessoa> pessoas = new ArrayList<>();

        while (result.next()) {
            int codigo = result.getInt("codigo");
            String nome = result.getString("nome");
            pessoas.add(new Pessoa(codigo, nome));
        }

        for (Pessoa p : pessoas) {
            System.out.println(p);
        }

        sc.close();
        stmt.close();
        conexao.close();
    }
}