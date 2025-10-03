package main

import (
	// "context"
	"log"
	"net"
	// "net/http"
	// "os"
	// "os/signal"
	// "time"

	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"

	// "github.com/AchrafMoreau/gRPC-in-Go/server/handlers"
	// "github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const(
	port = ":9090"
)

type HelloServer struct{
	pb.GreetServiceServer
}

func main() {

	l, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Failed to Start the server at the Port : %v", port)
	}

	grpcServer := grpc.NewServer()
	log.Printf("The Server is Listening on the Port %v", port)

	pb.RegisterGreetServiceServer(grpcServer, &HelloServer{})
	err = grpcServer.Serve(l);
	if err != nil{
		log.Fatalf("Failed to Start the server at the Port : %v \n Error: %v", port, err)
	}


	// You can use print statements as follows for debugging, they'll be visible when running tests.
//	l := log.New(os.Stdout, "product-api", log.LstdFlags)
//	productHnalder := handlers.NewProduct(l)
//
//	sm := mux.NewRouter()
//	getRouter := sm.Methods(http.MethodGet).Subrouter()
//	getRouter.HandleFunc("/", productHnalder.GetAllProducts)
//
//	putRouter := sm.Methods(http.MethodPut).Subrouter()
//	putRouter.HandleFunc("/{id:[0-9]+}", productHnalder.UpdateProduct)
//	putRouter.Use(productHnalder.MiddlewareProduct)
//
//	postRouter := sm.Methods(http.MethodPost).Subrouter()
//	postRouter.HandleFunc("/", productHnalder.AddProducts)
//	postRouter.Use(productHnalder.MiddlewareProduct)
//
//	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
//	deleteRouter.HandleFunc("/{id:[0-9]+}", productHnalder.DeleteProduct)
//
//	server := &http.Server{
//		Addr: ":9090",
//		Handler: sm,
//		IdleTimeout: 120*time.Second,
//		ReadTimeout: 1*time.Second,
//		WriteTimeout: 1*time.Second,
//	}
//
//	go func() {
//
//		err := http.ListenAndServe(":9090", sm)
//		if err != nil{
//			l.Fatal(err)
//		}
//	}()
//
//	sigChan := make(chan os.Signal)
//	signal.Notify(sigChan, os.Interrupt)
//	signal.Notify(sigChan, os.Kill)
//
//	sig := <- sigChan;
//	l.Println("Recived Terminate, gracful shuddown !", sig);
//
//	cx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//	server.Shutdown(cx)
}
