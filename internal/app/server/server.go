package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	balance_handler "github.com/avito-test-case/internal/pkg/balance/handler"
	balance_repository "github.com/avito-test-case/internal/pkg/balance/repository"
	balance_usecase "github.com/avito-test-case/internal/pkg/balance/usecase"
	transaction_handler "github.com/avito-test-case/internal/pkg/transaction/handler"
	transaction_repository "github.com/avito-test-case/internal/pkg/transaction/repository"
	transaction_usecase "github.com/avito-test-case/internal/pkg/transaction/usecase"
	"github.com/avito-test-case/pkg/tools/configer"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Start() {
	// Init config
	configer.Init("configs/app/balance_service.yaml")

	// Connect to postgreSql db
	postgreSqlConn, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
			configer.AppConfig.Postgresql.User,
			configer.AppConfig.Postgresql.Password,
			configer.AppConfig.Postgresql.DBName,
			configer.AppConfig.Postgresql.Host,
			configer.AppConfig.Postgresql.Port,
			configer.AppConfig.Postgresql.Sslmode,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer postgreSqlConn.Close()
	if err := postgreSqlConn.Ping(); err != nil {
		log.Fatal(err)
	}

	balanceRepository := balance_repository.NewSessionPostgresqlRepository(postgreSqlConn)
	transactionRepository := transaction_repository.NewSessionPostgresqlRepository(postgreSqlConn)

	balanceUseCase := balance_usecase.NewUseCase(balanceRepository)
	transactionUseCase := transaction_usecase.NewUseCase(transactionRepository)

	balanceHandler := balance_handler.NewHandler(balanceUseCase)
	transactionHandler := transaction_handler.NewHandler(transactionUseCase)

	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/api/v1/user/balance/{id:[0-9]+}", balanceHandler.GetUserBalance).Methods("GET", "OPTIONS")
	mainRouter.HandleFunc("/api/v1/user/balance/improve", balanceHandler.ImproveUserBalance).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/v1/user/balance/withdraw", balanceHandler.WithdrawUserBalance).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/v1/user/balance/transfer", balanceHandler.MakeUserBalanceTransfer).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/v1/user/transaction/{id:[0-9]+}", transactionHandler.GetUserTransactions).Methods("GET", "OPTIONS")

	server := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%s",
			configer.AppConfig.Server.Host,
			configer.AppConfig.Server.Port,
		),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mainRouter,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
