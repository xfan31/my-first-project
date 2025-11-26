package config

type Config struct {
	DBHost     string //数据库服务器地址
	DBPort     string //数据库端口号
	DBUser     string //数据库用户名
	DBPassword string //数据库密码
	DBName     string //要连接的数据库名称
	JWTSecret  string
} //把项目中所有需要配置的参数都放在一个地方，方便统一管理

// 返回指针，避免数据拷贝，更高效
func Load() *Config { //返回指针
	//创建Config结构体的实例
	return &Config{ //取地址
		DBHost:     "localhost",
		DBPort:     "3306",
		DBUser:     "root",
		DBPassword: "123456",
		DBName:     "login_demo",
		JWTSecret:  "your-secret-key",
	}
} //Load()函数是一个配置初始化器，它封装了创建配置对象的复杂过程
