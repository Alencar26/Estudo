<?php

namespace Alura\Pdo\Infra\Database;

use PDO;

class FabricaDeConexaoDB {

    public static function criarConexao(): PDO {
        $databasePath = __DIR__.'/../../../db.sqlite';
        return new PDO('sqlite:'.$databasePath);
    }
}