<?php

namespace Alura\Pdo\Infra\Database;

use PDO;

class FabricaDeConexaoDB {

    public static function criarConexao(): PDO {
        $databasePath = __DIR__.'/../../../db.sqlite';
        $connection = new PDO('sqlite:'.$databasePath);
        $connection->setAttribute(PDO::ATTR_ERRMODE,PDO::ERRMODE_EXCEPTION);
        
        return $connection;
    }
}