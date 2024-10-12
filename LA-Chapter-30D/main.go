// package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// func main() {
// 	// 1. Menetapkan nilai default
// 	viper.SetDefault("app_name", "DefaultApp")
// 	viper.SetDefault("port", 8080)
// 	viper.SetDefault("debug", false)
// 	viper.SetDefault("database.host", "localhost")
// 	viper.SetDefault("database.port", 5432)

// 	// 2. Mengatur file konfigurasi (misalnya, config.yaml) jika ada
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")

// 	// 3. Membaca file konfigurasi jika ada
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Println("No config file found, using defaults")
// 	}

// 	// 4. Mengakses nilai konfigurasi
// 	appName := viper.GetString("app_name")
// 	port := viper.GetInt("port")
// 	debug := viper.GetBool("debug")

// 	dbHost := viper.GetString("database.host")
// 	dbPort := viper.GetInt("database.port")

// 	// 5. Menampilkan nilai konfigurasi
// 	fmt.Printf("App Name: %s\n", appName)
// 	fmt.Printf("Port: %d\n", port)
// 	fmt.Printf("Debug: %t\n", debug)
// 	fmt.Printf("Database Host: %s\n", dbHost)
// 	fmt.Printf("Database Port: %d\n", dbPort)
// }

// package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// func main() {
// 	// 1. Mengatur Viper untuk membaca dari file JSON
// 	viper.SetConfigName("config1") // Nama file tanpa ekstensi
// 	viper.SetConfigType("json")    // Format file
// 	viper.AddConfigPath(".")       // Jalur ke file konfigurasi

// 	// 2. Membaca file konfigurasi JSON
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Println("Error reading JSON config:", err)
// 	}

// 	// 3. Mengatur Viper untuk membaca dari file YAML
// 	viper.SetConfigName("config2") // Nama file tanpa ekstensi
// 	viper.SetConfigType("yaml")    // Format file
// 	viper.AddConfigPath(".")       // Jalur ke file konfigurasi

// 	// 4. Membaca file konfigurasi YAML
// 	if err := viper.MergeInConfig(); err != nil {
// 		fmt.Println("Error reading YAML config:", err)
// 	}

// 	// 5. Mengakses nilai konfigurasi
// 	appName := viper.GetString("app_name")
// 	port := viper.GetInt("port")
// 	debug := viper.GetBool("debug")

// 	dbHost := viper.GetString("database.host")
// 	dbPort := viper.GetInt("database.port")

// 	// 6. Menampilkan nilai konfigurasi
// 	fmt.Printf("App Name: %s\n", appName)
// 	fmt.Printf("Port: %d\n", port)
// 	fmt.Printf("Debug: %t\n", debug)
// 	fmt.Printf("Database Host: %s\n", dbHost)
// 	fmt.Printf("Database Port: %d\n", dbPort)
// }

package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// 1. Menetapkan flags menggunakan paket flag
	flag.Int("port", 8080, "Port for the application")
	flag.Bool("debug", false, "Enable debug mode")
	flag.String("app_name", "MyApp", "Name of the application")

	// 2. Memparsing flags
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// 4. Mengakses nilai flags melalui Viper
	port := viper.GetInt("port")
	debug := viper.GetBool("debug")
	appName := viper.GetString("app_name")

	// 5. Menampilkan nilai flags
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug Mode: %t\n", debug)
}
