package collectors

import "fmt"

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/distribution/context"
	log "github.com/Sirupsen/logrus"
	"time"
	"strconv"
)

var HostForImages hostImagesData


type ImageBulkData struct {
	ImgData [] imageData
	DataType string
	CollectionTime time.Time
}
type imageData struct {
	Id     string
	Created  time.Time
	Labels string
	ParentId string
	RepoDigest string
	Size int64
	VirtualSize int64
	Host hostImagesData
}
type hostImagesData struct {
	Host                HostStaticData
	TotalLayers   int
	TotalImages   int
	TotalSizeMB   int64
}


func (ibd *ImageBulkData) addImageData(id imageData) [] imageData {
	ibd.ImgData = append(ibd.ImgData, id)
	return ibd.ImgData
}

func ImageStats(client doClient) {

	HostForImages.Host = getHostStaticData()

	log.Println("Collecting images data...")
	log.Println(client.dc.Info(context.Background()))
	info, _ := client.dc.Info(context.Background())

	var contBulk ContainersBulkData
	contBulk.DataType="image_monitor"

	options := types.ImageListOptions{All:false}	//not including intermediate images
	for {
		contBulk.CollectionTime=time.Now()
		HostForImages.TotalLayers = info.Images
		log.Info("Found layers: ",HostForImages.TotalLayers)

		images, err := client.dc.ImageList(context.Background(), options)
		if err != nil {
			panic(err)
		}
		var imgTotalSize int64
		for _, img := range images {
			log.Debug("Image found: ",img.ID, img.Created, img.Labels, img.ParentID, img.RepoDigests, img.Size, img.VirtualSize)
			imgTotalSize += img.Size
		}

		log.Info("Found Images: ",len(images), " with total size: " + strconv.FormatInt(imgTotalSize/1024/1024,10) + " MB")


		fmt.Println("Going to sleep under the next collection")
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}

func mapToArray2(m map[string]string) [] string  {

	res := make([]string, 0)

	for ind,val := range m {
		res = append(res,ind+"="+val)
	}

	return res
}


