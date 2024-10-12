// package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// func main() {
// 	// 1. Menentukan nama file konfigurasi
// 	viper.SetConfigName(".env")

// 	// 2. Menentukan format file konfigurasi sebagai env
// 	viper.SetConfigType("env")

// 	// 3. Menentukan jalur direktori tempat file konfigurasi berada
// 	viper.AddConfigPath(".")

// 	// 4. Membaca file konfigurasi
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Printf("Error membaca file konfigurasi: %s\n", err)
// 		return
// 	}

// 	// 5. Mengakses nilai konfigurasi
// 	appName := viper.GetString("APP_NAME")
// 	port := viper.GetInt("PORT")
// 	debug := viper.GetBool("DEBUG")

// 	dbUser := viper.GetString("DATABASE_USER")
// 	dbPassword := viper.GetString("DATABASE_PASSWORD")
// 	dbHost := viper.GetString("DATABASE_HOST")
// 	dbPort := viper.GetInt("DATABASE_PORT")
// 	dbName := viper.GetString("DATABASE_NAME")

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
	// 1. Mengaktifkan pembacaan variabel lingkungan secara otomatis
	viper.AutomaticEnv()

	// 3. Mengakses nilai konfigurasi dari variabel lingkungan
	appName := viper.GetString("APP_NAME")
	port := viper.GetInt("PORT")
	debug := viper.GetBool("DEBUG")

	dbUser := viper.GetString("DATABASE_USER")
	dbPassword := viper.GetString("DATABASE_PASSWORD")
	dbHost := viper.GetString("DATABASE_HOST")
	dbPort := viper.GetInt("DATABASE_PORT")
	dbName := viper.GetString("DATABASE_NAME")

	// 4. Menampilkan nilai konfigurasi
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug: %t\n", debug)
	fmt.Printf("Database User: %s\n", dbUser)
	fmt.Printf("Database Password: %s\n", dbPassword)
	fmt.Printf("Database Host: %s\n", dbHost)
	fmt.Printf("Database Port: %d\n", dbPort)
	fmt.Printf("Database Name: %s\n", dbName)
}
