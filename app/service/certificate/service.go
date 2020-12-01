package certificate

import (
	"context"
	"encoding/json"
	"time"

	pb "gf-app/grpc/ofc_certificate"

	"github.com/gogf/gf/util/gconv"
)

//ios添加
func IosCreate(req *pb.IosCertificateRequest, client pb.OfcCertificateServiceClient) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.IosCertificateAdd(ctx, req)
	return err
}

//Android添加
func AndroidCreate(req *pb.AndroidCertificateRequest, client pb.OfcCertificateServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.AndroidCertificateAdd(ctx, req)
	return err
}

//删除证书
func DelCert(req *pb.IdCertificateRequest, client pb.OfcCertificateServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.DelCertificate(ctx, req)
	return err
}

//获取单个证书信息
func OneGet(req *pb.OneCertificateRequest, client pb.OfcCertificateServiceClient) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := client.OneCertificate(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := getInfo(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//获取证书列表
func ListGet(req *pb.CertificateListRequest, client pb.OfcCertificateServiceClient) ([]map[string]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	data, err := client.CertificateList(ctx, req)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	for _, v := range data.List {
		temp, err := getInfo(v)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}

func getInfo(in *pb.OneCertificateReply) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal(in.CertificateData.Value, &m)
	if err != nil {
		return nil, err
	}
	temp := gconv.Map(in)
	temp["certificate_data"] = m
	return temp, nil
}
