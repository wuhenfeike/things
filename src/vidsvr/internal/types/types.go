// Code generated by goctl. DO NOT EDIT.
package types

type HooksApiResp struct {
	Code int64  `json:"code"` //状态码回复:0
	Msg  string `json:"msg"`  //msg:success
}

type HooksApiFlowReportReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Duration      int64  `json:"duration"`
	Params        string `json:"params"`
	Player        bool   `json:"player"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	TotalBytes    int64  `json:"totalBytes"`
	Vhost         string `json:"vhost"`
	Ip            string `json:"ip"`
	Port          int64  `json:"port"`
	Id            string `json:"id"`
}

type HooksApiHttpAccessReq struct {
	MediaServerId                 string `json:"mediaServerId"`
	HeaderAccept                  string `json:"header.Accept"`
	HeaderAcceptEncod             string `json:"header.Accept-Encoding"`
	HeaderAcceptLanguage          string `json:"header.Accept-Language"`
	HeaderCacheControl            string `json:"header.Cache-Control"`
	HeaderConnection              string `json:"header.Connection"`
	HeaderHost                    string `json:"header.Host"`
	HeaderUpgradeInsecureRequests string `json:"header.Upgrade-Insecure-Requests"`
	HeaderUserAgent               string `json:"header.User-Agent"`
	ID                            string `json:"id"`
	IP                            string `json:"ip"`
	IsDir                         bool   `json:"is_dir"`
	Params                        string `json:"params"`
	Path                          string `json:"path"`
	Port                          int64  `json:"port"`
}

type HooksApiHttpAccessResp struct {
	Code   int64  `json:"code"`
	Err    string `json:"err"`
	Path   string `json:"path"`
	Second int64  `json:"second"`
}

type HooksApiPlayReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Params        string `json:"params"`
	Port          int64  `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type HooksApiPublishReq struct {
	HooksApiPlayReq
}

type HooksApiPublishResp struct {
	Code           int64  `json:"code"`
	AddMuteAudio   bool   `json:"add_mute_audio"`
	ContinuePushMs int64  `json:"continue_push_ms"`
	EnableAudio    bool   `json:"enable_audio"`
	EnableFmp4     bool   `json:"enable_fmp4"`
	EnableHls      bool   `json:"enable_hls"`
	EnableHlsFmp4  bool   `json:"enable_hls_fmp4"`
	EnableMp4      bool   `json:"enable_mp4"`
	EnableRtmp     bool   `json:"enable_rtmp"`
	EnableRtsp     bool   `json:"enable_rtsp"`
	EnableTs       bool   `json:"enable_ts"`
	HlsSavePath    string `json:"hls_save_path"`
	ModifyStamp    bool   `json:"modify_stamp"`
	Mp4AsPlayer    bool   `json:"mp4_as_player"`
	Mp4MaxSecond   int64  `json:"mp4_max_second"`
	Mp4SavePath    string `json:"mp4_save_path"`
	AutoClose      bool   `json:"auto_close"`
	StreamReplace  string `json:"stream_replace"`
}

type HooksApiRecordMp4Req struct {
	MediaServerId string  `json:"mediaServerId"`
	App           string  `json:"app"`
	FileName      string  `json:"file_name"`
	FilePath      string  `json:"file_path"`
	FileSize      int64   `json:"file_size"`
	Folder        string  `json:"folder"`
	StartTime     int64   `json:"start_time"`
	Stream        string  `json:"stream"`
	TimeLen       float32 `json:"time_len"`
	Url           string  `json:"url"`
	Vhost         string  `json:"vhost"`
}

type HooksApiRtspAuthReq struct {
	MustNoEncrypt bool `json:"must_no_encrypt"`
	HooksApiPlayReq
	Realm    string `json:"realm"`
	UserName string `json:"user_name"`
}

type HooksApiRtspAuthResp struct {
	Code      int64  `json:"code"`
	Encrypted bool   `json:"encrypted"`
	Passwd    string `json:"passwd"`
}

type HooksApiRtspRealmReq struct {
	HooksApiPlayReq
}

type HooksApiRtspRealmResp struct {
	Code  int64  `json:"code"`
	Realm string `json:"realm"`
}

type HooksApiShellLoginReq struct {
	MediaServerId string `json:"mediaServerId"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Passwd        string `json:"passwd"`
	Port          int64  `json:"port"`
	User_name     string `json:"user_name"`
}

type HooksApiStreamChangedReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Regist        bool   `json:"regist"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type HooksApiStreamChangedRep struct {
	MediaServerId    string        `json:"mediaServerId"`
	App              string        `json:"app"`
	Regist           bool          `json:"regist"`
	Stream           string        `json:"stream"`
	Vhost            string        `json:"vhost"`
	AliveSecond      int64         `json:"aliveSecond"`
	BytesSpeed       int64         `json:"bytesSpeed"`
	CreateStamp      int64         `json:"createStamp"`
	OriginSock       OriginSock    `json:"originSock"`
	OriginType       int64         `json:"originType"`
	OriginTypeStr    string        `json:"originTypeStr"`
	OriginUrl        string        `json:"originUrl"`
	ReaderCount      int64         `json:"readerCount"`
	Schema           string        `json:"schema"`
	TotalReaderCount int64         `json:"totalReaderCount"`
	Tracks           []StreamTrack `json:"tracks"`
	IsRecordingMp4   bool          `json:"isRecordingMp4"`
	IsRecordingHLS   bool          `json:"isRecordingHLS"`
	IsShareChannel   bool          `json:"isShareChannel"`
	IsAutoPush       bool          `json:"isAutoPush"`
	IsAutoRecord     bool          `json:"isAutoRecord"`
	IsPTZ            bool          `json:"isPTZ"`
}

type HooksApiStreamNoneReaderReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type HooksApiStreamNoneReaderResp struct {
	Close bool `json:"close"`
	Code  bool `json:"code"`
}

type HooksApiStreamNotFoundReq struct {
	HooksApiPlayReq
}

type HooksApiServerStartedReq struct {
	ServerConfig
}

type HooksApiServerKeepaliveReq struct {
	Data          Statistic `json:"data"`
	MediaServerId string    `json:"mediaServerId"`
}

type HooksApiRtpServerTimeoutReq struct {
	LocalPort     int    `json:"local_port"`
	ReUsePort     bool   `json:"re_use_port"`
	Ssrc          int    `json:"ssrc"`
	StreamId      string `json:"stream_id"`
	TcpMode       int    `json:"tcp_mode"`
	MediaServerId string `json:"mediaServerId"`
}

type SvcZlmedia struct {
	IP     string `json:"ip"`
	Port   int64  `json:"port"`
	Secret string `json:"Secret"`
}

type ThreadLoad struct {
	Delay int64 `json:"delay"`
	Load  int64 `json:"load"`
}

type MediaPlayer struct {
	OriginSock
	Typeid string `json:"typeid"`
}

type OriginSock struct {
	Identifier string `json:"identifier"`
	LocalIp    string `json:"local_ip"`
	LocalPort  int64  `json:"local_port""`
	PeerIp     string `json:"peer_ip"`
	PeerPort   int64  `json:"peer_port"`
}

type StreamTrack struct {
	Channels    int64  `json:"channels"`
	CodecId     int64  `json:"codec_id"`
	CodecIdName string `json:"codec_id_name"`
	CodecType   int64  `json:"codec_type"`
	Ready       bool   `json:"ready"`
	SampleBit   int64  `json:"sample_bit"`
	SampleRate  int64  `json:"sample_rate"`
	Fps         int64  `json:"fps"`
	Height      int64  `json:"height"`
	Width       int64  `json:"width"`
}

type Statistic struct {
	Buffer                int `json:"Buffer"`
	BufferLikeString      int `json:"BufferLikeString"`
	BufferList            int `json:"BufferList"`
	BufferRaw             int `json:"BufferRaw"`
	Frame                 int `json:"Frame"`
	FrameImp              int `json:"FrameImp"`
	MediaSource           int `json:"MediaSource"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer"`
	RtmpPacket            int `json:"RtmpPacket"`
	RtpPacket             int `json:"RtpPacket"`
	Socket                int `json:"Socket"`
	TcpClient             int `json:"TcpClient"`
	TcpServer             int `json:"TcpServer"`
	TcpSession            int `json:"TcpSession"`
	UdpServer             int `json:"UdpServer"`
	UdpSession            int `json:"UdpSession"`
}

type AllSession struct {
	Id        string `json:"id"`
	LocalIp   string `json:"local_ip"`
	LocalPort int64  `json:"local_port""`
	PeerIp    string `json:"peer_ip"`
	PeerPort  int64  `json:"peer_port"`
	Typeid    string `json:"typeid"`
}

type MediaList struct {
	App              string        `json:"app"`
	ReaderCount      int64         `json:"readerCount"`
	TotalReaderCount int64         `json:"totalReaderCount"`
	Schema           string        `json:"schema"`
	Stream           string        `json:"stream"`
	OriginSock       OriginSock    `json:"originSock"`
	OriginType       int64         `json:"originType"`
	OriginTypeStr    string        `json:"originTypeStr"`
	OriginUrl        string        `json:"originUrl"`
	CreateStamp      string        `json:"createStamp"`
	AliveSecond      string        `json:"aliveSecond"`
	BytesSpeed       int64         `json:"bytesSpeed"`
	Tracks           []StreamTrack `json:"tracks"`
	Vhost            string        `json:"vhost"`
}

type Version struct {
	BranchName string `json:"branchName"`
	BuildTime  string `json:"buildTime"`
	CommitHash string `json:"commitHash"`
}

type ServerConfig struct {
	ApiDebug                       string `json:"api.apiDebug,omitempty"`
	ApiDefaultSnap                 string `json:"api.defaultSnap,omitempty"`
	ApiSecret                      string `json:"api.secret,omitempty"`
	ApiSnapRoot                    string `json:"api.snapRoot,omitempty"`
	ClusterOriginUrl               string `json:"cluster.origin_url,omitempty"`
	ClusterRetryCount              string `json:"cluster.retry_count,omitempty"`
	ClusterTimeoutSec              string `json:"cluster.timeout_sec,omitempty"`
	FfmpegBin                      string `json:"ffmpeg.bin,omitempty"`
	FfmpegCmd                      string `json:"ffmpeg.cmd,omitempty"`
	FfmpegLog                      string `json:"ffmpeg.log,omitempty"`
	FfmpegRestartSec               string `json:"ffmpeg.restart_sec,omitempty"`
	FfmpegSnap                     string `json:"ffmpeg.snap,omitempty"`
	GeneralCheckNvidiaDev          string `json:"general.check_nvidia_dev,omitempty"`
	GeneralEnableVhost             string `json:"general.enableVhost,omitempty"`
	GeneralEnableFfmpegLog         string `json:"general.enable_ffmpeg_log,omitempty"`
	GeneralFlowThreshold           string `json:"general.flowThreshold,omitempty"`
	GeneralMaxStreamWaitMS         string `json:"general.maxStreamWaitMS,omitempty"`
	GeneralMediaServerId           string `json:"general.mediaServerId,omitempty"`
	GeneralMergeWriteMS            string `json:"general.mergeWriteMS,omitempty"`
	GeneralResetWhenRePlay         string `json:"general.resetWhenRePlay,omitempty"`
	GeneralStreamNoneReaderDelayMS string `json:"general.streamNoneReaderDelayMS,omitempty"`
	GeneralUnreadyFrameCache       string `json:"general.unready_frame_cache,omitempty"`
	GeneralWaitAddTrackMs          string `json:"general.wait_add_track_ms,omitempty"`
	GeneralWaitTrackReadyMs        string `json:"general.wait_track_ready_ms,omitempty"`
	HlsBroadcastRecordTs           string `json:"hls.broadcastRecordTs,omitempty"`
	HlsDeleteDelaySec              string `json:"hls.deleteDelaySec,omitempty"`
	HlsFileBufSize                 string `json:"hls.fileBufSize,omitempty"`
	HlsSegDur                      string `json:"hls.segDur,omitempty"`
	HlsSegKeep                     string `json:"hls.segKeep,omitempty"`
	HlsSegNum                      string `json:"hls.segNum,omitempty"`
	HlsSegRetain                   string `json:"hls.segRetain,omitempty"`
	HookAliveInterval              string `json:"hook.alive_interval,omitempty"`
	HookEnable                     string `json:"hook.enable,omitempty"`
	HookOnFlowReport               string `json:"hook.on_flow_report,omitempty"`
	HookOnHttpAccess               string `json:"hook.on_http_access,omitempty"`
	HookOnPlay                     string `json:"hook.on_play,omitempty"`
	HookOnPublish                  string `json:"hook.on_publish,omitempty"`
	HookOnRecordMp4                string `json:"hook.on_record_mp4,omitempty"`
	HookOnRecordTs                 string `json:"hook.on_record_ts,omitempty"`
	HookOnRtpServerTimeout         string `json:"hook.on_rtp_server_timeout,omitempty"`
	HookOnRtspAuth                 string `json:"hook.on_rtsp_auth,omitempty"`
	HookOnRtspRealm                string `json:"hook.on_rtsp_realm,omitempty"`
	HookOnSendRtpStopped           string `json:"hook.on_send_rtp_stopped,omitempty"`
	HookOnServerExited             string `json:"hook.on_server_exited,omitempty"`
	HookOnServerKeepalive          string `json:"hook.on_server_keepalive,omitempty"`
	HookOnServerStarted            string `json:"hook.on_server_started,omitempty"`
	HookOnShellLogin               string `json:"hook.on_shell_login,omitempty"`
	HookOnStreamChanged            string `json:"hook.on_stream_changed,omitempty"`
	HookOnStreamNoneReader         string `json:"hook.on_stream_none_reader,omitempty"`
	HookOnStreamNotFound           string `json:"hook.on_stream_not_found,omitempty"`
	HookRetry                      string `json:"hook.retry,omitempty"`
	HookRetryDelay                 string `json:"hook.retry_delay,omitempty"`
	HookStreamChangedSchemas       string `json:"hook.stream_changed_schemas,omitempty"`
	HookTimeoutSec                 string `json:"hook.timeoutSec,omitempty"`
	HttpAllowCrossDomains          string `json:"http.allow_cross_domains,omitempty"`
	HttpAllowIpRange               string `json:"http.allow_ip_range,omitempty"`
	HttpCharSet                    string `json:"http.charSet,omitempty"`
	HttpDirMenu                    string `json:"http.dirMenu,omitempty"`
	HttpForbidCacheSuffix          string `json:"http.forbidCacheSuffix,omitempty"`
	HttpForwardedIpHeader          string `json:"http.forwarded_ip_header,omitempty"`
	HttpKeepAliveSecond            string `json:"http.keepAliveSecond,omitempty"`
	HttpMaxReqSize                 string `json:"http.maxReqSize,omitempty"`
	HttpNotFound                   string `json:"http.notFound,omitempty"`
	HttpPort                       string `json:"http.port,omitempty"`
	HttpRootPath                   string `json:"http.rootPath,omitempty"`
	HttpSendBufSize                string `json:"http.sendBufSize,omitempty"`
	HttpSslport                    string `json:"http.sslport,omitempty"`
	HttpVirtualPath                string `json:"http.virtualPath,omitempty"`
	MulticastAddrMax               string `json:"multicast.addrMax,omitempty"`
	MulticastAddrMin               string `json:"multicast.addrMin,omitempty"`
	MulticastUdpTTL                string `json:"multicast.udpTTL,omitempty"`
	ProtocolAddMuteAudio           string `json:"protocol.add_mute_audio,omitempty"`
	ProtocolAutoClose              string `json:"protocol.auto_close,omitempty"`
	ProtocolContinuePushMs         string `json:"protocol.continue_push_ms,omitempty"`
	ProtocolEnableAudio            string `json:"protocol.enable_audio,omitempty"`
	ProtocolEnableFmp4             string `json:"protocol.enable_fmp4,omitempty"`
	ProtocolEnableHls              string `json:"protocol.enable_hls,omitempty"`
	ProtocolEnableHlsFmp4          string `json:"protocol.enable_hls_fmp4,omitempty"`
	ProtocolEnableMp4              string `json:"protocol.enable_mp4,omitempty"`
	ProtocolEnableRtmp             string `json:"protocol.enable_rtmp,omitempty"`
	ProtocolEnableRtsp             string `json:"protocol.enable_rtsp,omitempty"`
	ProtocolEnableTs               string `json:"protocol.enable_ts,omitempty"`
	ProtocolFmp4Demand             string `json:"protocol.fmp4_demand,omitempty"`
	ProtocolHlsDemand              string `json:"protocol.hls_demand,omitempty"`
	ProtocolHlsSavePath            string `json:"protocol.hls_save_path,omitempty"`
	ProtocolModifyStamp            string `json:"protocol.modify_stamp,omitempty"`
	ProtocolMp4AsPlayer            string `json:"protocol.mp4_as_player,omitempty"`
	ProtocolMp4MaxSecond           string `json:"protocol.mp4_max_second,omitempty"`
	ProtocolMp4SavePath            string `json:"protocol.mp4_save_path,omitempty"`
	ProtocolRtmpDemand             string `json:"protocol.rtmp_demand,omitempty"`
	ProtocolRtspDemand             string `json:"protocol.rtsp_demand,omitempty"`
	ProtocolTsDemand               string `json:"protocol.ts_demand,omitempty"`
	RecordAppName                  string `json:"record.appName,omitempty"`
	RecordFastStart                string `json:"record.fastStart,omitempty"`
	RecordFileBufSize              string `json:"record.fileBufSize,omitempty"`
	RecordFileRepeat               string `json:"record.fileRepeat,omitempty"`
	RecordSampleMS                 string `json:"record.sampleMS,omitempty"`
	RtcExternIP                    string `json:"rtc.externIP,omitempty"`
	RtcPort                        string `json:"rtc.port,omitempty"`
	RtcPreferredCodecA             string `json:"rtc.preferredCodecA,omitempty"`
	RtcPreferredCodecV             string `json:"rtc.preferredCodecV,omitempty"`
	RtcRembBitRate                 string `json:"rtc.rembBitRate,omitempty"`
	RtcTcpPort                     string `json:"rtc.tcpPort,omitempty"`
	RtcTimeoutSec                  string `json:"rtc.timeoutSec,omitempty"`
	RtmpHandshakeSecond            string `json:"rtmp.handshakeSecond,omitempty"`
	RtmpKeepAliveSecond            string `json:"rtmp.keepAliveSecond,omitempty"`
	RtmpPort                       string `json:"rtmp.port,omitempty"`
	RtmpSslport                    string `json:"rtmp.sslport,omitempty"`
	RtpAudioMtuSize                string `json:"rtp.audioMtuSize,omitempty"`
	RtpH264StapA                   string `json:"rtp.h264_stap_a,omitempty"`
	RtpLowLatency                  string `json:"rtp.lowLatency,omitempty"`
	RtpRtpMaxSize                  string `json:"rtp.rtpMaxSize,omitempty"`
	RtpVideoMtuSize                string `json:"rtp.videoMtuSize,omitempty"`
	RtpProxyDumpDir                string `json:"rtp_proxy.dumpDir,omitempty"`
	RtpProxyGopCache               string `json:"rtp_proxy.gop_cache,omitempty"`
	RtpProxyH264Pt                 string `json:"rtp_proxy.h264_pt,omitempty"`
	RtpProxyH265Pt                 string `json:"rtp_proxy.h265_pt,omitempty"`
	RtpProxyOpusPt                 string `json:"rtp_proxy.opus_pt,omitempty"`
	RtpProxyPort                   string `json:"rtp_proxy.port,omitempty"`
	RtpProxyPortRange              string `json:"rtp_proxy.port_range,omitempty"`
	RtpProxyPsPt                   string `json:"rtp_proxy.ps_pt,omitempty"`
	RtpProxyTimeoutSec             string `json:"rtp_proxy.timeoutSec,omitempty"`
	RtspAuthBasic                  string `json:"rtsp.authBasic,omitempty"`
	RtspDirectProxy                string `json:"rtsp.directProxy,omitempty"`
	RtspHandshakeSecond            string `json:"rtsp.handshakeSecond,omitempty"`
	RtspKeepAliveSecond            string `json:"rtsp.keepAliveSecond,omitempty"`
	RtspLowLatency                 string `json:"rtsp.lowLatency,omitempty"`
	RtspPort                       string `json:"rtsp.port,omitempty"`
	RtspRtpTransportType           string `json:"rtsp.rtpTransportType,omitempty"`
	RtspSslport                    string `json:"rtsp.sslport,omitempty"`
	ShellMaxReqSize                string `json:"shell.maxReqSize,omitempty"`
	ShellPort                      string `json:"shell.port,omitempty"`
	SrtLatencyMul                  string `json:"srt.latencyMul,omitempty"`
	SrtPktBufSize                  string `json:"srt.pktBufSize,omitempty"`
	SrtPort                        string `json:"srt.port,omitempty"`
	SrtTimeoutSec                  string `json:"srt.timeoutSec,omitempty"`
}

type IndexApiResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

type IndexApiListResp struct {
	IndexApiResp
	Data []string `json:"data"`
}

type IndexApiThreadLoadResp struct {
	IndexApiResp
	Data []ThreadLoad `json:"data"`
}

type IndexApiWorkThreadLoadResp struct {
	IndexApiResp
	Data []ThreadLoad `json:"data"`
}

type IndexApiServerConfigResp struct {
	IndexApiResp
	Data []ServerConfig `json:"data,omitempty"`
}

type IndexApiSetServerConfigResp struct {
	Changed int64 `json:"changed"`
	IndexApiResp
}

type IndexApiRestartServerResp struct {
	IndexApiResp
}

type IndexApiMediaListResp struct {
	IndexApiResp
	Data []MediaList `json:"data"`
}

type IndexApiCloseStreamResp struct {
	IndexApiResp
	Result int64 `json:"result"`
}

type IndexApiCloseStreamsResp struct {
	IndexApiResp
	CountHit    int64 `json:"count_hit"`
	CountClosed int64 `json:"count_closed"`
}

type IndexApiAllSession struct {
	Id        string `json:"id"`
	LocalIp   string `json:"local_ip"`
	LocalPort int64  `json:"local_port""`
	PeerIp    string `json:"peer_ip"`
	PeerPort  int64  `json:"peer_port"`
	Typeid    string `json:"typeid"`
}

type IndexApiAllSessionResp struct {
	IndexApiResp
	Data []IndexApiAllSession `json:"data"`
}

type IndexApiKickSessionResp struct {
	IndexApiResp
}

type IndexApiKickSessionsResp struct {
	IndexApiResp
	CountHit int64 `json:"count_hit"`
}

type IndexApiAddStreamKey struct {
	Key string `json:"key"`
}

type IndexApiAddStreamProxyResp struct {
	IndexApiResp
	Data IndexApiAddStreamKey `json:"data"`
}

type IndexApiAddStreamFlag struct {
	Flag bool `json:"flag"`
}

type IndexApiDelStreamProxyResp struct {
	IndexApiResp
	Data IndexApiAddStreamFlag `json:"data"`
}

type IndexApiAddFFmpegSourceResp struct {
	IndexApiResp
	Data IndexApiAddStreamKey `json:"data"`
}

type IndexApiDelFFmpegSourceResp struct {
	IndexApiResp
	Data IndexApiAddStreamFlag `json:"data"`
}

type IndexApiIsMediaOnlineResp struct {
	IndexApiResp
	Online bool `json:"online"`
}

type IndexApiMediaInfoResp struct {
	IndexApiResp
	Online           bool          `json:"online"`
	ReaderCount      int64         `json:"readerCount"`
	TotalReaderCount int64         `json:"totalReaderCount"`
	Tracks           []StreamTrack `json:"tracks"`
}

type IndexApiRtpInfoResp struct {
	IndexApiResp
	Exist     bool   `json:"exist"`
	PeerIp    string `json:"peer_ip"`
	PeerPort  int64  `json:"peer_port"`
	LocalIp   string `json:"local_ip"`
	LocalPort int64  `json:"local_port"`
}

type IndexApiRecord struct {
	Paths    []string `json:"paths"`
	RootPath string   `json:"rootPath"`
}

type IndexApiMp4RecordFileResp struct {
	IndexApiResp
	Data IndexApiRecord `json:"data"`
}

type IndexApiStartRecordResp struct {
	IndexApiResp
	Result bool `json:"result"`
}

type IndexApiStopRecordResp struct {
	IndexApiResp
	Result bool `json:"result"`
}

type IndexApiIsRecordingResp struct {
	IndexApiResp
	Status bool `json:"status"`
}

type IndexApiSnapResp struct {
	Data []byte `json:"Data"`
}

type IndexApiOpenRtpServerResp struct {
	IndexApiResp
	Port int64 `json:"port"`
}

type IndexApiCloseRtpServerResp struct {
	IndexApiResp
	Hit int64 `json:"hit"`
}

type IndexApiRtp struct {
	Port     int64  `json:"port"`
	StreamId string `json:"stream_id"`
}

type IndexApiListRtpServerResp struct {
	IndexApiResp
	Data []IndexApiRtp `json:"data"`
}

type IndexApiStartSendRtpResp struct {
	IndexApiResp
	LocalProt int64 `json:"local_port"`
}

type IndexApiStartSendRtpPassiveResp struct {
	IndexApiResp
	LocalProt int64 `json:"local_port"`
}

type IndexApiStopSendRtpResp struct {
	IndexApiResp
}

type IndexApiStatisticResp struct {
	IndexApiResp
	Data Statistic `json:"data"`
}

type IndexApiAddStreamPusherProxyResp struct {
	IndexApiResp
	Data IndexApiAddStreamKey `json:"data"`
}

type IndexDelStreamPusherProxyResp struct {
	IndexApiResp
	Data IndexApiAddStreamFlag `json:"data"`
}

type IndexApiVersion struct {
	BranchName string `json:"branchName"`
	BuildTime  string `json:"buildTime"`
	CommitHash string `json:"commitHash"`
}

type IndexApiVersionResp struct {
	IndexApiResp
	Data IndexApiVersion `json:"data"`
}

type IndexApiMediaPlayerListResp struct {
	IndexApiResp
	Data MediaPlayer `json:"data"`
}
