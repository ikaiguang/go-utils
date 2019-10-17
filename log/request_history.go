package golog

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

// SaveHistory save history
func SaveHistory(historyInterface RequestHistoryInterface) error {
	if err := historyInterface.Save(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// RequestHistoryInterface history
type RequestHistoryInterface interface {
	Save() error
}

// NewRequestHistory history
func NewRequestHistory(db *gorm.DB, history *RequestHistory) *requestHistory {
	return &requestHistory{db: db, history: history}
}

// requestHistory history
type requestHistory struct {
	db      *gorm.DB        // db
	history *RequestHistory // history
}

// Save save
func (r *requestHistory) Save() error {
	// create table
	var createErr error
	requestHistoryTableInit.Do(func() {
		if initErr := r.history.CreateTable(r.db); initErr != nil {
			createErr = errors.WithStack(initErr)
		}
	})
	if createErr != nil {
		return createErr
	}

	// insert one
	if err := r.history.InsertOne(r.db); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// table
var (
	// mutex
	requestHistoryTableInit sync.Once

	// table
	RequestHistoryTableName    = "request_history"
	RequestHistoryTableNameSep = "{{$table}}"
	RequestHistoryTableSQL     = `
CREATE TABLE IF NOT EXISTS {{$table}}
(
    id               BIGINT AUTO_INCREMENT COMMENT 'request_history',
    user_id          BIGINT        DEFAULT '0' NOT NULL COMMENT 'user_id',
    user_type        VARCHAR(32)   DEFAULT ''  NOT NULL COMMENT 'user_type',
    user_ip          VARCHAR(150)  DEFAULT ''  NOT NULL COMMENT 'ipv4:port/ipv6:port',
    request_time     BIGINT        DEFAULT '0' NOT NULL COMMENT 'request_time',
    request_method   VARCHAR(32)   DEFAULT ''  NOT NULL COMMENT 'request_method',
    request_duration BIGINT        DEFAULT '0' NOT NULL COMMENT 'request_duration',
    request_latency  VARCHAR(255)  DEFAULT ''  NOT NULL COMMENT 'request_latency',
    request_uri      VARCHAR(255)  DEFAULT ''  NOT NULL COMMENT 'request_uri',
    request_url      VARCHAR(2000) DEFAULT ''  NOT NULL COMMENT 'request_url',
    response_code    INT           DEFAULT '0' NOT NULL COMMENT 'response_code',
    app_version      VARCHAR(255)  DEFAULT ''  NOT NULL COMMENT 'app_version',
    content_type     VARCHAR(255)  DEFAULT ''  NOT NULL COMMENT 'content_type',
    user_agent       VARCHAR(1000) DEFAULT ''  NOT NULL COMMENT 'user_agent',
    response_message TEXT                      NULL     COMMENT 'response_message',
    request_body     TEXT                      NULL     COMMENT 'request_body',
    PRIMARY KEY (id),
    KEY user_id (user_id),
    KEY request_time (request_time)
)
    COMMENT 'request_history'
    DEFAULT CHARSET utf8mb4;
`
)

// RequestHistory request history
type RequestHistory struct {
	Id              interface{} `json:"id"`
	UserId          interface{} `json:"user_id"`
	UserType        string      `json:"user_type"`
	UserIp          string      `json:"user_ip"`
	RequestTime     interface{} `json:"request_time"`
	RequestMethod   string      `json:"request_method"`
	RequestDuration interface{} `json:"request_duration"`
	RequestLatency  string      `json:"request_latency"`
	RequestUri      string      `json:"request_uri"`
	RequestUrl      string      `json:"request_url"`
	ResponseCode    interface{} `json:"response_code"`
	AppVersion      string      `json:"app_version"`
	ContentType     string      `json:"content_type"`
	UserAgent       string      `json:"user_agent"`
	ResponseMessage string      `json:"response_message"`
	RequestBody     interface{} `json:"request_body"`
}

// TableName table name
func (r *RequestHistory) TableName() string {
	return RequestHistoryTableName
}

// CreateTable
func (r *RequestHistory) CreateTable(dbConn *gorm.DB) error {
	tableSQL := strings.Replace(RequestHistoryTableSQL, RequestHistoryTableNameSep, RequestHistoryTableName, 1)
	if err := dbConn.Exec(tableSQL).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertOne insert one
func (r *RequestHistory) InsertOne(dbConn *gorm.DB) error {
	if err := dbConn.Create(r).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
