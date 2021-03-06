package handlers

import (
	"fmt"
	"path"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/optikon/api/api/v0/convert"
	"github.com/optikon/api/api/v0/mock"
	"github.com/optikon/api/api/v0/models"

	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/clusters"
)

func NewGetClusterByID() *getClusterById {
	return &getClusterById{}
}

type getClusterById struct{}

func (d *getClusterById) Handle(params clusters.GetClusterByIDParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	// Call cluster registry
	regCluster, err := config.ClusterClient.ClusterregistryV1alpha1().Clusters().Get(params.ClusterID, v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		return clusters.NewGetClusterByIDInternalServerError()
	}

	conv := convert.RegToOptikonCluster(*regCluster)

	// Use regular kube api client for this cluster, to get # pods
	pods, err := config.EdgeClients[conv.Metadata.Name].CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return clusters.NewGetClusterByIDInternalServerError()
	}
	conv.Metadata.Annotations["NumPods"] = strconv.Itoa(len(pods.Items))
	return clusters.NewGetClusterByIDOK().WithPayload(&conv)
}

func (d *getClusterById) MockHandle(params clusters.GetClusterByIDParams) middleware.Responder {
	payload := models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster{}
	statusCode, err := mock.GetMock(
		path.Join(config.MockBasePath, fmt.Sprintf("get-cluster-%s.json", params.ClusterID)), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewGetClusterByIDInternalServerError()
	}

	switch statusCode {
	case 201:
		return clusters.NewGetClusterByIDOK().WithPayload(&payload)
	case 400:
		return clusters.NewGetClusterByIDBadRequest()
	case 401:
		return clusters.NewGetClusterByIDUnauthorized()
	case 404:
		return clusters.NewGetClusterByIDNotFound()
	case 500:
		return clusters.NewGetClusterByIDInternalServerError()
	}
	return clusters.NewGetClusterByIDInternalServerError()
}
