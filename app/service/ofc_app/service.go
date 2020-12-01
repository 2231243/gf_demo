package ofc_app

import (
	"context"
	"time"

	"gf-app/grpc/ofc_app"
	//"gf-app/library"
)

//应用添加
func AppCreate(req *ofc_app.AppRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.AppReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.AppAdd(ctx, req)
}

//应用修改
func AppUpdate(req *ofc_app.AppRequest, client ofc_app.OfcAppServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.AppSave(ctx, req)
	return err
}

//获取单个应用信息
func AppOne(req *ofc_app.AppGetRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.AppOneReply, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return client.OneApp(ctx, req)
}

//获取应用列表
func AppList(req *ofc_app.AppGetRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.AppListReply, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return client.ListApp(ctx, req)
}

//刷新密钥
func AppSecret(req *ofc_app.AppKeyRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.SecretSaveReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.SecretSave(ctx, req)
}

//消息抄送
func AppMessage(req *ofc_app.MessageUrlRequest, client ofc_app.OfcAppServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.MessageUrl(ctx, req)
	return err
}

//标识管理
func AppIdentifier(req *ofc_app.IdentifierRequest, client ofc_app.OfcAppServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Identifier(ctx, req)
	return err
}

//文案添加
func CopyWritingCreate(req *ofc_app.CopywritingRequest, client ofc_app.OfcAppServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.CopywritingAdd(ctx, req)
	return err
}

//文案修改
func CopyWritingUpdate(req *ofc_app.CopywritingRequest, client ofc_app.OfcAppServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.CopywritingSave(ctx, req)
	return err
}

//获取单个文案信息
func CopywritingGetOne(req *ofc_app.CopywritingGetRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.CopywritingOneReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.CopywritingOne(ctx, req)
}

//获取文案列表
func CopywritingGetList(req *ofc_app.CopywritingGetRequest, client ofc_app.OfcAppServiceClient) (*ofc_app.CopywritingListReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.CopywritingList(ctx, req)
}
