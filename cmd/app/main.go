package main

import (
	"booking-system-server-lab/internal/adapter/controller"
	"booking-system-server-lab/internal/adapter/controller/handler"
	application "booking-system-server-lab/internal/application/service"
	"booking-system-server-lab/internal/domain/factory"
	domain "booking-system-server-lab/internal/domain/service"
	"booking-system-server-lab/internal/infrastructure/database"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	db, err := database.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// DB接続を渡してUserRepositoryのインスタンスを生成
	userRepository := database.NewUserRepository(db)

	// UserFactoryのインスタンスを生成
	userFactory := factory.NewUserFactory()

	// UserServiceのインスタンスを作成
	domainService := domain.NewUserService(userRepository)

	// ユースケース層のインスタンスを作成（実装はService層に記載している）
	userUsecase := application.NewUserService(userRepository, userFactory, domainService)

	// ユースケース層のインスタンスをハンドラに渡す
	userHandler := handler.NewUserHandler(userUsecase)

	handlersDependency := controller.HandlersDependency{
		UserHandler: userHandler,
	}

	// ルーティングやミドルウェアの設定を行い、Echoの*echo.Echoインスタンスを返す
	e := controller.SetupRoutes(&handlersDependency)

	// Echoのサーバーをポート8080で起動
	// 指定したポートが使用可能でない場合や、サーバーが正常に起動できなかった場合はエラーが発生
	// e.Start("") でコンテナ内のポートを指定する
	e.Logger.Fatal(e.Start(":8080"))
}
