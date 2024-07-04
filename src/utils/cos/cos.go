package cos

import (
	"context"
	"singapore/src/utils/config"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"io"
	"net/http"
	"net/url"
)

// var cosClient *cos.Client
var cosClient *Client
var xCosClient *Client

type Client struct {
	Cli *cos.Client
	Env string
	Cdn string
}

func InitCos(c *config.TencentOSS) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	cosClient = new(Client)
	u, _ := url.Parse(c.Url)
	b := &cos.BaseURL{BucketURL: u}
	cosClient.Cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: c.SecretId,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: c.SecretKey,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	cosClient.Env = c.Env
	cosClient.Cdn = c.Cdn

	if cosClient.Cli == nil {
		panic(fmt.Sprintf("InitCos fail config: %+v", c))
	}

}

// x项目私有桶
func InitCosXClient(c *config.TencentOSS) *Client {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	xCosClient = new(Client)
	u, _ := url.Parse(c.Url)
	b := &cos.BaseURL{BucketURL: u}
	xCosClient.Cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: c.SecretId,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: c.SecretKey,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	xCosClient.Env = c.Env
	xCosClient.Cdn = c.Cdn

	if xCosClient.Cli == nil {
		panic(fmt.Sprintf("xCosClient InitCos fail config: %+v", c))
	}

	return xCosClient
}

// 获取client
func GetCli() *Client {
	return cosClient
}

// 获取client
func GetXCli() *Client {
	return xCosClient
}

// 上传对象
func (c *Client) Put(ctx context.Context, name string, r io.Reader, opt *cos.ObjectPutOptions) (*cos.Response, error) {
	return c.Cli.Object.Put(ctx, name, r, opt)
}

// 上传文件对象
func (c *Client) PutFromFile(ctx context.Context, name string, path string, opt *cos.ObjectPutOptions) (*cos.Response, error) {
	return c.Cli.Object.PutFromFile(ctx, name, path, opt)
}

// 判断文件存在
func (c *Client) IsExist(ctx context.Context, name string, id ...string) (bool, error) {
	return c.Cli.Object.IsExist(ctx, name, id...)
}

// 删除对象
func (c *Client) Delete(ctx context.Context, name string) (*cos.Response, error) {
	return c.Cli.Object.Delete(ctx, name)
}

// 删除对象
func (c *Client) Copy(ctx context.Context, name string, sourceURL string) (*cos.ObjectCopyResult, *cos.Response, error) {
	return c.Cli.Object.Copy(ctx, name, sourceURL, nil)
}

// 下载对象
func (c *Client) Download(ctx context.Context, name string, filepath string) (*cos.Response, error) {
	opts := &cos.MultiDownloadOptions{
		PartSize:       16,
		ThreadPoolSize: 20,
	}
	return c.Cli.Object.Download(ctx, name, filepath, opts)
}

// 查询媒体服务可用
func (c *Client) DescribeMediaProcessBuckets(ctx context.Context) (*cos.DescribeMediaProcessBucketsResult, *cos.Response, error) {
	opt := &cos.DescribeMediaProcessBucketsOptions{
		Regions:     "ap-beijing",
		BucketNames: "",
		BucketName:  "",
	}

	return c.Cli.CI.DescribeMediaProcessBuckets(ctx, opt)
}

// 获取媒体截图
func (c *Client) GetSnapshot(ctx context.Context, name string) (*cos.Response, error) {
	var opt *cos.GetSnapshotOptions = &cos.GetSnapshotOptions{
		Time:   1.0,
		Height: 0,
		Width:  0,
		Format: "png",
		Rotate: "auto",
		Mode:   "exactframe",
		// Mode:   "keyframe",
	}

	return c.Cli.CI.GetSnapshot(ctx, name, opt)
}

func (c *Client) GetSnapshot2(ctx context.Context, url string, height int) (*cos.Response, error) {
	var opt *cos.GetSnapshotOptions = &cos.GetSnapshotOptions{
		Time:   1.0,
		Height: height,
		Width:  0,
		Format: "jpg",
		Rotate: "auto",
		Mode:   "keyframe",
		// Mode:   "keyframe",
	}

	return c.Cli.CI.GetSnapshot(ctx, url, opt)
}
