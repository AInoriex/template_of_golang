package obs

import (
	"context"
	"fmt"
	"gitee.com/slopy/tools/log"
	obsSdk "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"singapore/src/utils/config"
	"strings"
)

// var obsClient *obs.Client
var obsClient *Client

type Client struct {
	Cli    *obsSdk.ObsClient
	Env    string
	Cdn    string
	Bucket string
}

func InitObs(c *config.HuaweiOBS) (err error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/obs5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	//推荐通过环境变量获取AKSK，这里也可以使用其他外部引入方式传入，如果使用硬编码可能会存在泄露风险。
	//您可以登录访问管理控制台获取访问密钥AK/SK，获取方式请参见https://support.huaweicloud.com/usermanual-ca/ca_01_0003.html。
	// 【可选】如果使用临时AK/SK和SecurityToken访问OBS，同样建议您尽量避免使用硬编码，以降低信息泄露风险。您可以通过环境变量获取访问密钥AK/SK，也可以使用其他外部引入方式传入。
	// securityToken := os.Getenv("SecurityToken")
	// endpoint填写Bucket对应的Endpoint, 这里以华北-北京四为例，其他地区请按实际情况填写。
	obsClient = new(Client)
	endPoint := "https://obs.cn-north-4.myhuaweicloud.com"
	// 创建obsClient实例
	// 如果使用临时AKSK和SecurityToken访问OBS，需要在创建实例时通过obs.WithSecurityToken方法指定securityToken值。
	cli, err := obsSdk.New(c.SecretId, c.SecretKey, endPoint /*, obs.WithSecurityToken(securityToken)*/)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
		panic(fmt.Sprintf("InitObs fail config: %+v", c))
		// return err
	}
	obsClient.Env = c.Env
	obsClient.Cdn = c.Cdn
	obsClient.Cli = cli
	obsClient.Bucket = c.Bucket
	return nil
}

// 获取client
func GetCli() *Client {
	return obsClient
}

// 断点续传
// name: tmp/xxx.jpg 不要带上
// https://support.huaweicloud.com/sdk-go-devg-obs/obs_23_0409.html
func (c *Client) Upload(ctx context.Context, name string, path string) (err error) {

	// name 格式化成tmp/xxx.jpg
	if strings.Index(name, "/") == 0 {
		name = strings.TrimLeft(name, "/")
	}

	input := &obsSdk.UploadFileInput{}
	// 指定存储桶名称
	input.Bucket = c.Bucket
	// 指定上传对象，此处以 example/objectname 为例。
	input.Key = name
	// 指定待上传的本地文件，此处以/tmp/objectname为例。
	input.UploadFile = path
	// 指定是否开启断点续传模式，此处以true为例。默认为False，表示不开启。
	input.EnableCheckpoint = true
	// 指定分段大小，单位字节。此处以每段9M为例。
	input.PartSize = 9 * 1024 * 1024
	// 指定分段上传时的最大并发数，此处以并发数5为例
	input.TaskNum = 5
	// 断点续传上传对象
	output, err := c.Cli.UploadFile(input)
	if err == nil {
		log.Infof("Upload file(%s) under the bucket(%s) successful!\n", input.UploadFile, input.Bucket)
		log.Infof("ETag:%s\n", output.ETag)
		return nil
	}
	if obsError, ok := err.(obsSdk.ObsError); ok {
		//log.Errorf("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		log.Errorf("ObsError err:%+v", obsError)
	} else {
		// An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network
		log.Errorf("Exception err:%+v", err)
	}
	return err
}

func (c *Client) Download(ctx context.Context, name string, filepath string) (err error) {
	log.Infof("obs.Download(%s) start filepath:%s", name, filepath)
	// name 格式化成tmp/xxx.jpg
	if strings.Index(name, "/") == 0 {
		name = strings.TrimLeft(name, "/")
	}

	input := &obsSdk.DownloadFileInput{}
	// 指定存储桶名称
	input.Bucket = c.Bucket
	// 指定下载对象，此处以 example/objectname 为例。
	input.Key = name
	// 指定下载对象的本地文件全路径，此处以/tmp/objectname为例。当该值为空时，默认为当前程序的运行目录。
	input.DownloadFile = filepath
	// 指定是否开启断点续传模式，此处以true为例。默认为False，表示不开启。
	input.EnableCheckpoint = true
	// 指定分段大小，单位字节。此处以每段9M为例。
	input.PartSize = 9 * 1024 * 1024
	// 指定分段下载时的最大并发数，此处以并发数5为例
	input.TaskNum = 5
	output, err := c.Cli.DownloadFile(input)
	if err == nil {
		log.Infof("Download file(%s) under the bucket(%s) successful!\n", input.Key, input.Bucket)
		log.Infof("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		return nil
	}
	if obsError, ok := err.(obsSdk.ObsError); ok {
		//log.Errorf("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		log.Errorf("ObsError err:%+v", obsError)
	} else {
		// An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network
		log.Errorf("Exception err:%+v", err)
	}
	return err
}

// 判断文件存在
func (c *Client) IsExist(ctx context.Context, name string) (bool, error) {
	log.Infof("obs.IsExistGetObjectMetadata(%s) start ", name)
	// name 格式化成tmp/xxx.jpg
	if strings.Index(name, "/") == 0 {
		name = strings.TrimLeft(name, "/")
	}

	//var isExist = false
	input := &obsSdk.GetObjectMetadataInput{}
	// 指定存储桶名称
	input.Bucket = c.Bucket
	// 指定下载对象，此处以 example/objectname 为例。
	input.Key = name
	log.Infof("obs.IsExistGetObjectMetadata(%s) under the bucket(%s) successful!", input.Key, input.Bucket)
	output, err := c.Cli.GetObjectMetadata(input)
	if err == nil {
		log.Infof("Download file(%s) under the bucket(%s) successful!", input.Key, input.Bucket)
		log.Infof("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		return true, nil
	}
	log.Infof("Download file(%s) under the bucket(%s) successful!", input.Key, input.Bucket)
	if obsError, ok := err.(obsSdk.ObsError); ok {
		//log.Errorf("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		log.Errorf("ObsError err:%+v", obsError)
	} else {
		// An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network
		log.Errorf("Exception err:%+v", err)
	}
	return false, err
}
