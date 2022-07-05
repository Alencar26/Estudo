package model.umpramuitos;

import javax.persistence.*;
import java.util.Date;
import java.util.List;

@Entity
@Table(name = "pedidos")
public class Pedido {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false)
    private Date data;

    /* quando a operação é UmParaMuitos ou UmParaUm por padrão a
       operação é fetch = FetchType.EAGER que trás os elementos
       do banco em uma única operação. (não precisa informar o
       parâmetro nesse caso, pois é default).

       No caso de MuitosParaUm ou MuitosParaMuitos por padrão o
       fetch é = FetchType.LAZY que não tráz tudo.
    */
    @OneToMany(mappedBy = "pedido", fetch = FetchType.EAGER)
    private List<ItemPedido> itens;

    public Pedido() {
        this(new Date());
    }

    public Pedido(Date data) {
        this.data = data;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Date getData() {
        return data;
    }

    public void setData(Date data) {
        this.data = data;
    }

    public List<ItemPedido> getItens() {
        return itens;
    }

    public void setItens(List<ItemPedido> itens) {
        this.itens = itens;
    }
}
