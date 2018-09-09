package registry

import (
	"github.com/itouri/sgx-iaas/pkg/domain/glance"
	uuid "github.com/satori/go.uuid"
)

// TODO 意味あるかな この辺
var registeredImages []glance.Image

func init() {
	registeredImages = []glance.Image{}
}

func RegisterImage(image *glance.Image) {
	registeredImages = append(registeredImages, *image)
}

func DeleteImage(id uuid.UUID) {
	delIndex := -1
	for i, r := range registeredImages {
		if r.ID == id {
			delIndex = i
			break
		}
	}
	if delIndex == -1 {
		return
	}
	registeredImages = append(registeredImages[:delIndex], registeredImages[delIndex+1:]...)
}

func GetRegisteredImages() []glance.Image {
	return registeredImages
}
