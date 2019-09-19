package veiculo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//É um contrato de implementação
type Storage interface {
	GetVeiculos() ([]Veiculo, error)
	CreateVeiculo(nome, marca string, ano, modelo int) error
	UpdateVeiculo(id int, veiculo *Veiculo) error
	DeleteVeiculo(id int) error
}
type MySQLStorage struct {
	dbConn *sql.DB
}

// func (dono = MySQLStorage) nomeFuncao = GetVeiculos() (parametros = não tem) (retorno = ([]Veiculo, error))
func (s *MySQLStorage) GetVeiculos() ([]Veiculo, error) {
	sql := "select id, nome, marca, ano, modelo from veiculos"
	rows, err := s.dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	//garante que será fechada a conexão ao término do método
	defer rows.Close()
	//define um slice (lista) de veiculos
	var veiculos []Veiculo

	for rows.Next() {
		//define variavel do tipo Veiculo
		var veiculo Veiculo
		//1 - Pega o resultSet (linhas (rows) que retornaram do banco)
		//2 - Pega o poteiro & da variavel Veiculo e armazena os dados nela
		rows.Scan(&veiculo.ID, &veiculo.Nome, &veiculo.Marca, &veiculo.Ano, &veiculo.Modelo)
		//3 - Pega o item (variavel Veiculo) e adiciona no slice
		veiculos = append(veiculos, veiculo)
	}
	return veiculos, nil
}
func (s *MySQLStorage) CreateVeiculo(nome, marca string, ano, modelo int) error {
	insert := "insert into veiculos (nome, marca, ano, modelo) values (?,?,?,?);"
	//prepara o banco para receber os parametros
	stmt, err := s.dbConn.Prepare(insert)
	if err != nil {
		return err
	}
	//garante que será fechada a conexão
	defer stmt.Close()
	//executa a query que estava preparada com os parametros
	_, err = stmt.Exec(nome, marca, ano, modelo)
	if err != nil {
		return err
	}
	//se tudo ocorrer bem retornará NIL Pointer, ou seja, s/ erro
	return nil
}
func (s *MySQLStorage) UpdateVeiculo(id int, veiculo *Veiculo) error {
	update := "update veiculos set nome=?, marca=?, ano=?, modelo=? where id=?;"
	stmt, err := s.dbConn.Prepare(update)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(veiculo.Nome, veiculo.Marca, veiculo.Ano, veiculo.Modelo, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *MySQLStorage) Delete(id int) error {
	deleteSQL := "delete from veiculos where id=?"
	stmt, err := s.dbConn.Prepare(deleteSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//construtor do MySQL Storage (conexão c/ o banco)
func NewStorage(conStr string) MySQLStorage {
	conn, err := sql.Open("mysql", conStr)
	if err != nil {
		panic("MySQL connection has failed!")
	}
	return MySQLStorage{
		dbConn: conn,
	}
}
