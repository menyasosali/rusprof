package main

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"os"
	pb "rusprofile/rusprofile/proto"
	"strings"
)

const (
	grpcServerEndpoint = ":50051"
	searchURL          = "https://www.rusprofile.ru/search?query="
)

// server represents the gRPC server for RusProfile service.
type server struct {
	pb.UnimplementedRusProfileServiceServer
}

// GetCompanyInfo retrieves information about a company by INN.
//
// swagger:route GET /companies/{inn} getCompanyInfo
//
// Retrieves information about a company by INN.
//
// Parameters:
//   - name: inn
//     in: path
//     description: INN (Taxpayer Identification Number) of the company.
//     required: true
//     schema:
//     type: string
//
// Responses:
//
//	200: companyResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
func (s *server) GetCompanyInfo(ctx context.Context, req *pb.CompanyRequest) (*pb.CompanyResponse, error) {
	inn := req.GetInn()
	companyURL := searchURL + inn

	resp, err := http.Get(companyURL)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make a request to rusprofile.ru")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.NotFound, "company not found on rusprofile.ru")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse HTML response")
	}

	companyName := doc.Find("h1[itemprop='name']").Text()
	companyName = strings.TrimSpace(companyName)

	if companyName == "" {
		return nil, status.Errorf(codes.NotFound, "name of company not found on rusprofile.ru")
	}

	companyKpp := doc.Find("span#clip_kpp").Text()
	companyKpp = strings.TrimSpace(companyKpp)

	if companyKpp == "" {
		return nil, status.Errorf(codes.NotFound, "kpp of company not found")
	}

	companyGlavar := doc.Find("div.company-row.hidden-parent span.company-info__text a.link-arrow").Text()
	companyGlavar = strings.TrimSpace(companyGlavar)

	if companyGlavar == "" {
		return nil, status.Errorf(codes.NotFound, "glavar of company not found")
	}

	return &pb.CompanyResponse{
		Inn:    inn,
		Kpp:    companyKpp,
		Name:   companyName,
		Glavar: companyGlavar,
	}, nil
}

func readSwaggerFile() ([]byte, error) {
	filePath := "swagger/proto/rusprofile.swagger.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *server) GetSwaggerUI(ctx context.Context, req *pb.SwaggerUIRequest) (*pb.SwaggerUIResponse, error) {
	swaggerData, err := readSwaggerFile()
	if err != nil {
		return nil, err
	}

	html := string(swaggerData)

	return &pb.SwaggerUIResponse{
		Html: html,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRusProfileServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
