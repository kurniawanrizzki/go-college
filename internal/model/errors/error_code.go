package errors

import "github.com/palantir/stacktrace"

type (
	ServiceType  int
	Code         = stacktrace.ErrorCode
	ErrorMessage map[Code]Message

	Message struct {
		EN            string
		ID            string
		StatusCode    int
		HasAnnotation bool
	}
)

const (
	COMMON ServiceType = 1
)

const (
	CodeHTTPBadRequest = Code(iota + 100)
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPMarshal
	CodeHTTPConflict
	CodeHTTPForbidden
	CodeHTTPUnprocessableEntity
	CodeHTTPTooManyRequest
	CodeHTTPValidatorError
	CodeHTTPServiceUnavailable
	CodeHTTPParamDecode
	CodeHTTPErrorOnReadBody
	CodeHTTPExternalAPI
)

const (
	CodeSQLBuilder = Code(iota + 200)
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLCreate
	CodeSQLUpdate
	CodeSQLDelete
	CodeSQLUnlink
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLPrepareStmt
	CodeSQLRecordMustExist
	CodeSQLCannotRetrieveLastInsertID
	CodeSQLCannotRetrieveAffectedRows
	CodeSQLUniqueConstraint
	CodeSQLRecordDoesNotMatch
	CodeSQLRecordIsExpired
	CodeSQLRecordDoesNotExist
	CodeSQLForeignKeyMissing
	CodeSQLTxRollback
	CodeRequestIDIsNotMatch
	CodeSQLConflict
	CodeSQLEmptyRow
	CodeSQLTableNotExist
	CodeSQLQueryBuild
)

const (
	CodeTokenStillValid = Code(iota + 300)
	CodeTokenRefreshStillValid
)

const (
	CodeCacheMarshal = Code(iota + 400)
	CodeCacheUnmarshal
	CodeCacheGetSimpleKey
	CodeCacheSetSimpleKey
	CodeCacheDeleteSimpleKey
	CodeCacheGetHashKey
	CodeCacheSetHashKey
	CodeCacheDeleteHashKey
	CodeCacheSetExpiration
	CodeCacheDecode
	CodeCacheLockNotAcquired
	CodeCacheLockFailed
	CodeCacheInvalidCastType
	CodeCacheNotFound
)

const (
	CodeFileNotFound = Code(iota + 500)
	CodeFileRead
	CodeDuplicateQuery
	CodeTemplateParse
	CodeTemplateExecute
	CodeSQLQueryNotFound
	CodeInvalidIdentifier
)

var (
	svcError map[ServiceType]ErrorMessage

	ErrCode      = stacktrace.GetCode
	New          = stacktrace.NewError
	NewWithCode  = stacktrace.NewErrorWithCode
	RootCause    = stacktrace.RootCause
	Wrap         = stacktrace.Propagate
	WrapWithCode = stacktrace.PropagateWithCode
)
