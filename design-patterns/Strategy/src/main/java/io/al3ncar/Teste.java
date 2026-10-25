package io.al3ncar;

public class Teste {

    public static void main(String[] args) {

        //STRATEGY
        Comportamento normal = new ComportamentoNormal();
        Comportamento defencivo = new ComportamentoDefensivo();
        Comportamento agrssivo = new ComportamentoAgressivo();

        RoboContexto robo = new RoboContexto();
        robo.setStrategy(normal);
        robo.mover();

        robo.setStrategy(agrssivo);
        robo.mover();

        robo.setStrategy(defencivo);
        robo.mover();


    }
}
