package transactionenv

import (
	"fmt"

	"github.com/pachyderm/pachyderm/src/client/auth"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pps"
)

func unimplementedError(name string) error {
	return fmt.Errorf("%s is not implemented in this mock", name)
}

// MockPfsWrites provides a simple mock that can be used to instantiate a transaction env
type MockPfsWrites struct{}

func (mpw *MockPfsWrites) CreateRepo(*pfs.CreateRepoRequest) error {
	return unimplementedError("PfsWrites.CreateRepo")
}
func (mpw *MockPfsWrites) DeleteRepo(*pfs.DeleteRepoRequest) error {
	return unimplementedError("PfsWrites.DeleteRepo")
}
func (mpw *MockPfsWrites) StartCommit(*pfs.StartCommitRequest, *pfs.Commit) (*pfs.Commit, error) {
	return nil, unimplementedError("PfsWrites.StartCommit")
}
func (mpw *MockPfsWrites) FinishCommit(*pfs.FinishCommitRequest) error {
	return unimplementedError("PfsWrites.FinishCommit")
}
func (mpw *MockPfsWrites) DeleteCommit(*pfs.DeleteCommitRequest) error {
	return unimplementedError("PfsWrites.DeleteCommit")
}
func (mpw *MockPfsWrites) CreateBranch(*pfs.CreateBranchRequest) error {
	return unimplementedError("PfsWrites.CreateBranch")
}
func (mpw *MockPfsWrites) DeleteBranch(*pfs.DeleteBranchRequest) error {
	return unimplementedError("PfsWrites.DeleteBranch")
}

type MockPpsWrites struct{}

func (mpw *MockPpsWrites) UpdateJobState(*pps.UpdateJobStateRequest) error {
	return unimplementedError("PpsWrites.UpdateJobState")
}

type MockAuthWrites struct{}

func (maw *MockAuthWrites) SetScope(*auth.SetScopeRequest) (*auth.SetScopeResponse, error) {
	return unimplementedError("AuthWrites.SetScope")
}
func (maw *MockAuthWrites) SetACL(*auth.SetACLRequest) (*auth.SetACLResponse, error) {
	return unimplementedError("AuthWrites.SetACL")
}

type MockPfsPropagater struct{}

func (mpp *MockPfsPropagater) PropagateCommit(branch *pfs.Branch, isNewCommit bool) error {
	return unimplementedError("PfsPropagater.PropagateCommit")
}
func (mpp *MockPfsPropagater) Run() error {
	return unimplementedError("PfsPropagater.Run")
}

type MockAuthTransactionServer struct{}

func (mats *MockAuthTransactionServer) AuthorizeInTransaction(*TransactionContext, *auth.AuthorizeRequest) (*auth.AuthorizeResponse, error) {
	return nil, unimplementedError("AuthTransactionServer.AuthorizeInTransaction")
}
func (mats *MockAuthTransactionServer) GetScopeInTransaction(*TransactionContext, *auth.GetScopeRequest) (*auth.GetScopeResponse, error) {
	return nil, unimplementedError("AuthTransactionServer.GetScopeInTransaction")
}
func (mats *MockAuthTransactionServer) SetScopeInTransaction(*TransactionContext, *auth.SetScopeRequest) (*auth.SetScopeResponse, error) {
	return nil, unimplementedError("AuthTransactionServer.SetScopeInTransaction")
}
func (mats *MockAuthTransactionServer) GetACLInTransaction(*TransactionContext, *auth.GetACLRequest) (*auth.GetACLResponse, error) {
	return nil, unimplementedError("AuthTransactionServer.GetACLInTransaction")
}
func (mats *MockAuthTransactionServer) SetACLInTransaction(*TransactionContext, *auth.SetACLRequest) (*auth.SetACLResponse, error) {
	return nil, unimplementedError("AuthTransactionServer.SetACLInTransaction")
}

type MockPfsTransactionServer struct{}

func (mpts *MockPfsTransactionServer) NewPropagater(col.STM) PfsPropagater {
	return newMockPfsPropagater()
}
func (mpts *MockPfsTransactionServer) CreateRepoInTransaction(*TransactionContext, *pfs.CreateRepoRequest) error {
	return unimplementedError("PfsTransactionServer.CreateRepoInTransaction")
}
func (mpts *MockPfsTransactionServer) InspectRepoInTransaction(*TransactionContext, *pfs.InspectRepoRequest) (*pfs.RepoInfo, error) {
	return nil, unimplementedError("PfsTransactionServer.InspectRepoInTransaction")
}
func (mpts *MockPfsTransactionServer) DeleteRepoInTransaction(*TransactionContext, *pfs.DeleteRepoRequest) error {
	return unimplementedError("PfsTransactionServer.DeleteRepoInTransaction")
}
func (mpts *MockPfsTransactionServer) StartCommitInTransaction(*TransactionContext, *pfs.StartCommitRequest, *pfs.Commit) (*pfs.Commit, error) {
	return nil, unimplementedError("PfsTransactionServer.StartCommitInTransaction")
}
func (mpts *MockPfsTransactionServer) FinishCommitInTransaction(*TransactionContext, *pfs.FinishCommitRequest) error {
	return unimplementedError("PfsTransactionServer.FinishCommitInTransaction")
}
func (mpts *MockPfsTransactionServer) DeleteCommitInTransaction(*TransactionContext, *pfs.DeleteCommitRequest) error {
	return unimplementedError("PfsTransactionServer.DeleteCommitInTransaction")
}
func (mpts *MockPfsTransactionServer) CreateBranchInTransaction(*TransactionContext, *pfs.CreateBranchRequest) error {
	return unimplementedError("PfsTransactionServer.CreateBranchInTransaction")
}
func (mpts *MockPfsTransactionServer) DeleteBranchInTransaction(*TransactionContext, *pfs.DeleteBranchRequest) error {
	return unimplementedError("PfsTransactionServer.DeleteBranchInTransaction")
}

type MockPpsTransactionServer struct{}

func (mpts *MockPpsTransactionServer) UpdateJobStateInTransaction(*TransactionContext, *pps.UpdateJobStateRequest) error {
	return unimplementedError("PpsTransactionServer.UpdateJobStateInTransaction")
}
