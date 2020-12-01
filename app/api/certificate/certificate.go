package certificate

import (
	"gf-app/app/api"
	service "gf-app/app/service/certificate"
	"gf-app/boot"
	pb "gf-app/grpc/ofc_certificate"
	"gf-app/inter"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

type Cert struct {
	api.Control
	certClient pb.OfcCertificateServiceClient
	conn inter.RpcInter
}

func (c *Cert) Init(r *ghttp.Request) {
	//c.Control.Init(r)
	c.conn, _ = boot.RpcPoolMap.Conn(boot.OFC_COM_SVR)
	c.certClient = pb.NewOfcCertificateServiceClient(c.conn.GetRpc())
}

func (c *Cert) Shut(r *ghttp.Request) {
	c.conn.Close()
}

/**
 * @Author lifeifei
 * @Date 9:34 上午 2020/11/18
 * @Summary 证书模块
 * @Tags 上传iOS证书
 * @Produce json
 * @Router  /api/v1/cert/upload
 * @Param fileList formData file true "上传文件"
 * @Return 0 {object} response.JsonResponse "执行结果"
 **/
func (c *Cert) Upload(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		c.Fail(r, gerror.New("文件不存在"))
	}
	fileName, err := file.Save(gfile.TempDir(), true)
	if err != nil {
		c.Fail(r, err)
	}
	c.Success(r, g.MapAnyAny{"filename": gfile.TempDir() + "/" + fileName})
}

//添加IOS证书
func (c *Cert) Ios(r *ghttp.Request) {
	var req *pb.IosCertificateRequest
	if err := c.Parse(r, &req); err != nil {
		c.Fail(r, err, nil, req)
	}
	//req := boot.M.Get(r.URL.Path).(*pb.IosCertificateRequest)
	if err := service.IosCreate(req, c.certClient); err != nil {
		c.Fail(r, err, nil, req)
	}
	c.Success(r, nil)
}

//添加android证书
func (c *Cert) Android(r *ghttp.Request) {
	var request *pb.AndroidCertificateRequest
	if err := c.Parse(r, &request); err != nil {
		c.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*pb.AndroidCertificateRequest)
	if err := service.AndroidCreate(request, c.certClient); err != nil {
		c.Fail(r, err, nil, request)
	}
	c.Success(r, nil, request)
}

//获取单个证书信息
func (c *Cert) One(r *ghttp.Request) {
	var request *pb.OneCertificateRequest
	if err := c.Parse(r, &request); err != nil {
		c.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*pb.OneCertificateRequest)
	data, err := service.OneGet(request, c.certClient)
	if err != nil {
		c.Fail(r, err, nil, request)
	}
	c.Success(r, data, nil)
}

//获取证书列表
func (c *Cert) List(r *ghttp.Request) {
	var request *pb.CertificateListRequest
	if err := c.Parse(r, &request); err != nil {
		c.Fail(r, err, nil, request)
	}
	data, err := service.ListGet(request, c.certClient)
	if err != nil {
		c.Fail(r, err, nil, request)
	}
	c.Success(r, data, nil, request)
}

//删除证书
func (c *Cert) Del(r *ghttp.Request) {
	var request *pb.IdCertificateRequest
	if err := c.Parse(r, &request); err != nil {
		c.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*pb.IdCertificateRequest)
	if err := service.DelCert(request, c.certClient); err != nil {
		c.Fail(r, err, nil, request)
	}
	c.Success(r, nil)
}

func NewCert() *Cert {
	cert := &Cert{}
	return cert
}
