package config

const (

	//
	// MySQLUsername is the username of the database to connect to.
	//
	MySQLUsername string = "root"

	//
	// MySQLPassword is the password of the database to connect to.
	//
	MySQLPassword string = "root"

	//
	// MySQLDatabase is the name of the database to connect to.
	//
	MySQLDatabase string = "noteshare"

	//
	// MySQLPort is the port number to connect to the database on.
	//
	MySQLPort string = "3306"
)

//
// MySQLToFormatDNS converts the mysql data options to string format.
//
func MySQLToFormatDNS() string {
	return MySQLUsername +
		":" + MySQLPassword +
		"@/" + MySQLDatabase
}
