package main

import (
	"Calculator/internal/controller"
	calculator_v1 "Calculator/proto/gen/go/calc"
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
)

type ServerStruct struct {
	calculator_v1.UnimplementedCalculatorServer
	calculator Calculator
}

func SetViperConfig() {
	viper.SetEnvPrefix("calcg")
	viper.SetDefault("port", "8081")
	err := viper.BindEnv("port")
	if err != nil {
		slog.Info("Used default port", "port", viper.GetString("CalcPort"))
	}
	viper.AutomaticEnv()
}

type Calculator interface {
	Calc(payload string) (int32, string)
}

// CalculatorImpl is an implementation of the Calculator interface
type CalculatorImpl struct{}

// NewCalculatorImpl creates a new calculator implementation
func NewCalculatorImpl() *CalculatorImpl {
	return &CalculatorImpl{}
}

// Calc processes the calculation payload
func (c *CalculatorImpl) Calc(payload string) (int32, string) {
	ctx := context.Background()
	var input []map[string]interface{}
	err := json.Unmarshal([]byte(payload), &input)
	if err != nil {
		return 1, "Invalid JSON payload"
	}

	res, err := controller.CalculateSentence(ctx, input...)
	if err != nil {
		return 1, err.Error()
	}

	type Item struct {
		Var   string `json:"var"`
		Value int64  `json:"value"`
	}
	type Result struct {
		Items []Item `json:"items"`
	}
	result := &Result{
		make([]Item, 0),
	}

	for key, item := range res {
		result.Items = append(result.Items, Item{
			key,
			item,
		})
	}

	data, err := json.Marshal(result)
	if err != nil {
		return 1, "Failed to marshal result"
	}

	return 0, string(data)
}

func RegisterCalculator(grpcServer *grpc.Server, calculator Calculator) {
	calculator_v1.RegisterCalculatorServer(grpcServer, &ServerStruct{calculator: calculator})
}

func (s *ServerStruct) Calc(ctx context.Context, in *calculator_v1.Sentence) (*calculator_v1.Result, error) {
	if in.Payload == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Empty payload")
	}

	var input []map[string]interface{}
	err := json.Unmarshal([]byte(in.Payload), &input)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot unmarshal payload")
	}
	res, err := controller.CalculateSentence(ctx, input...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot calculate sentence, err=%s", err.Error())
	}

	type Item struct {
		Var   string `json:"var"`
		Value int64  `json:"value"`
	}
	type Result struct {
		Items []Item `json:"items"`
	}
	result := &Result{
		make([]Item, 0),
	}

	for key, item := range res {
		result.Items = append(result.Items, Item{
			key,
			item,
		})
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot marshal result, err=%s", err.Error())
	}

	return &calculator_v1.Result{Code: 0, Message: string(data)}, nil
}

func main() {
	SetViperConfig()

	grpcServer := grpc.NewServer()

	// Create calculator implementation
	calcImpl := NewCalculatorImpl()

	// Register calculator service
	RegisterCalculator(grpcServer, calcImpl)

	// Set up listener
	l, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		panic(err)
	}

	slog.Info("Starting gRPC Server", "port", viper.GetString("port"))

	// Start serving
	if err := grpcServer.Serve(l); err != nil {
		slog.Error("Failed to serve", "error", err)
		panic(err)
	}
}
