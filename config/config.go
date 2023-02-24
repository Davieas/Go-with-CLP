package config

import(

"github.com/spf13/viper"


)

var cfg *config

type config struct {

	Client TcpClient
	MODBUS Adress
	Rack Location
	Slot Number

}

type TcpClient struct {
	IP string
}

type Adress struct {

	adressInput string
	adressOutput string

}
type Rack struct {

	Location int


}


type Slot struct {

	Number int
	

}

func init(){


	
}
func load(){
	viper.SetConfigName("config");
	viper.SetConfigType("toml");
	viper.AddConfigPath(".");

	err := viper.ReadInConfig();

	if err != nil {
		if err _, ok := err.(viper.ConfigFileNotFoundError); !ok {

			return err

		}
	}

	cfg = new(config)

	cfg.Client = TcpClient{
		IP: viper.GetString("tcp.client"),
	}

	cfg.MODBUS = Adress{
		adressInput: viper.GetString("modbus.adressIn")
		adressOutput: viper.GetString("modbus.adressOut")
	}

	cfg.Rack = Location{

		Location: viper.GetString("rack.location")
	}

	cfg.Slot = Number{

		Number: viper.GetString("number.cpu")

	}

	return nil
}

func GetClient() TcpClient{
	return cfg.Client.IP
}

func GetModbusAdress() string {
	return cfg.MODBUS.adressInput
	return cfg.MODBUS.adressOutput
}

func GetRackLocation() string {


	return cfg.Rack.Location
}

func GetSlotNumber() string {


	return cfg.Slot.Number
}


