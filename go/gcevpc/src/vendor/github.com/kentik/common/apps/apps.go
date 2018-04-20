package apps

// Mapping for application types to column names for krpobe.

const (
	KPROBE_DEVICE_TYPE  = "kprobe"
	KPROBE_APP_PROTOCOL = "APP_PROTOCOL"

	DNS  = 1
	DHCP = 4
	TLS  = 3
	HTTP = 2

	STR00 = "STR00"
	STR01 = "STR01"
	STR02 = "STR02"
	STR03 = "STR03"
	STR04 = "STR04"
	STR05 = "STR05"
	INT00 = "INT00"
	INT01 = "INT01"
	INT02 = "INT02"
	INT03 = "INT03"
	INT04 = "INT04"
	INT05 = "INT05"

	KFLOW_DNS_QUERY      = "KFLOW_DNS_QUERY"
	KFLOW_DNS_QUERY_TYPE = "KFLOW_DNS_QUERY_TYPE"
	KFLOW_DNS_RET_CODE   = "KFLOW_DNS_RET_CODE"
	KFLOW_DNS_RESPONSE   = "KFLOW_DNS_RESPONSE"
	KFLOW_HTTP_URL       = "KFLOW_HTTP_URL"
	KFLOW_HTTP_HOST      = "KFLOW_HTTP_HOST"
	KFLOW_HTTP_REFERER   = "KFLOW_HTTP_REFERER"
	KFLOW_HTTP_UA        = "KFLOW_HTTP_UA"
	KFLOW_HTTP_STATUS    = "KFLOW_HTTP_STATUS"
	TLS_SERVER_NAME      = "TLS_SERVER_NAME"
	TLS_SERVER_VERSION   = "TLS_SERVER_VERSION"
	TLS_CIPHER_SUITE     = "TLS_CIPHER_SUITE"
	DHCP_OP              = "DHCP_OP"
	DHCP_MSG_TYPE        = "DHCP_MSG_TYPE"
	DHCP_CI_ADDR         = "DHCP_CI_ADDR"
	DHCP_YI_ADDR         = "DHCP_YI_ADDR"
	DHCP_SI_ADDR         = "DHCP_SI_ADDR"
	DHCP_LEASE           = "DHCP_LEASE"
	DHCP_CH_ADDR         = "DHCP_CH_ADDR"
	DHCP_HOSTNAME        = "DHCP_HOSTNAME"
	DHCP_DOMAIN          = "DHCP_DOMAIN"
)

var (
	COLUMN_TO_APP = map[string]uint32{
		KFLOW_DNS_QUERY:      DNS,
		KFLOW_DNS_QUERY_TYPE: DNS,
		KFLOW_DNS_RET_CODE:   DNS,
		KFLOW_DNS_RESPONSE:   DNS,
		KFLOW_HTTP_URL:       HTTP,
		KFLOW_HTTP_HOST:      HTTP,
		KFLOW_HTTP_REFERER:   HTTP,
		KFLOW_HTTP_UA:        HTTP,
		KFLOW_HTTP_STATUS:    HTTP,
		TLS_SERVER_NAME:      TLS,
		TLS_SERVER_VERSION:   TLS,
		TLS_CIPHER_SUITE:     TLS,
		DHCP_OP:              DHCP,
		DHCP_MSG_TYPE:        DHCP,
		DHCP_CI_ADDR:         DHCP,
		DHCP_YI_ADDR:         DHCP,
		DHCP_SI_ADDR:         DHCP,
		DHCP_LEASE:           DHCP,
		DHCP_CH_ADDR:         DHCP,
		DHCP_HOSTNAME:        DHCP,
		DHCP_DOMAIN:          DHCP,
	}

	APP_TO_COLUMN = map[uint32]map[string]string{
		DNS: map[string]string{
			KFLOW_DNS_QUERY:      STR00,
			KFLOW_DNS_QUERY_TYPE: INT00,
			KFLOW_DNS_RET_CODE:   INT01,
			KFLOW_DNS_RESPONSE:   STR01,
		},
		HTTP: map[string]string{
			KFLOW_HTTP_URL:     STR00,
			KFLOW_HTTP_HOST:    STR01,
			KFLOW_HTTP_REFERER: STR02,
			KFLOW_HTTP_UA:      STR03,
			KFLOW_HTTP_STATUS:  INT00,
		},
		TLS: map[string]string{
			TLS_SERVER_NAME:    STR00,
			TLS_SERVER_VERSION: INT00,
			TLS_CIPHER_SUITE:   INT01,
		},
		DHCP: map[string]string{
			DHCP_OP:       INT00,
			DHCP_MSG_TYPE: INT01,
			DHCP_CI_ADDR:  INT02,
			DHCP_YI_ADDR:  INT03,
			DHCP_SI_ADDR:  INT04,
			DHCP_LEASE:    INT05,
			DHCP_CH_ADDR:  STR00,
			DHCP_HOSTNAME: STR01,
			DHCP_DOMAIN:   STR02,
		},
	}
)
