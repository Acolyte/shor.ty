package config

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rbcervilla/redisstore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"text/template"
	"time"
)

type ServerConfig struct {
	Host string
	Port int
}

type DBConnection struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
	Schema string
	Proto  string
	DSN    string
}

type RedisConnection struct {
	Host     string
	Port     int
	Password string
	DB       int
	Proto    string
}

type RedisConfig struct {
	ConnectionSettings RedisConnection
	Client             *redis.Client
}

type DatabaseConfig struct {
	ConnectionType     string
	ConnectionSettings DBConnection
	Connection         *sqlx.DB
}

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	Redis        RedisConfig
	SessionRedis *redisstore.RedisStore
}

var Settings *Config

func (dc *DatabaseConfig) GetDSN() (string, error) {
	switch dc.ConnectionType {
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dc.ConnectionSettings.Host,
			dc.ConnectionSettings.Port,
			dc.ConnectionSettings.User,
			dc.ConnectionSettings.Pass,
			dc.ConnectionSettings.Name,
		), nil
	case "mysql":
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8",
			dc.ConnectionSettings.User,
			dc.ConnectionSettings.Pass,
			dc.ConnectionSettings.Proto,
			dc.ConnectionSettings.Host,
			dc.ConnectionSettings.Port,
			dc.ConnectionSettings.Name,
		), nil
	}

	return "", errors.New("unknown database type")
}

// Helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func init() {
	// load .env file
	err := godotenv.Load(".env")

	Templates = make(map[string]*template.Template)
	Templates["index"] = template.Must(template.ParseFiles("./web/template/index.html"))
	Templates["found"] = template.Must(template.ParseFiles("./web/template/found.html"))

	Settings = &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_PORT", 80),
		},

		Database: DatabaseConfig{
			ConnectionType: getEnv("DB_CONNECTION", "postgres"),
			ConnectionSettings: DBConnection{
				Host:   getEnv("DB_HOST", "127.0.0.1"),
				Port:   getEnvAsInt("DB_PORT", 5432),
				User:   getEnv("DB_USERNAME", "db"),
				Pass:   getEnv("DB_PASSWORD", "db"),
				Name:   getEnv("DB_DATABASE", "db"),
				Schema: getEnv("DB_SCHEMA", "public"),
				Proto:  getEnv("DB_PROTOCOL", "tcp"),
			},
		},
		Redis: RedisConfig{
			ConnectionSettings: RedisConnection{
				Host:     getEnv("REDIS_HOST", "127.0.0.1"),
				Port:     getEnvAsInt("REDIS_PORT", 6379),
				Password: getEnv("REDIS_PASSWORD", ""),
				DB:       getEnvAsInt("REDIS_DB", 0),
				Proto:    getEnv("REDIS_PROTOCOL", "tcp"),
			},
		},
	}

	dsn, err := Settings.Database.GetDSN()
	if err != nil {
		panic(err)
	}
	conn, err := sqlx.Connect(Settings.Database.ConnectionType, dsn)
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)
	Settings.Database.Connection = conn
	Database = conn

	Gorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", Settings.Redis.ConnectionSettings.Host, Settings.Redis.ConnectionSettings.Port),
		Password:    Settings.Redis.ConnectionSettings.Password,
		DB:          Settings.Redis.ConnectionSettings.DB,
		IdleTimeout: 5 * time.Minute,
		MaxRetries:  3,
	})

	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	Settings.Redis.Client = client

	Settings.SessionRedis, err = redisstore.NewRedisStore(client)
	if err != nil {
		panic(err)
	}
}
