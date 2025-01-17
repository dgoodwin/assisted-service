// Code generated by MockGen. DO NOT EDIT.
// Source: metricsManager.go

// Package metrics is a generated GoMock package.
package metrics

import (
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
	time "time"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ClusterRegistered mocks base method
func (m *MockAPI) ClusterRegistered(clusterVersion string, clusterID strfmt.UUID, emailDomain string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterRegistered", clusterVersion, clusterID, emailDomain)
}

// ClusterRegistered indicates an expected call of ClusterRegistered
func (mr *MockAPIMockRecorder) ClusterRegistered(clusterVersion, clusterID, emailDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRegistered", reflect.TypeOf((*MockAPI)(nil).ClusterRegistered), clusterVersion, clusterID, emailDomain)
}

// HostValidationFailed mocks base method
func (m *MockAPI) HostValidationFailed(clusterVersion, emailDomain string, hostValidationType models.HostValidationID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostValidationFailed", clusterVersion, emailDomain, hostValidationType)
}

// HostValidationFailed indicates an expected call of HostValidationFailed
func (mr *MockAPIMockRecorder) HostValidationFailed(clusterVersion, emailDomain, hostValidationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostValidationFailed", reflect.TypeOf((*MockAPI)(nil).HostValidationFailed), clusterVersion, emailDomain, hostValidationType)
}

// HostValidationChanged mocks base method
func (m *MockAPI) HostValidationChanged(clusterVersion, emailDomain string, hostValidationType models.HostValidationID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostValidationChanged", clusterVersion, emailDomain, hostValidationType)
}

// HostValidationChanged indicates an expected call of HostValidationChanged
func (mr *MockAPIMockRecorder) HostValidationChanged(clusterVersion, emailDomain, hostValidationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostValidationChanged", reflect.TypeOf((*MockAPI)(nil).HostValidationChanged), clusterVersion, emailDomain, hostValidationType)
}

// ClusterValidationFailed mocks base method
func (m *MockAPI) ClusterValidationFailed(clusterVersion, emailDomain string, clusterValidationType models.ClusterValidationID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterValidationFailed", clusterVersion, emailDomain, clusterValidationType)
}

// ClusterValidationFailed indicates an expected call of ClusterValidationFailed
func (mr *MockAPIMockRecorder) ClusterValidationFailed(clusterVersion, emailDomain, clusterValidationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterValidationFailed", reflect.TypeOf((*MockAPI)(nil).ClusterValidationFailed), clusterVersion, emailDomain, clusterValidationType)
}

// ClusterValidationChanged mocks base method
func (m *MockAPI) ClusterValidationChanged(clusterVersion, emailDomain string, clusterValidationType models.ClusterValidationID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterValidationChanged", clusterVersion, emailDomain, clusterValidationType)
}

// ClusterValidationChanged indicates an expected call of ClusterValidationChanged
func (mr *MockAPIMockRecorder) ClusterValidationChanged(clusterVersion, emailDomain, clusterValidationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterValidationChanged", reflect.TypeOf((*MockAPI)(nil).ClusterValidationChanged), clusterVersion, emailDomain, clusterValidationType)
}

// InstallationStarted mocks base method
func (m *MockAPI) InstallationStarted(clusterVersion string, clusterID strfmt.UUID, emailDomain, userManagedNetworking string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InstallationStarted", clusterVersion, clusterID, emailDomain, userManagedNetworking)
}

// InstallationStarted indicates an expected call of InstallationStarted
func (mr *MockAPIMockRecorder) InstallationStarted(clusterVersion, clusterID, emailDomain, userManagedNetworking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallationStarted", reflect.TypeOf((*MockAPI)(nil).InstallationStarted), clusterVersion, clusterID, emailDomain, userManagedNetworking)
}

// ClusterHostInstallationCount mocks base method
func (m *MockAPI) ClusterHostInstallationCount(clusterID strfmt.UUID, emailDomain string, hostCount int, clusterVersion string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterHostInstallationCount", clusterID, emailDomain, hostCount, clusterVersion)
}

// ClusterHostInstallationCount indicates an expected call of ClusterHostInstallationCount
func (mr *MockAPIMockRecorder) ClusterHostInstallationCount(clusterID, emailDomain, hostCount, clusterVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterHostInstallationCount", reflect.TypeOf((*MockAPI)(nil).ClusterHostInstallationCount), clusterID, emailDomain, hostCount, clusterVersion)
}

// ClusterHostsNTPFailures mocks base method
func (m *MockAPI) ClusterHostsNTPFailures(clusterID strfmt.UUID, emailDomain string, hostNTPFailureCount int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterHostsNTPFailures", clusterID, emailDomain, hostNTPFailureCount)
}

// ClusterHostsNTPFailures indicates an expected call of ClusterHostsNTPFailures
func (mr *MockAPIMockRecorder) ClusterHostsNTPFailures(clusterID, emailDomain, hostNTPFailureCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterHostsNTPFailures", reflect.TypeOf((*MockAPI)(nil).ClusterHostsNTPFailures), clusterID, emailDomain, hostNTPFailureCount)
}

// Duration mocks base method
func (m *MockAPI) Duration(operation string, duration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Duration", operation, duration)
}

// Duration indicates an expected call of Duration
func (mr *MockAPIMockRecorder) Duration(operation, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Duration", reflect.TypeOf((*MockAPI)(nil).Duration), operation, duration)
}

// ClusterInstallationFinished mocks base method
func (m *MockAPI) ClusterInstallationFinished(log logrus.FieldLogger, result, clusterVersion string, clusterID strfmt.UUID, emailDomain string, installationStartedTime strfmt.DateTime) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterInstallationFinished", log, result, clusterVersion, clusterID, emailDomain, installationStartedTime)
}

// ClusterInstallationFinished indicates an expected call of ClusterInstallationFinished
func (mr *MockAPIMockRecorder) ClusterInstallationFinished(log, result, clusterVersion, clusterID, emailDomain, installationStartedTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInstallationFinished", reflect.TypeOf((*MockAPI)(nil).ClusterInstallationFinished), log, result, clusterVersion, clusterID, emailDomain, installationStartedTime)
}

// ReportHostInstallationMetrics mocks base method
func (m *MockAPI) ReportHostInstallationMetrics(log logrus.FieldLogger, clusterVersion string, clusterID strfmt.UUID, emailDomain string, boot *models.Disk, h *models.Host, previousProgress *models.HostProgressInfo, currentStage models.HostStage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportHostInstallationMetrics", log, clusterVersion, clusterID, emailDomain, boot, h, previousProgress, currentStage)
}

// ReportHostInstallationMetrics indicates an expected call of ReportHostInstallationMetrics
func (mr *MockAPIMockRecorder) ReportHostInstallationMetrics(log, clusterVersion, clusterID, emailDomain, boot, h, previousProgress, currentStage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportHostInstallationMetrics", reflect.TypeOf((*MockAPI)(nil).ReportHostInstallationMetrics), log, clusterVersion, clusterID, emailDomain, boot, h, previousProgress, currentStage)
}

// DiskSyncDuration mocks base method
func (m *MockAPI) DiskSyncDuration(clusterID, hostID strfmt.UUID, diskPath string, syncDuration int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DiskSyncDuration", clusterID, hostID, diskPath, syncDuration)
}

// DiskSyncDuration indicates an expected call of DiskSyncDuration
func (mr *MockAPIMockRecorder) DiskSyncDuration(clusterID, hostID, diskPath, syncDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskSyncDuration", reflect.TypeOf((*MockAPI)(nil).DiskSyncDuration), clusterID, hostID, diskPath, syncDuration)
}

// ImagePullStatus mocks base method
func (m *MockAPI) ImagePullStatus(clusterID, hostID strfmt.UUID, imageName, resultStatus string, downloadRate float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImagePullStatus", clusterID, hostID, imageName, resultStatus, downloadRate)
}

// ImagePullStatus indicates an expected call of ImagePullStatus
func (mr *MockAPIMockRecorder) ImagePullStatus(clusterID, hostID, imageName, resultStatus, downloadRate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePullStatus", reflect.TypeOf((*MockAPI)(nil).ImagePullStatus), clusterID, hostID, imageName, resultStatus, downloadRate)
}

// FileSystemUsage mocks base method
func (m *MockAPI) FileSystemUsage(usageInPercentage float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FileSystemUsage", usageInPercentage)
}

// FileSystemUsage indicates an expected call of FileSystemUsage
func (mr *MockAPIMockRecorder) FileSystemUsage(usageInPercentage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileSystemUsage", reflect.TypeOf((*MockAPI)(nil).FileSystemUsage), usageInPercentage)
}
