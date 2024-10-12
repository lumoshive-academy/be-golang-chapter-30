// package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// func main() {
// 	// 1. Menentukan nama file konfigurasi (tanpa ekstensi)
// 	viper.SetConfigName("config")

// 	// 2. Menentukan jalur direktori tempat file konfigurasi berada
// 	viper.AddConfigPath(".")

// 	// 3. Menentukan format file konfigurasi sebagai JSON
// 	viper.SetConfigType("json")

// 	// 4. Membaca file konfigurasi
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Printf("Error membaca file konfigurasi: %s\n", err)
// 		return
// 	}

// 	// 5. Mengakses nilai konfigurasi
// 	appName := viper.GetString("app_name")
// 	port := viper.GetInt("port")
// 	debug := viper.GetBool("debug")

// 	dbUser := viper.GetString("database.user")
// 	dbPassword := viper.GetString("database.password")
// 	dbHost := viper.GetString("database.host")
// 	dbPort := viper.GetInt("database.port")
// 	dbName := viper.GetString("database.name")

// 	// 6. Menampilkan nilai konfigurasi
// 	fmt.Printf("App Name: %s\n", appName)
// 	fmt.Printf("Port: %d\n", port)
// 	fmt.Printf("Debug: %t\n", debug)
// 	fmt.Printf("Database User: %s\n", dbUser)
// 	fmt.Printf("Database Password: %s\n", dbPassword)
// 	fmt.Printf("Database Host: %s\n", dbHost)
// 	fmt.Printf("Database Port: %d\n", dbPort)
// 	fmt.Printf("Database Name: %s\n", dbName)
// }

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 1. Menentukan nama file konfigurasi (tanpa ekstensi)
	viper.SetConfigName("config")

	// 2. Menentukan jalur direktori tempat file konfigurasi berada
	viper.AddConfigPath(".")

	// 3. Menentukan format file konfigurasi sebagai YAML
	viper.SetConfigType("yaml")

	// 4. Membaca file konfigurasi
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error membaca file konfigurasi: %s\n", err)
		return
	}

	// 5. Mengakses nilai konfigurasi
	appName := viper.GetString("app_name")
	port := viper.GetInt("port")
	debug := viper.GetBool("debug")

	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbName := viper.GetString("database.name")

	// 6. Menampilkan nilai konfigurasi
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug: %t\n", debug)
	fmt.Printf("Database User: %s\n", dbUser)
	fmt.Printf("Database Password: %s\n", dbPassword)
	fmt.Printf("Database Host: %s\n", dbHost)
	fmt.Printf("Database Port: %d\n", dbPort)
	fmt.Printf("Database Name: %s\n", dbName)
}
