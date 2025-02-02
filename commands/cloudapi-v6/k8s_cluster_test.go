package commands

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	clusterTestPost = resources.K8sClusterForPost{
		KubernetesClusterForPost: ionoscloud.KubernetesClusterForPost{
			Properties: &ionoscloud.KubernetesClusterPropertiesForPost{
				Name:       &testClusterVar,
				K8sVersion: &testClusterVar,
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterVar},
			},
		},
	}
	clusterTestPut = resources.K8sClusterForPut{
		KubernetesClusterForPut: ionoscloud.KubernetesClusterForPut{
			Properties: &ionoscloud.KubernetesClusterPropertiesForPut{
				Name:       &testClusterVar,
				K8sVersion: &testClusterVar,
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterVar},
			},
		},
	}
	clusterNewTestPut = resources.K8sClusterForPut{
		KubernetesClusterForPut: ionoscloud.KubernetesClusterForPut{
			Properties: &ionoscloud.KubernetesClusterPropertiesForPut{
				Name:       &testClusterNewVar,
				K8sVersion: &testClusterNewVar,
				MaintenanceWindow: &ionoscloud.KubernetesMaintenanceWindow{
					DayOfTheWeek: &testClusterNewVar,
					Time:         &testClusterNewVar,
				},
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterNewVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterNewVar},
			},
		},
	}
	clusterTest = resources.K8sCluster{
		KubernetesCluster: ionoscloud.KubernetesCluster{
			Properties: &ionoscloud.KubernetesClusterProperties{
				Name:       &testClusterVar,
				K8sVersion: &testClusterVar,
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterVar},
			},
		},
	}
	clustersList = resources.K8sClusters{
		KubernetesClusters: ionoscloud.KubernetesClusters{
			Id: &testClusterVar,
			Items: &[]ionoscloud.KubernetesCluster{
				clusterTestId.KubernetesCluster,
				clusterTestId.KubernetesCluster,
			},
		},
	}
	clusterTestId = resources.K8sCluster{
		KubernetesCluster: ionoscloud.KubernetesCluster{
			Id: &testClusterVar,
			Properties: &ionoscloud.KubernetesClusterProperties{
				Name:       &testClusterVar,
				K8sVersion: &testClusterVar,
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterVar},
			},
		},
	}
	clusterTestGet = resources.K8sCluster{
		KubernetesCluster: ionoscloud.KubernetesCluster{
			Id: &testClusterVar,
			Properties: &ionoscloud.KubernetesClusterProperties{
				Name:                     &testClusterVar,
				K8sVersion:               &testClusterVar,
				AvailableUpgradeVersions: &testClusterSliceVar,
				ViableNodePoolVersions:   &testClusterSliceVar,
				MaintenanceWindow: &ionoscloud.KubernetesMaintenanceWindow{
					DayOfTheWeek: &testClusterVar,
					Time:         &testClusterVar,
				},
				S3Buckets: &[]ionoscloud.S3Bucket{
					{
						Name: &testClusterVar,
					},
				},
				ApiSubnetAllowList: &[]string{testClusterVar},
			},
			Metadata: &ionoscloud.DatacenterElementMetadata{
				State: &testStateVar,
			},
		},
	}
	clusters = resources.K8sClusters{
		KubernetesClusters: ionoscloud.KubernetesClusters{
			Id:    &testClusterVar,
			Items: &[]ionoscloud.KubernetesCluster{clusterTest.KubernetesCluster},
		},
	}
	clusterProperties = resources.K8sClusterProperties{
		KubernetesClusterProperties: ionoscloud.KubernetesClusterProperties{
			Name:       &testClusterNewVar,
			K8sVersion: &testClusterNewVar,
			MaintenanceWindow: &ionoscloud.KubernetesMaintenanceWindow{
				DayOfTheWeek: &testClusterNewVar,
				Time:         &testClusterNewVar,
			},
			S3Buckets: &[]ionoscloud.S3Bucket{
				{
					Name: &testClusterNewVar,
				},
			},
			ApiSubnetAllowList: &[]string{testClusterNewVar},
		},
	}
	clusterNew = resources.K8sCluster{
		KubernetesCluster: ionoscloud.KubernetesCluster{
			Properties: &clusterProperties.KubernetesClusterProperties,
		},
	}
	testClusterVar      = "test-cluster"
	testClusterSliceVar = []string{"test-cluster"}
	testClusterNewVar   = "test-new-cluster"
	testClusterErr      = errors.New("cluster test error")
)

func TestK8sCmd(t *testing.T) {
	var err error
	core.RootCmdTest.AddCommand(K8sCmd())
	if ok := K8sCmd().IsAvailableCommand(); !ok {
		err = errors.New("non-available cmd")
	}
	assert.NoError(t, err)
}

func TestPreRunK8sClusterId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		err := PreRunK8sClusterId(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunK8sClusterIdErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		err := PreRunK8sClusterId(cfg)
		assert.Error(t, err)
	})
}

func TestPreRunK8sClusterList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		err := PreRunK8sClusterList(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunK8sClusterListFilters(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{fmt.Sprintf("createdBy=%s", testQueryParamVar)})
		err := PreRunK8sClusterList(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunK8sClusterListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{fmt.Sprintf("%s=%s", testQueryParamVar, testQueryParamVar)})
		err := PreRunK8sClusterList(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(clusters, &testResponse, nil)
		err := RunK8sClusterList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterListQueryParams(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{fmt.Sprintf("%s=%s", testQueryParamVar, testQueryParamVar)})
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgOrderBy), testQueryParamVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgMaxResults), testMaxResultsVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(resources.K8sClusters{}, &testResponse, nil)
		err := RunK8sClusterList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(clusters, nil, testClusterErr)
		err := RunK8sClusterList(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, &testResponse, nil)
		err := RunK8sClusterGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterGetWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, testClusterErr)
		err := RunK8sClusterGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterGetWait(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		err := RunK8sClusterGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, testClusterErr)
		err := RunK8sClusterGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, &testResponse, nil)
		err := RunK8sClusterCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterCreateWaitIdErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, nil, nil)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreateWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestId, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, testClusterErr)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreateWaitState(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestId, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		err := RunK8sClusterCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterCreateWaitReqErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestId, &testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreateVersion(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetVersion().Return(testClusterVar, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, nil, nil)
		err := RunK8sClusterCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterCreateVersionErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetVersion().Return(testClusterVar, nil, testClusterErr)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreateResponseErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, &testResponse, testClusterErr)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterCreateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().CreateCluster(clusterTestPost, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, nil, testClusterErr)
		err := RunK8sClusterCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterUpdate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceTime), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceDay), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterNewVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().UpdateCluster(testClusterVar, clusterNewTestPut, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterNew, &testResponse, nil)
		err := RunK8sClusterUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterUpdateWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceTime), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceDay), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterNewVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().UpdateCluster(testClusterVar, clusterNewTestPut, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterNew, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, testClusterErr)
		err := RunK8sClusterUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterUpdateWaitState(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceTime), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceDay), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterNewVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().UpdateCluster(testClusterVar, clusterNewTestPut, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterNew, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		err := RunK8sClusterUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterUpdateOldUser(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().UpdateCluster(testClusterVar, clusterTestPut, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTest, nil, nil)
		err := RunK8sClusterUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterUpdateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceTime), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sMaintenanceDay), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterNewVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().UpdateCluster(testClusterVar, clusterNewTestPut, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterNew, nil, testClusterErr)
		err := RunK8sClusterUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterUpdateGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForState), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sVersion), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgS3Bucket), testClusterNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgApiSubnets), []string{testClusterNewVar})
		rm.CloudApiV6Mocks.K8s.EXPECT().GetCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&clusterTestGet, nil, testClusterErr)
		err := RunK8sClusterUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDelete(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, nil)
		err := RunK8sClusterDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterDeleteAll(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgServerUrl, constants.DefaultApiURL)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAll), true)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(clustersList, &testResponse, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, nil)
		err := RunK8sClusterDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterDeleteAllListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgServerUrl, constants.DefaultApiURL)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAll), true)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(clustersList, nil, testClusterErr)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteAllItemsErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgServerUrl, constants.DefaultApiURL)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAll), true)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(resources.K8sClusters{}, &testResponse, nil)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteAllLenErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgServerUrl, constants.DefaultApiURL)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAll), true)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(
			resources.K8sClusters{KubernetesClusters: ionoscloud.KubernetesClusters{Items: &[]ionoscloud.KubernetesCluster{}}}, &testResponse, nil)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteAllErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgVerbose, false)
		viper.Set(constants.ArgServerUrl, constants.DefaultApiURL)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAll), true)
		rm.CloudApiV6Mocks.K8s.EXPECT().ListClusters(gomock.AssignableToTypeOf(testListQueryParam)).Return(clustersList, &testResponse, nil)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, testClusterErr)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, nil)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteWaitReqErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), true)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(constants.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(nil, testClusterErr)
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunK8sClusterDeleteAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		rm.CloudApiV6Mocks.K8s.EXPECT().DeleteCluster(testClusterVar, gomock.AssignableToTypeOf(testQueryParamOther)).Return(nil, nil)
		err := RunK8sClusterDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunK8sClusterDeleteAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(constants.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, constants.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgK8sClusterId), testClusterVar)
		cfg.Stdin = os.Stdin
		err := RunK8sClusterDelete(cfg)
		assert.Error(t, err)
	})
}
