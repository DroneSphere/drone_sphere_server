package mavic3

import "time"

// PropertyMessage 用于接受MQTT消息的结构体
type PropertyMessage struct {
	Country                  string                  `json:"country"`                               // 国家区域码
	ModeCode                 ModeCode                `json:"mode_code"`                             // 飞行器状态
	ModeCodeReason           ModeCodeReason          `json:"mode_code_reason"`                      // 飞行器进入当前状态的原因
	Cameras                  []Camera                `json:"cameras"`                               // 飞行器相机信息
	DongleInfos              []DongleInfo            `json:"dongle_infos"`                          // 4G Dongle信息
	ObstacleAvoidance        ObstacleAvoidance       `json:"obstacle_avoidance"`                    // 飞行器避障状态
	IsNearAreaLimit          int                     `json:"is_near_area_limit"`                    // 是否接近限飞区
	IsNearHeightLimit        int                     `json:"is_near_height_limit"`                  // 是否接近设定的限制高度
	HeightLimit              int                     `json:"height_limit"`                          // 飞行器限高
	NightLightsState         int                     `json:"night_lights_state"`                    // 飞行器夜航灯状态
	ActivationTime           int64                   `json:"activation_time"`                       // 飞行器激活时间（Unix时间戳）
	MaintainStatus           MaintainStatus          `json:"maintain_status"`                       // 保养信息
	TotalFlightSorties       int                     `json:"total_flight_sorties"`                  // 飞行器累计飞行总架次
	TrackID                  string                  `json:"track_id"`                              // 航迹ID
	PositionState            PositionState           `json:"position_state"`                        // 搜星状态
	Storage                  Storage                 `json:"storage"`                               // 存储容量
	Battery                  Battery                 `json:"battery"`                               // 飞行器电池信息
	TotalFlightDistance      float64                 `json:"total_flight_distance"`                 // 飞行器累计飞行总里程
	TotalFlightTime          int                     `json:"total_flight_time"`                     // 飞行器累计飞行航时
	SeriousLowBatteryWarning int                     `json:"serious_low_battery_warning_threshold"` // 严重低电量告警
	LowBatteryWarning        int                     `json:"low_battery_warning_threshold"`         // 低电量告警
	ControlSource            string                  `json:"control_source"`                        // 当前控制源
	WindDirection            int                     `json:"wind_direction"`                        // 当前风向
	WindSpeed                float64                 `json:"wind_speed"`                            // 风速
	HomeDistance             float64                 `json:"home_distance"`                         // 距离Home点的距离
	HomeLatitude             float64                 `json:"home_latitude"`                         // Home点纬度
	HomeLongitude            float64                 `json:"home_longitude"`                        // Home点经度
	AttitudeHead             int                     `json:"attitude_head"`                         // 偏航轴角度
	AttitudeRoll             float64                 `json:"attitude_roll"`                         // 横滚轴角度
	AttitudePitch            float64                 `json:"attitude_pitch"`                        // 俯仰轴角度
	Elevation                float64                 `json:"elevation"`                             // 相对起飞点高度
	Height                   float64                 `json:"height"`                                // 绝对高度
	Latitude                 float64                 `json:"latitude"`                              // 当前位置纬度
	Longitude                float64                 `json:"longitude"`                             // 当前位置经度
	VerticalSpeed            float64                 `json:"vertical_speed"`                        // 垂直速度
	HorizontalSpeed          float64                 `json:"horizontal_speed"`                      // 水平速度
	FirmwareUpgradeStatus    int                     `json:"firmware_upgrade_status"`               // 固件升级状态
	CompatibleStatus         int                     `json:"compatible_status"`                     // 固件一致性
	FirmwareVersion          string                  `json:"firmware_version"`                      // 固件版本
	Gear                     int                     `json:"gear"`                                  // 档位
	CameraWatermarkSettings  CameraWatermarkSettings `json:"camera_watermark_settings"`             // 相机水印设置
}

// ModeCode 飞行器状态
type ModeCode int

const (
	ModeCodeStandby               ModeCode = 0  // 待机
	ModeCodeTakeoff               ModeCode = 1  // 起飞准备
	ModeCodeTakeoffReady          ModeCode = 2  // 起飞准备完毕
	ModeCodeManualFlight          ModeCode = 3  // 手动飞行
	ModeCodeAutoTakeoff           ModeCode = 4  // 自动起飞
	ModeCodeRouteFlight           ModeCode = 5  // 航线飞行
	ModeCodePanoramaPhoto         ModeCode = 6  // 全景拍照
	ModeCodeSmartFollow           ModeCode = 7  // 智能跟随
	ModeCodeADSBAvoidance         ModeCode = 8  // ADS-B 躲避
	ModeCodeAutoReturn            ModeCode = 9  // 自动返航
	ModeCodeAutoLanding           ModeCode = 10 // 自动降落
	ModeCodeForceLanding          ModeCode = 11 // 强制降落
	ModeCodeThreeBladeLanding     ModeCode = 12 // 三桨叶降落
	ModeCodeUpgrading             ModeCode = 13 // 升级中
	ModeCodeDisconnected          ModeCode = 14 // 未连接
	ModeCodeAPAS                  ModeCode = 15 // APAS
	ModeCodeVirtualJoystick       ModeCode = 16 // 虚拟摇杆状态
	ModeCodeCommandFlight         ModeCode = 17 // 指令飞行
	ModeCodeAirRTKConvergenceMode ModeCode = 18 // 空中 RTK 收敛模式
)

// ModeCodeReason 飞行器进入当前状态的原因
type ModeCodeReason int

const (
	ModeCodeReasonNone                         ModeCodeReason = 0  // 无意义
	ModeCodeReasonLowBattery                   ModeCodeReason = 1  // 电池电量不足（返航、降落）
	ModeCodeReasonLowBatteryVoltage            ModeCodeReason = 2  // 电池电压不足（返航、降落）
	ModeCodeReasonLowBatteryVoltageSerious     ModeCodeReason = 3  // 电压严重过低（返航、降落）
	ModeCodeReasonRemoteControlKeyRequest      ModeCodeReason = 4  // 遥控器按键请求（起飞、返航、降落）
	ModeCodeReasonAppRequest                   ModeCodeReason = 5  // App 请求（起飞、返航、降落）
	ModeCodeReasonRemoteSignalLost             ModeCodeReason = 6  // 遥控信号丢失（返航、降落、悬停）
	ModeCodeReasonNavigationSDKTrigger         ModeCodeReason = 7  // 导航、SDK 等外部设备触发（起飞、返航、降落）
	ModeCodeReasonEnterAirportLimitArea        ModeCodeReason = 8  // 进入机场限飞区（降落）
	ModeCodeReasonReturnHomeDistanceTooClose   ModeCodeReason = 9  // 虽然触发了返航但是因为距离 Home 点距离太近（降落）
	ModeCodeReasonReturnHomeDistanceTooFar     ModeCodeReason = 10 // 虽然触发了返航但是因为距离 Home 点距离太远（降落）
	ModeCodeReasonWaypointTaskRequest          ModeCodeReason = 11 // 执行航点任务时请求（起飞）
	ModeCodeReasonReturnHomeAboveHomePoint     ModeCodeReason = 12 // 返航阶段到达 Home 点上方后请求（降落）
	ModeCodeReasonHeightDecrease               ModeCodeReason = 13 // 飞行器高度下降，距地面 0.7m（二段降落限低）时，继续下降导致（降落）
	ModeCodeReasonForceBreakLowLimitProtection ModeCodeReason = 14 // App、SDK 等设备强制突破限低保护进行（降落）
	ModeCodeReasonFlightPassingBy              ModeCodeReason = 15 // 因为周围有航班经过而请求（返航、降落）
	ModeCodeReasonHeightControlFailed          ModeCodeReason = 16 // 因为高度控制失败请求（返航、降落）
	ModeCodeReasonSmartLowBatteryReturn        ModeCodeReason = 17 // 智能低电量返航后进入（降落）
	ModeCodeReasonAPControlFlightMode          ModeCodeReason = 18 // AP控制飞行模式（手动飞行）
	ModeCodeReasonHardwareException            ModeCodeReason = 19 // 硬件异常（返航、降落）
	ModeCodeReasonAntiTouchGroundProtection    ModeCodeReason = 20 // 防触地保护结束（降落）
	ModeCodeReasonReturnHomeCancel             ModeCodeReason = 21 // 返航取消（悬停）
	ModeCodeReasonReturnHomeObstacle           ModeCodeReason = 22 // 返航时遇到障碍物（降落）
)

// Camera 飞行器相机信息
type Camera struct {
	RemainPhotoNum                  int                 `json:"remain_photo_num"`                    // 剩余拍照张数
	RemainRecordDuration            int                 `json:"remain_record_duration"`              // 剩余录像时间
	RecordTime                      int                 `json:"record_time"`                         // 视频录制时长
	PayloadIndex                    string              `json:"payload_index"`                       // 负载编号
	CameraMode                      CameraMode          `json:"camera_mode"`                         // 相机模式
	PhotoState                      PhotoState          `json:"photo_state"`                         // 拍照状态
	ScreenSplitEnable               bool                `json:"screen_split_enable"`                 // 分屏是否使能
	RecordingState                  int                 `json:"recording_state"`                     // 录像状态
	ZoomFactor                      float64             `json:"zoom_factor"`                         // 变焦倍数
	IRZoomFactor                    float64             `json:"ir_zoom_factor"`                      // 红外变焦倍数
	LiveviewWorldRegion             LiveviewWorldRegion `json:"liveview_world_region"`               // 视场角（FOV）在liveview中的区域
	PhotoStorageSettings            []string            `json:"photo_storage_settings"`              // 照片存储设置集合
	VideoStorageSettings            []string            `json:"video_storage_settings"`              // 视频存储设置集合
	WideExposureMode                WideExposureMode    `json:"wide_exposure_mode"`                  // 广角镜头曝光模式
	WideISO                         int                 `json:"wide_iso"`                            // 广角镜头感光度
	WideShutterSpeed                int                 `json:"wide_shutter_speed"`                  // 广角镜头快门速度
	WideExposureValue               int                 `json:"wide_exposure_value"`                 // 广角镜头曝光值
	ZoomExposureMode                int                 `json:"zoom_exposure_mode"`                  // 变焦镜头曝光模式
	ZoomISO                         int                 `json:"zoom_iso"`                            // 变焦镜头感光度
	ZoomShutterSpeed                int                 `json:"zoom_shutter_speed"`                  // 变焦镜头快门速度
	ZoomExposureValue               int                 `json:"zoom_exposure_value"`                 // 变焦镜头曝光值
	ZoomFocusMode                   int                 `json:"zoom_focus_mode"`                     // 变焦镜头对焦模式
	ZoomFocusValue                  int                 `json:"zoom_focus_value"`                    // 变焦镜头对焦值
	ZoomMaxFocusValue               int                 `json:"zoom_max_focus_value"`                // 变焦镜头最大对焦值
	ZoomMinFocusValue               int                 `json:"zoom_min_focus_value"`                // 变焦镜头最小对焦值
	ZoomCalibrateFarthestFocusValue int                 `json:"zoom_calibrate_farthest_focus_value"` // 变焦镜头标定的最远对焦值
	ZoomCalibrateNearestFocusValue  int                 `json:"zoom_calibrate_nearest_focus_value"`  // 变焦镜头标定的最近对焦值
	ZoomFocusState                  int                 `json:"zoom_focus_state"`                    // 变焦镜头对焦状态
	IRMeteringMode                  int                 `json:"ir_metering_mode"`                    // 红外测温模式
	IRMeteringPoint                 IRMeteringPoint     `json:"ir_metering_point"`                   // 红外测温点
	IRMeteringArea                  IRMeteringArea      `json:"ir_metering_area"`                    // 红外测温区域
}

type CameraMode int

const (
	CameraModePhoto    CameraMode = 0 // 拍照
	CameraModeVideo    CameraMode = 1 // 录像
	CameraModeLowLight CameraMode = 2 // 智能低光
	CameraModePanorama CameraMode = 3 // 全景拍照
)

type PhotoState int

const (
	PhotoStateIdle       PhotoState = 0 // 空闲
	PhotoStatePhotograph PhotoState = 1 // 拍照中
)

type RecordState int

const (
	RecordStateIdle      RecordState = 0 // 空闲
	RecordStateRecording RecordState = 1 // 录像中
)

type WideExposureMode int

const (
	WideExposureModeAuto             WideExposureMode = 1 // 自动
	WideExposureModeShutterPriority  WideExposureMode = 2 // 快门优先曝光
	WideExposureModeAperturePriority WideExposureMode = 3 // 光圈优先曝光
	WideExposureModeManual           WideExposureMode = 4 // 手动曝光
)

// LiveviewWorldRegion 视场角（FOV）在liveview中的区域
type LiveviewWorldRegion struct {
	Left   float64 `json:"left"`   // 左上角的x轴起始点
	Top    float64 `json:"top"`    // 左上角的y轴起始点
	Right  float64 `json:"right"`  // 右下角的x轴起始点
	Bottom float64 `json:"bottom"` // 右下角的y轴起始点
}

// IRMeteringPoint 红外测温点
type IRMeteringPoint struct {
	X           float64 `json:"x"`           // 测温点坐标x
	Y           float64 `json:"y"`           // 测温点坐标y
	Temperature float64 `json:"temperature"` // 测温点的温度
}

// IRMeteringArea 红外测温区域
type IRMeteringArea struct {
	X                   float64          `json:"x"`                     // 测温区域左上角点坐标x
	Y                   float64          `json:"y"`                     // 测温区域左上角点坐标y
	Width               float64          `json:"width"`                 // 测温区域宽度
	Height              float64          `json:"height"`                // 测温区域高度
	AverTemperature     float64          `json:"aver_temperature"`      // 测温区域平均温度
	MinTemperaturePoint TemperaturePoint `json:"min_temperature_point"` // 测温区域最低温度点
	MaxTemperaturePoint TemperaturePoint `json:"max_temperature_point"` // 测温区域最高温度点
}

// TemperaturePoint 温度点信息
type TemperaturePoint struct {
	X           float64 `json:"x"`           // 温度点坐标x
	Y           float64 `json:"y"`           // 温度点坐标y
	Temperature float64 `json:"temperature"` // 温度点的温度
}

// DongleInfo 4G Dongle信息
type DongleInfo struct {
	IMEI              string     `json:"imei"`                // Dongle的唯一识别标志
	DongleType        int        `json:"dongle_type"`         // Dongle类型
	EID               string     `json:"eid"`                 // eSIM的唯一识别标志
	ESIMActivateState int        `json:"esim_activate_state"` // eSIM激活状态
	SIMCardState      int        `json:"sim_card_state"`      // SIM卡状态
	SIMSlot           int        `json:"sim_slot"`            // SIM卡槽使能状态
	ESIMInfos         []ESIMInfo `json:"esim_infos"`          // eSIM信息
	SIMInfo           SIMInfo    `json:"sim_info"`            // SIM卡信息
}

// ESIMInfo eSIM信息
type ESIMInfo struct {
	TelecomOperator int    `json:"telecom_operator"` // 支持的运营商
	Enabled         bool   `json:"enabled"`          // eSIM使能状态
	ICCID           string `json:"iccid"`            // SIM卡的唯一识别标志
}

// SIMInfo SIM卡信息
type SIMInfo struct {
	TelecomOperator int    `json:"telecom_operator"` // 支持的运营商
	SIMType         int    `json:"sim_type"`         // SIM卡类型
	ICCID           string `json:"iccid"`            // SIM卡的唯一识别标志
}

// ObstacleAvoidance 飞行器避障状态
type ObstacleAvoidance struct {
	Horizon  int `json:"horizon"`  // 水平避障状态
	Upside   int `json:"upside"`   // 上视避障状态
	Downside int `json:"downside"` // 下视避障状态
}

// MaintainStatus 保养信息
type MaintainStatus struct {
	MaintainStatusArray []MaintainStatusItem `json:"maintain_status_array"` // 保养信息数组
}

// MaintainStatusItem 保养信息项
type MaintainStatusItem struct {
	State                     int       `json:"state"`                        // 保养状态
	LastMaintainType          int       `json:"last_maintain_type"`           // 上一次保养类型
	LastMaintainTime          time.Time `json:"last_maintain_time"`           // 上一次保养时间
	LastMaintainFlightTime    int       `json:"last_maintain_flight_time"`    // 上一次保养时飞行航时
	LastMaintainFlightSorties int       `json:"last_maintain_flight_sorties"` // 上一次保养时飞行架次
}

// PositionState 搜星状态
type PositionState struct {
	IsFixed   int `json:"is_fixed"`   // 是否收敛
	Quality   int `json:"quality"`    // 搜星档位
	GPSNumber int `json:"gps_number"` // GPS搜星数量
	RTKNumber int `json:"rtk_number"` // RTK搜星数量
}

// Storage 存储容量
type Storage struct {
	Total int `json:"total"` // 总容量
	Used  int `json:"used"`  // 已使用容量
}

// Battery 飞行器电池信息
type Battery struct {
	CapacityPercent  int             `json:"capacity_percent"`   // 电池的总剩余电量
	RemainFlightTime int             `json:"remain_flight_time"` // 剩余飞行时间
	ReturnHomePower  int             `json:"return_home_power"`  // 返航所需电量百分比
	LandingPower     int             `json:"landing_power"`      // 强制降落电量百分比
	Batteries        []BatteryDetail `json:"batteries"`          // 电池详细信息
}

// BatteryDetail 电池详细信息
type BatteryDetail struct {
	CapacityPercent        int     `json:"capacity_percent"`          // 电池剩余电量
	Index                  int     `json:"index"`                     // 电池序号
	SN                     string  `json:"sn"`                        // 电池序列号（SN）
	Type                   int     `json:"type"`                      // 电池类型
	SubType                int     `json:"sub_type"`                  // 电池子类型
	FirmwareVersion        string  `json:"firmware_version"`          // 固件版本
	LoopTimes              int     `json:"loop_times"`                // 电池循环次数
	Voltage                int     `json:"voltage"`                   // 电压
	Temperature            float64 `json:"temperature"`               // 温度
	HighVoltageStorageDays int     `json:"high_voltage_storage_days"` // 高电压存储天数
}

// CameraWatermarkSettings 相机水印设置
type CameraWatermarkSettings struct {
	GlobalEnable           int    `json:"global_enable"`             // 水印显示全局开关
	DroneTypeEnable        int    `json:"drone_type_enable"`         // 机型显示开关
	DroneSNEnable          int    `json:"drone_sn_enable"`           // 飞行器序列号显示开关
	DatetimeEnable         int    `json:"datetime_enable"`           // 日期时间显示开关
	GPSEnable              int    `json:"gps_enable"`                // 经纬度&海拔高显示开关
	UserCustomStringEnable int    `json:"user_custom_string_enable"` // 自定义文案显示开关
	UserCustomString       string `json:"user_custom_string"`        // 自定义文案内容
	Layout                 int    `json:"layout"`                    // 水印在画面中位置
}
