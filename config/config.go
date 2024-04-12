package config

// Importando um leitor de arquivos .TOML
import "github.com/spf13/viper"

// Definindo uma variavel Ponteiro da Struct de configuração do bando
var cfg *configPostgres
// var cfm *configMysql

// Struct de Configuraçao
type configPostgres struct {
	DB DBConfigPostgres
}
// type configMysql struct {
// 	DB DBConfigMysql
// }

// Struct de Conexão
type DBConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// type DBConfigMysql struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Pass     string
// 	Database string
// }

// Inicializando o leitor/Viper e de finindo algumas variaveis com valores padrao
func init() {
	viper.SetDefault("postgresql.host", "localhost")
	viper.SetDefault("postgresql.port", "5432")
	// viper.SetDefault("mysql.host", "localhost")
	// viper.SetDefault("mysql.port", "3306")
}

// Carregando o arquivo .TOML
func Load() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// Recuperando o Ponteiro
	cfg = new(configPostgres)
	// cfm = new(configMysql)

	// Preenchendo as variaveis da Struct do Ponteiro com os valores recuperando do arquivo .TOML
	cfg.DB = DBConfigPostgres{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetString("postgresql.port"),
		User:     viper.GetString("postgresql.user"),
		Pass:     viper.GetString("postgresql.pass"),
		Database: viper.GetString("postgresql.name"),
	}
	// cfm.DB = DBConfigMysql{
	// 	Host:     viper.GetString("mysql.host"),
	// 	Port:     viper.GetString("mysql.port"),
	// 	User:     viper.GetString("mysql.user"),
	// 	Pass:     viper.GetString("mysql.pass"),
	// 	Database: viper.GetString("mysql.name"),
	// }

	// Retornando o ponteiro
	return nil
}

// Exportando a conexao com o Database
func GetDB() DBConfigPostgres {
	return cfg.DB //, cfm.DB
}
