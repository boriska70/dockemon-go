package collectors

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"strconv"
	"strings"
	"time"
)

// Host on which images are been monitored
var HostForImages hostImagesData

// Data type that keeps image related data for all images in the given host
type ImageBulkData struct {
	ImgData        []imageData
	DataType       string
	CollectionTime time.Time
	TotalLayers    int
	TotalImages    int
	TotalSizeMB    int64
}
type imageData struct {
	Id          string
	Created     string
	Labels      string
	ParentId    string
	RepoDigest  string
	Size        int64
	VirtualSize int64
	Host        hostImagesData
}
type hostImagesData struct {
	Host *HostStaticData
}

func (ibd *ImageBulkData) addImageData(id imageData) []imageData {
	ibd.ImgData = append(ibd.ImgData, id)
	return ibd.ImgData
}

// Endless loop to collect image related data
func ImageStats(client doClient, ch chan ImageBulkData) {

	log.Println("Collecting images data...")
	var imgBulk ImageBulkData
	imgBulk.DataType = "image_monitor"

	HostForImages.Host = getHostStaticData()

	for {
		info, _ := client.dc.Info(context.Background())
		options := types.ImageListOptions{All: false} //not including intermediate images

		imgBulk.CollectionTime = time.Now()
		imgBulk.TotalLayers = info.Images
		log.Info("Found layers: ", imgBulk.TotalLayers)

		images, err := client.dc.ImageList(context.Background(), options)
		if err != nil {
			panic(err)
		}
		var imgTotalSize int64 = 0
		var imgTotalImages int = 0
		for _, img := range images {
			imgTotalImages++
			imgTotalSize += img.Size
		}
		imgBulk.TotalImages = imgTotalImages
		imgBulk.TotalSizeMB = imgTotalSize / 1024 / 1024
		for _, img := range images {
			log.Debug("Image found: ", img.ID, img.Created, img.Labels, img.ParentID, img.RepoDigests, img.Size, img.VirtualSize)

			var image imageData
			image.Id = img.ID
			image.Created = time.Unix(img.Created, 0).UTC().Format(time.RFC3339)
			image.Labels = strings.Join(MapToArray(img.Labels), ",")
			image.ParentId = img.ParentID
			image.RepoDigest = strings.Join(img.RepoDigests, ",")
			image.Size = img.Size
			image.VirtualSize = img.VirtualSize
			image.Host = HostForImages
			imgBulk.ImgData = imgBulk.addImageData(image)
		}

		log.Info("Found Images: ", len(images), " with total size: "+strconv.FormatInt(imgTotalSize/1024/1024, 10)+" MB")

		if len(imgBulk.ImgData) > 0 {
			ch <- imgBulk
		}

		log.Debug("Image monitor is going to sleep under the next collection")
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}
