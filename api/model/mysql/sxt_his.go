package mysql

import "time"

// 病例表 his_care_history
type HisCareHistory struct {
	ChID         string     `gorm:"column:ch_id;primaryKey" json:"ch_id"`
	UserID       int64      `gorm:"column:user_id" json:"user_id"`
	UserName     string     `gorm:"column:user_name" json:"user_name"`
	PatientID    string     `gorm:"column:patient_id" json:"patient_id"`
	PatientName  string     `gorm:"column:patient_name" json:"patient_name"`
	DeptID       int64      `gorm:"column:dept_id" json:"dept_id"`
	DeptName     string     `gorm:"column:dept_name" json:"dept_name"`
	ReceiveType  string     `gorm:"column:receive_type" json:"receive_type"`
	IsContagious string     `gorm:"column:is_contagious" json:"is_contagious"`
	CareTime     *time.Time `gorm:"column:care_time" json:"care_time"`
	CaseDate     string     `gorm:"column:case_date" json:"case_date"`
	RegID        string     `gorm:"column:reg_id" json:"reg_id"`
	CaseTitle    string     `gorm:"column:case_title" json:"case_title"`
	CaseResult   string     `gorm:"column:case_result" json:"case_result"`
	DoctorTips   string     `gorm:"column:doctor_tips" json:"doctor_tips"`
	Remark       string     `gorm:"column:remark" json:"remark"`
}

func (HisCareHistory) TableName() string {
	return "his_care_history"
}

// 药用处方表 his_care_order
type HisCareOrder struct {
	CoID        string     `gorm:"column:co_id;primaryKey" json:"co_id"`
	CoType      string     `gorm:"column:co_type" json:"co_type"`
	UserID      int64      `gorm:"column:user_id" json:"user_id"`
	PatientID   string     `gorm:"column:patient_id" json:"patient_id"`
	PatientName string     `gorm:"column:patient_name" json:"patient_name"`
	ChID        string     `gorm:"column:ch_id" json:"ch_id"`
	AllAmount   float64    `gorm:"column:all_amount" json:"all_amount"`
	CreateBy    string     `gorm:"column:create_by" json:"create_by"`
	CreateTime  *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateBy    string     `gorm:"column:update_by" json:"update_by"`
	UpdateTime  *time.Time `gorm:"column:update_time" json:"update_time"`
}

func (HisCareOrder) TableName() string {
	return "his_care_order"
}

// 开诊细表 his_care_order_item
type HisCareOrderItem struct {
	ItemID     string     `gorm:"column:item_id;primaryKey" json:"item_id"`
	CoID       string     `gorm:"column:co_id" json:"co_id"`
	ItemRefID  string     `gorm:"column:item_ref_id" json:"item_ref_id"`
	ItemName   string     `gorm:"column:item_name" json:"item_name"`
	ItemType   string     `gorm:"column:item_type" json:"item_type"`
	Num        float64    `gorm:"column:num" json:"num"`
	Price      float64    `gorm:"column:price" json:"price"`
	Amount     float64    `gorm:"column:amount" json:"amount"`
	Remark     string     `gorm:"column:remark" json:"remark"`
	Status     string     `gorm:"column:status" json:"status"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
}

func (HisCareOrderItem) TableName() string {
	return "his_care_order_item"
}

// 检查结果表 his_check_result
type HisCheckResult struct {
	CocID         string     `gorm:"column:coc_id;primaryKey" json:"coc_id"`
	CheckItemID   int        `gorm:"column:check_item_id" json:"check_item_id"`
	CheckItemName string     `gorm:"column:check_item_name" json:"check_item_name"`
	Price         float64    `gorm:"column:price" json:"price"`
	RegID         string     `gorm:"column:reg_id" json:"reg_id"`
	ResultMsg     string     `gorm:"column:result_msg" json:"result_msg"`
	ResultImg     string     `gorm:"column:result_img" json:"result_img"`
	PatientID     string     `gorm:"column:patient_id" json:"patient_id"`
	PatientName   string     `gorm:"column:patient_name" json:"patient_name"`
	ResultStatus  string     `gorm:"column:result_status" json:"result_status"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy      string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy      string     `gorm:"column:update_by" json:"update_by"`
}

func (HisCheckResult) TableName() string {
	return "his_check_result"
}

// 退费主表 his_order_backfee
type HisOrderBackfee struct {
	BackID         string     `gorm:"column:back_id;primaryKey" json:"back_id"`
	BackAmount     float64    `gorm:"column:back_amount" json:"back_amount"`
	ChID           string     `gorm:"column:ch_id" json:"ch_id"`
	RegID          string     `gorm:"column:reg_id" json:"reg_id"`
	PatientName    string     `gorm:"column:patient_name" json:"patient_name"`
	BackStatus     string     `gorm:"column:back_status" json:"back_status"`
	BackType       string     `gorm:"column:back_type" json:"back_type"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	BackPlatformID string     `gorm:"column:back_platform_id" json:"back_platform_id"`
	BackTime       *time.Time `gorm:"column:back_time" json:"back_time"`
	CreateTime     *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy       string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy       string     `gorm:"column:update_by" json:"update_by"`
}

func (HisOrderBackfee) TableName() string {
	return "his_order_backfee"
}

// 退费订单详情表 his_order_backfee_item
type HisOrderBackfeeItem struct {
	ItemID     string  `gorm:"column:item_id;primaryKey" json:"item_id"`
	CoID       string  `gorm:"column:co_id" json:"co_id"`
	ItemName   string  `gorm:"column:item_name" json:"item_name"`
	ItemPrice  float64 `gorm:"column:item_price" json:"item_price"`
	ItemNum    int     `gorm:"column:item_num" json:"item_num"`
	ItemAmount float64 `gorm:"column:item_amount" json:"item_amount"`
	BackID     string  `gorm:"column:back_id" json:"back_id"`
	ItemType   string  `gorm:"column:item_type" json:"item_type"`
	Status     string  `gorm:"column:status" json:"status"`
}

func (HisOrderBackfeeItem) TableName() string {
	return "his_order_backfee_item"
}

// 收费表 his_order_charge
type HisOrderCharge struct {
	OrderID       string     `gorm:"column:order_id;primaryKey" json:"order_id"`
	OrderAmount   float64    `gorm:"column:order_amount" json:"order_amount"`
	ChID          string     `gorm:"column:ch_id" json:"ch_id"`
	RegID         string     `gorm:"column:reg_id" json:"reg_id"`
	PatientName   string     `gorm:"column:patient_name" json:"patient_name"`
	OrderStatus   string     `gorm:"column:order_status" json:"order_status"`
	PayPlatformID string     `gorm:"column:pay_platform_id" json:"pay_platform_id"`
	PayTime       *time.Time `gorm:"column:pay_time" json:"pay_time"`
	PayType       string     `gorm:"column:pay_type" json:"pay_type"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy      string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy      string     `gorm:"column:update_by" json:"update_by"`
}

func (HisOrderCharge) TableName() string {
	return "his_order_charge"
}

// 支付订单详情表 his_order_charge_item
type HisOrderChargeItem struct {
	ItemID     string  `gorm:"column:item_id;primaryKey" json:"item_id"`
	CoID       string  `gorm:"column:co_id" json:"co_id"`
	ItemName   string  `gorm:"column:item_name" json:"item_name"`
	ItemPrice  float64 `gorm:"column:item_price" json:"item_price"`
	ItemNum    int     `gorm:"column:item_num" json:"item_num"`
	ItemAmount float64 `gorm:"column:item_amount" json:"item_amount"`
	OrderID    string  `gorm:"column:order_id" json:"order_id"`
	ItemType   string  `gorm:"column:item_type" json:"item_type"`
	Status     string  `gorm:"column:status" json:"status"`
}

func (HisOrderChargeItem) TableName() string {
	return "his_order_charge_item"
}

// 患者信息表 his_patient
type HisPatient struct {
	PatientID     string     `gorm:"column:patient_id;primaryKey" json:"patient_id"`
	Name          string     `gorm:"column:name" json:"name"`
	Phone         string     `gorm:"column:phone" json:"phone"`
	Sex           string     `gorm:"column:sex" json:"sex"`
	Birthday      string     `gorm:"column:birthday" json:"birthday"`
	IDCard        string     `gorm:"column:id_card" json:"id_card"`
	Address       string     `gorm:"column:address" json:"address"`
	AllergyInfo   string     `gorm:"column:allergy_info" json:"allergy_info"`
	IsFinal       string     `gorm:"column:is_final" json:"is_final"`
	Password      string     `gorm:"column:password" json:"password"`
	Openid        string     `gorm:"column:openid" json:"openid"`
	LastLoginIP   string     `gorm:"column:last_login_ip" json:"last_login_ip"`
	LastLoginTime *time.Time `gorm:"column:last_login_time" json:"last_login_time"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    *time.Time `gorm:"column:update_time" json:"update_time"`
}

func (HisPatient) TableName() string {
	return "his_patient"
}

// 患者档案表 his_patient_file
type HisPatientFile struct {
	PatientID                string     `gorm:"column:patient_id;primaryKey" json:"patient_id"`
	EmergencyContactName     string     `gorm:"column:emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactPhone    string     `gorm:"column:emergency_contact_phone" json:"emergency_contact_phone"`
	EmergencyContactRelation string     `gorm:"column:emergency_contact_relation" json:"emergency_contact_relation"`
	LeftEarHearing           string     `gorm:"column:left_ear_hearing" json:"left_ear_hearing"`
	RightEarHearing          string     `gorm:"column:right_ear_hearing" json:"right_ear_hearing"`
	LeftVision               float64    `gorm:"column:left_vision" json:"left_vision"`
	RightVision              float64    `gorm:"column:right_vision" json:"right_vision"`
	Height                   float64    `gorm:"column:height" json:"height"`
	Weight                   float64    `gorm:"column:weight" json:"weight"`
	BloodType                string     `gorm:"column:blood_type" json:"blood_type"`
	PersonalInfo             string     `gorm:"column:personal_info" json:"personal_info"`
	FamilyInfo               string     `gorm:"column:family_info" json:"family_info"`
	CreateTime               *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime               *time.Time `gorm:"column:update_time" json:"update_time"`
}

func (HisPatientFile) TableName() string {
	return "his_patient_file"
}

// 挂号表 his_registration
type HisRegistration struct {
	RegistrationID     string     `gorm:"column:registration_id;primaryKey" json:"registration_id"`
	PatientID          string     `gorm:"column:patient_id" json:"patient_id"`
	PatientName        string     `gorm:"column:patient_name" json:"patient_name"`
	UserID             int64      `gorm:"column:user_id" json:"user_id"`
	DoctorName         string     `gorm:"column:doctor_name" json:"doctor_name"`
	DeptID             int64      `gorm:"column:dept_id" json:"dept_id"`
	RegistrationItemID int64      `gorm:"column:registration_item_id" json:"registration_item_id"`
	RegistrationAmount float64    `gorm:"column:registration_amount" json:"registration_amount"`
	RegistrationNumber int        `gorm:"column:registration_number" json:"registration_number"`
	RegistrationStatus string     `gorm:"column:registration_status" json:"registration_status"`
	VisitDate          string     `gorm:"column:visit_date" json:"visit_date"`
	SchedulingType     string     `gorm:"column:scheduling_type" json:"scheduling_type"`
	SubsectionType     string     `gorm:"column:subsection_type" json:"subsection_type"`
	CreateTime         *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime         *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy           string     `gorm:"column:create_by" json:"create_by"`
}

func (HisRegistration) TableName() string {
	return "his_registration"
}

// 排班信息表 his_scheduling
type HisScheduling struct {
	UserID         int        `gorm:"column:user_id;primaryKey" json:"user_id"`
	DeptID         int        `gorm:"column:dept_id;primaryKey" json:"dept_id"`
	SchedulingDay  string     `gorm:"column:scheduling_day;primaryKey" json:"scheduling_day"`
	SubsectionType string     `gorm:"column:subsection_type" json:"subsection_type"`
	SchedulingType string     `gorm:"column:scheduling_type" json:"scheduling_type"`
	CreateTime     *time.Time `gorm:"column:create_time" json:"create_time"`
	CreateBy       string     `gorm:"column:create_by" json:"create_by"`
}

func (HisScheduling) TableName() string {
	return "his_scheduling"
}

// 验证码表 his_verification_code
type HisVerificationCode struct {
	ID               int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	VerificationCode int        `gorm:"column:verification_code" json:"verification_code"`
	PhoneNumber      string     `gorm:"column:phone_number" json:"phone_number"`
	CreateTime       *time.Time `gorm:"column:create_time" json:"create_time"`
	IsCheck          bool       `gorm:"column:is_check" json:"is_check"`
}

func (HisVerificationCode) TableName() string {
	return "his_verification_code"
}

// 入库日志表 stock_inventory_log
type StockInventoryLog struct {
	InventoryLogID          string     `gorm:"column:inventory_log_id;primaryKey" json:"inventory_log_id"`
	PurchaseID              string     `gorm:"column:purchase_id" json:"purchase_id"`
	MedicinesID             string     `gorm:"column:medicines_id" json:"medicines_id"`
	InventoryLogNum         int        `gorm:"column:inventory_log_num" json:"inventory_log_num"`
	TradePrice              float64    `gorm:"column:trade_price" json:"trade_price"`
	PrescriptionPrice       float64    `gorm:"column:prescription_price" json:"prescription_price"`
	TradeTotalAmount        float64    `gorm:"column:trade_total_amount" json:"trade_total_amount"`
	PrescriptionTotalAmount float64    `gorm:"column:prescription_total_amount" json:"prescription_total_amount"`
	BatchNumber             string     `gorm:"column:batch_number" json:"batch_number"`
	MedicinesName           string     `gorm:"column:medicines_name" json:"medicines_name"`
	MedicinesType           string     `gorm:"column:medicines_type" json:"medicines_type"`
	PrescriptionType        string     `gorm:"column:prescription_type" json:"prescription_type"`
	ProducterID             string     `gorm:"column:producter_id" json:"producter_id"`
	Conversion              int        `gorm:"column:conversion" json:"conversion"`
	Unit                    string     `gorm:"column:unit" json:"unit"`
	ProviderID              string     `gorm:"column:provider_id" json:"provider_id"`
	CreateTime              *time.Time `gorm:"column:create_time" json:"create_time"`
	CreateBy                string     `gorm:"column:create_by" json:"create_by"`
}

func (StockInventoryLog) TableName() string {
	return "stock_inventory_log"
}

// 药品信息表 stock_medicines
type StockMedicines struct {
	MedicinesID             uint64     `gorm:"column:medicines_id;primaryKey;autoIncrement" json:"medicines_id"`
	MedicinesNumber         string     `gorm:"column:medicines_number" json:"medicines_number"`
	MedicinesName           string     `gorm:"column:medicines_name" json:"medicines_name"`
	MedicinesType           string     `gorm:"column:medicines_type" json:"medicines_type"`
	PrescriptionType        string     `gorm:"column:prescription_type" json:"prescription_type"`
	PrescriptionPrice       float64    `gorm:"column:prescription_price" json:"prescription_price"`
	Unit                    string     `gorm:"column:unit" json:"unit"`
	Conversion              int        `gorm:"column:conversion" json:"conversion"`
	Keywords                string     `gorm:"column:keywords" json:"keywords"`
	ProducterID             string     `gorm:"column:producter_id" json:"producter_id"`
	Status                  string     `gorm:"column:status" json:"status"`
	MedicinesStockNum       float64    `gorm:"column:medicines_stock_num" json:"medicines_stock_num"`
	MedicinesStockDangerNum float64    `gorm:"column:medicines_stock_danger_num" json:"medicines_stock_danger_num"`
	CreateTime              *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime              *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy                string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy                string     `gorm:"column:update_by" json:"update_by"`
	DelFlag                 string     `gorm:"column:del_flag" json:"del_flag"`
}

func (StockMedicines) TableName() string {
	return "stock_medicines"
}

// 生产厂家表 stock_producer
type StockProducer struct {
	ProducerID      int64      `gorm:"column:producer_id;primaryKey;autoIncrement" json:"producer_id"`
	ProducerName    string     `gorm:"column:producer_name" json:"producer_name"`
	ProducerCode    string     `gorm:"column:producer_code" json:"producer_code"`
	ProducerAddress string     `gorm:"column:producer_address" json:"producer_address"`
	ProducerTel     string     `gorm:"column:producer_tel" json:"producer_tel"`
	ProducerPerson  string     `gorm:"column:producer_person" json:"producer_person"`
	Keywords        string     `gorm:"column:keywords" json:"keywords"`
	Status          string     `gorm:"column:status" json:"status"`
	CreateTime      *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy        string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy        string     `gorm:"column:update_by" json:"update_by"`
}

func (StockProducer) TableName() string {
	return "stock_producer"
}

// 供应商信息表 stock_provider
type StockProvider struct {
	ProviderID      int64      `gorm:"column:provider_id;primaryKey;autoIncrement" json:"provider_id"`
	ProviderName    string     `gorm:"column:provider_name" json:"provider_name"`
	ContactName     string     `gorm:"column:contact_name" json:"contact_name"`
	ContactMobile   string     `gorm:"column:contact_mobile" json:"contact_mobile"`
	ContactTel      string     `gorm:"column:contact_tel" json:"contact_tel"`
	BankAccount     string     `gorm:"column:bank_account" json:"bank_account"`
	ProviderAddress string     `gorm:"column:provider_address" json:"provider_address"`
	Status          string     `gorm:"column:status" json:"status"`
	DelFlag         string     `gorm:"column:del_flag" json:"del_flag"`
	CreateTime      *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy        string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy        string     `gorm:"column:update_by" json:"update_by"`
}

func (StockProvider) TableName() string {
	return "stock_provider"
}

// 采购单 stock_purchase
type StockPurchase struct {
	PurchaseID               string     `gorm:"column:purchase_id;primaryKey" json:"purchase_id"`
	ProviderID               string     `gorm:"column:provider_id" json:"provider_id"`
	PurchaseTradeTotalAmount float64    `gorm:"column:purchase_trade_total_amount" json:"purchase_trade_total_amount"`
	Status                   string     `gorm:"column:status" json:"status"`
	ApplyUserID              int64      `gorm:"column:apply_user_id" json:"apply_user_id"`
	ApplyUserName            string     `gorm:"column:apply_user_name" json:"apply_user_name"`
	StorageOptUser           string     `gorm:"column:storage_opt_user" json:"storage_opt_user"`
	StorageOptTime           *time.Time `gorm:"column:storage_opt_time" json:"storage_opt_time"`
	CreateTime               *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime               *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy                 string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy                 string     `gorm:"column:update_by" json:"update_by"`
	Examine                  string     `gorm:"column:examine" json:"examine"`
	AuditMsg                 string     `gorm:"column:audit_msg" json:"audit_msg"`
}

func (StockPurchase) TableName() string {
	return "stock_purchase"
}

// 采购单明细表 stock_purchase_item
type StockPurchaseItem struct {
	ItemID           string  `gorm:"column:item_id;primaryKey" json:"item_id"`
	PurchaseID       string  `gorm:"column:purchase_id" json:"purchase_id"`
	MedicinesID      string  `gorm:"column:medicines_id" json:"medicines_id"`
	PurchaseNumber   int     `gorm:"column:purchase_number" json:"purchase_number"`
	TradePrice       float64 `gorm:"column:trade_price" json:"trade_price"`
	TradeTotalAmount float64 `gorm:"column:trade_total_amount" json:"trade_total_amount"`
	BatchNumber      string  `gorm:"column:batch_number" json:"batch_number"`
	Remark           string  `gorm:"column:remark" json:"remark"`
	MedicinesName    string  `gorm:"column:medicines_name" json:"medicines_name"`
	MedicinesType    string  `gorm:"column:medicines_type" json:"medicines_type"`
	PrescriptionType string  `gorm:"column:prescription_type" json:"prescription_type"`
	ProducterID      string  `gorm:"column:producter_id" json:"producter_id"`
	Conversion       int     `gorm:"column:conversion" json:"conversion"`
	Unit             string  `gorm:"column:unit" json:"unit"`
	Keywords         string  `gorm:"column:keywords" json:"keywords"`
}

func (StockPurchaseItem) TableName() string {
	return "stock_purchase_item"
}

// 通知公告表 sys_notice
type SysNotice struct {
	NoticeID      int        `gorm:"column:notice_id;primaryKey;autoIncrement" json:"notice_id"`
	NoticeTitle   string     `gorm:"column:notice_title" json:"notice_title"`
	NoticeType    string     `gorm:"column:notice_type" json:"notice_type"`
	NoticeContent string     `gorm:"column:notice_content" json:"notice_content"`
	Status        string     `gorm:"column:status" json:"status"`
	CreateBy      string     `gorm:"column:create_by" json:"create_by"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateBy      string     `gorm:"column:update_by" json:"update_by"`
	UpdateTime    *time.Time `gorm:"column:update_time" json:"update_time"`
	Remark        string     `gorm:"column:remark" json:"remark"`
}

func (SysNotice) TableName() string {
	return "sys_notice"
}

// 操作日志记录 sys_oper_log
type SysOperLog struct {
	OperID        int64      `gorm:"column:oper_id;primaryKey;autoIncrement" json:"oper_id"`
	Title         string     `gorm:"column:title" json:"title"`
	BusinessType  string     `gorm:"column:business_type" json:"business_type"`
	Method        string     `gorm:"column:method" json:"method"`
	RequestMethod string     `gorm:"column:request_method" json:"request_method"`
	OperatorType  int        `gorm:"column:operator_type" json:"operator_type"`
	OperName      string     `gorm:"column:oper_name" json:"oper_name"`
	OperURL       string     `gorm:"column:oper_url" json:"oper_url"`
	OperIP        string     `gorm:"column:oper_ip" json:"oper_ip"`
	OperLocation  string     `gorm:"column:oper_location" json:"oper_location"`
	OperParam     string     `gorm:"column:oper_param" json:"oper_param"`
	JSONResult    string     `gorm:"column:json_result" json:"json_result"`
	Status        string     `gorm:"column:status" json:"status"`
	ErrorMsg      string     `gorm:"column:error_msg" json:"error_msg"`
	OperTime      *time.Time `gorm:"column:oper_time" json:"oper_time"`
}

func (SysOperLog) TableName() string {
	return "sys_oper_log"
}

// 挂号项表 sys_registered_item
type SysRegisteredItem struct {
	RegItemID   int64      `gorm:"column:reg_item_id;primaryKey;autoIncrement" json:"reg_item_id"`
	RegItemName string     `gorm:"column:reg_item_name" json:"reg_item_name"`
	RegItemFee  float64    `gorm:"column:reg_item_fee" json:"reg_item_fee"`
	CreateTime  *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy    string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy    string     `gorm:"column:update_by" json:"update_by"`
	Status      string     `gorm:"column:status" json:"status"`
	DelFlag     string     `gorm:"column:del_flag" json:"del_flag"`
}

func (SysRegisteredItem) TableName() string {
	return "sys_registered_item"
}

// 角色信息表 sys_role
type SysRole struct {
	RoleID     int64      `gorm:"column:role_id;primaryKey;autoIncrement" json:"role_id"`
	RoleName   string     `gorm:"column:role_name" json:"role_name"`
	RoleCode   string     `gorm:"column:role_code" json:"role_code"`
	RoleSort   int        `gorm:"column:role_sort" json:"role_sort"`
	Remark     string     `gorm:"column:remark" json:"remark"`
	Status     string     `gorm:"column:status" json:"status"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"`
	DelFlag    string     `gorm:"column:del_flag" json:"del_flag"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

// 角色和菜单关联表 sys_role_menu
type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id"`
	MenuID int64 `gorm:"column:menu_id;primaryKey" json:"menu_id"`
}

// 用户和角色关联表 sys_role_user
type SysRoleUser struct {
	UserID int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id"`
}

// 短信发送记录表 sys_sms_log
type SysSmsLog struct {
	ID         int64      `gorm:"column:id;primaryKey" json:"id"`
	Mobile     string     `gorm:"column:mobile" json:"mobile"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	Code       string     `gorm:"column:code" json:"code"`
	Status     string     `gorm:"column:status" json:"status"`
	Type       string     `gorm:"column:type" json:"type"`
	ErrorInfo  string     `gorm:"column:error_info" json:"error_info"`
}

func (SysSmsLog) TableName() string {
	return "sys_sms_log"
}

// 用户信息表 sys_user
type SysUser struct {
	UserID         int64      `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	DeptID         int64      `gorm:"column:dept_id" json:"dept_id"`
	UserName       string     `gorm:"column:user_name" json:"user_name"`
	UserType       string     `gorm:"column:user_type" json:"user_type"`
	Sex            string     `gorm:"column:sex" json:"sex"`
	Age            int        `gorm:"column:age" json:"age"`
	Picture        string     `gorm:"column:picture" json:"picture"`
	Background     string     `gorm:"column:background" json:"background"`
	Phone          string     `gorm:"column:phone" json:"phone"`
	Email          string     `gorm:"column:email" json:"email"`
	Strong         string     `gorm:"column:strong" json:"strong"`
	Honor          string     `gorm:"column:honor" json:"honor"`
	Introduction   string     `gorm:"column:introduction" json:"introduction"`
	UserRank       string     `gorm:"column:user_rank" json:"user_rank"`
	Password       string     `gorm:"column:password" json:"password"`
	LastLoginTime  *time.Time `gorm:"column:last_login_time" json:"last_login_time"`
	LastLoginIP    string     `gorm:"column:last_login_ip" json:"last_login_ip"`
	Status         string     `gorm:"column:status" json:"status"`
	UnionID        string     `gorm:"column:union_id" json:"union_id"`
	OpenID         string     `gorm:"column:open_id" json:"open_id"`
	CreateTime     *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     *time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy       string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy       string     `gorm:"column:update_by" json:"update_by"`
	Salt           string     `gorm:"column:salt" json:"salt"`
	DelFlag        string     `gorm:"column:del_flag" json:"del_flag"`
	SchedulingFlag string     `gorm:"column:scheduling_flag" json:"scheduling_flag"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
