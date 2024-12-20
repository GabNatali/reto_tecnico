package email

type TookDetail struct {
	Total            int `json:"total"`
	IdxTook          int `json:"idx_took"`
	WaitQueue        int `json:"wait_queue"`
	ClusterTotal     int `json:"cluster_total"`
	ClusterWaitQueue int `json:"cluster_wait_queue"`
}

type Hit struct {
	Timestamp int64  `json:"_timestamp"`
	Body      string `json:"body"`
	Date      string `json:"date"`
	From      string `json:"from"`
	MessageID string `json:"message_id"`
	Subject   string `json:"subject"`
	To        string `json:"to"`
	Count     int    `json:"count,omitempty"`
}

type OpenObserverResponse struct {
	Took             int        `json:"took"`
	TookDetail       TookDetail `json:"took_detail"`
	Hits             []Hit      `json:"hits"`
	Total            int        `json:"total"`
	From             int        `json:"from"`
	Size             int        `json:"size"`
	CachedRatio      int        `json:"cached_ratio"`
	ScanSize         int        `json:"scan_size"`
	IdxScanSize      int        `json:"idx_scan_size"`
	ScanRecords      int        `json:"scan_records"`
	TraceID          string     `json:"trace_id"`
	IsPartial        bool       `json:"is_partial"`
	ResultCacheRatio int        `json:"result_cache_ratio"`
}
