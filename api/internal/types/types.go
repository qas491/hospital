// Code generated by goctl. DO NOT EDIT.
package types

type Appointment struct {
	Id            int32  `json:"id"`
	Patient_id    int32  `json:"patient_id"`
	Doctor_id     int32  `json:"doctor_id"`
	Department_id int32  `json:"department_id"`
	Timeslot_id   int32  `json:"timeslot_id"`
	Status        string `json:"status"`
	Created_at    string `json:"created_at"`
}

type CancelAppointmentReq struct {
	Appointment_id int32 `json:"appointment_id"`
}

type CancelAppointmentResp struct {
	Code    int64  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Success bool   `json:"success"` // 是否成功
}

type CareHistoryInfo struct {
	Ch_id         string `json:"ch_id"`         // 病例ID
	User_id       int64  `json:"user_id"`       // 医生ID
	User_name     string `json:"user_name"`     // 医生姓名
	Patient_id    string `json:"patient_id"`    // 患者ID
	Patient_name  string `json:"patient_name"`  // 患者姓名
	Dept_id       int64  `json:"dept_id"`       // 科室ID
	Dept_name     string `json:"dept_name"`     // 科室名称
	Receive_type  string `json:"receive_type"`  // 接诊类型
	Is_contagious string `json:"is_contagious"` // 是否传染病
	Care_time     string `json:"care_time"`     // 就诊时间
	Case_date     string `json:"case_date"`     // 病例日期
	Reg_id        string `json:"reg_id"`        // 挂号ID
	Case_title    string `json:"case_title"`    // 病例标题
	Case_result   string `json:"case_result"`   // 病例结果
	Doctor_tips   string `json:"doctor_tips"`   // 医生建议
	Remark        string `json:"remark"`        // 备注
}

type CreateCareHistoryReq struct {
	Ch_id         string `json:"ch_id"`         // 病例ID
	User_id       int64  `json:"user_id"`       // 医生ID
	User_name     string `json:"user_name"`     // 医生姓名
	Patient_id    string `json:"patient_id"`    // 患者ID
	Patient_name  string `json:"patient_name"`  // 患者姓名
	Dept_id       int64  `json:"dept_id"`       // 科室ID
	Dept_name     string `json:"dept_name"`     // 科室名称
	Receive_type  string `json:"receive_type"`  // 接诊类型
	Is_contagious string `json:"is_contagious"` // 是否传染病
	Case_date     string `json:"case_date"`     // 病例日期
	Reg_id        string `json:"reg_id"`        // 挂号ID
	Case_title    string `json:"case_title"`    // 病例标题
	Case_result   string `json:"case_result"`   // 病例结果
	Doctor_tips   string `json:"doctor_tips"`   // 医生建议
	Remark        string `json:"remark"`        // 备注
}

type CreateCareHistoryResp struct {
	Code    int64  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Ch_id   string `json:"ch_id"`   // 病例ID
}

type CreatePrescriptionReq struct {
	Co_id        string             `json:"co_id"`        // 处方ID
	Co_type      string             `json:"co_type"`      // 处方类型
	User_id      int64              `json:"user_id"`      // 医生ID
	Patient_id   string             `json:"patient_id"`   // 患者ID
	Patient_name string             `json:"patient_name"` // 患者姓名
	Ch_id        string             `json:"ch_id"`        // 病例ID
	All_amount   float64            `json:"all_amount"`   // 总金额
	Create_by    string             `json:"create_by"`    // 创建人
	Items        []PrescriptionItem `json:"items"`        // 处方项目列表
}

type CreatePrescriptionResp struct {
	Code    int64  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Co_id   string `json:"co_id"`   // 处方ID
}

type Department struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Doctor struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	Department_id int32  `json:"department_id"`
	Title         string `json:"title"`
	Profile       string `json:"profile"`
}

type DoctorPerformanceInfo struct {
	Doctor_id           int64               `json:"doctor_id"`           // 医生ID
	Doctor_name         string              `json:"doctor_name"`         // 医生姓名
	Dept_name           string              `json:"dept_name"`           // 科室名称
	Total_performance   float64             `json:"total_performance"`   // 总业绩
	Prescription_count  int64               `json:"prescription_count"`  // 处方数量
	Payment_count       int64               `json:"payment_count"`       // 缴费次数
	Performance_details []PerformanceDetail `json:"performance_details"` // 业绩明细
}

type GetAppointmentReq struct {
	Appointment_id int32 `json:"appointment_id"`
}

type GetAppointmentResp struct {
	Code        int64       `json:"code"`        // 状态码
	Message     string      `json:"message"`     // 提示信息
	Appointment Appointment `json:"appointment"` // 预约信息
}

type GetCareHistoryDetailReq struct {
	Ch_id string `json:"ch_id"` // 病例ID
}

type GetCareHistoryDetailResp struct {
	Code    int64           `json:"code"`    // 状态码
	Message string          `json:"message"` // 提示信息
	Detail  CareHistoryInfo `json:"detail"`  // 病例详情
}

type GetCareHistoryListReq struct {
	Page         int64  `json:"page"`         // 页码
	Page_size    int64  `json:"page_size"`    // 每页数量
	Patient_id   string `json:"patient_id"`   // 患者ID
	Patient_name string `json:"patient_name"` // 患者姓名
	User_id      int64  `json:"user_id"`      // 医生ID
	Dept_id      string `json:"dept_id"`      // 科室ID
	Case_date    string `json:"case_date"`    // 病例日期
	Start_time   string `json:"start_time"`   // 开始时间
	End_time     string `json:"end_time"`     // 结束时间
}

type GetCareHistoryListResp struct {
	Code    int64             `json:"code"`    // 状态码
	Message string            `json:"message"` // 提示信息
	List    []CareHistoryInfo `json:"list"`    // 病例列表
	Total   int64             `json:"total"`   // 总数
}

type GetDoctorPerformanceReq struct {
	Doctor_id  int64  `json:"doctor_id"`  // 医生ID
	Start_date string `json:"start_date"` // 开始日期
	End_date   string `json:"end_date"`   // 结束日期
}

type GetDoctorPerformanceResp struct {
	Code        int64                 `json:"code"`        // 状态码
	Message     string                `json:"message"`     // 提示信息
	Performance DoctorPerformanceInfo `json:"performance"` // 业绩信息
}

type GetMedicinesDetailReq struct {
	Medicines_id uint64 `json:"medicines_id"` // 药品ID
}

type GetMedicinesDetailResp struct {
	Code    int64         `json:"code"`    // 状态码
	Message string        `json:"message"` // 提示信息
	Detail  MedicinesInfo `json:"detail"`  // 药品详情
}

type GetMedicinesListReq struct {
	Page              int64  `json:"page"`              // 页码
	Page_size         int64  `json:"page_size"`         // 每页数量
	Medicines_name    string `json:"medicines_name"`    // 药品名称
	Medicines_type    string `json:"medicines_type"`    // 药品类型
	Prescription_type string `json:"prescription_type"` // 处方类型
	Status            string `json:"status"`            // 状态
	Keywords          string `json:"keywords"`          // 关键词
}

type GetMedicinesListResp struct {
	Code    int64           `json:"code"`    // 状态码
	Message string          `json:"message"` // 提示信息
	List    []MedicinesInfo `json:"list"`    // 药品列表
	Total   int64           `json:"total"`   // 总数
}

type GetPrescriptionDetailReq struct {
	Co_id string `json:"co_id"` // 处方ID
}

type GetPrescriptionDetailResp struct {
	Code    int64            `json:"code"`    // 状态码
	Message string           `json:"message"` // 提示信息
	Detail  PrescriptionInfo `json:"detail"`  // 处方详情
	Qrcode  string           `json:"qrcode"`  // 处方二维码
}

type GetPrescriptionListReq struct {
	Page         int64  `json:"page"`         // 页码
	Page_size    int64  `json:"page_size"`    // 每页数量
	Patient_id   string `json:"patient_id"`   // 患者ID
	Patient_name string `json:"patient_name"` // 患者姓名
	User_id      int64  `json:"user_id"`      // 医生ID
	Co_type      string `json:"co_type"`      // 处方类型
	Status       string `json:"status"`       // 状态
	Start_time   string `json:"start_time"`   // 开始时间
	End_time     string `json:"end_time"`     // 结束时间
}

type GetPrescriptionListResp struct {
	Code    int64              `json:"code"`    // 状态码
	Message string             `json:"message"` // 提示信息
	List    []PrescriptionInfo `json:"list"`    // 处方列表
	Total   int64              `json:"total"`   // 总数
}

type GetWeeklyRankingReq struct {
	Limit int64 `json:"limit"` // 获取前N名，默认10
}

type GetWeeklyRankingResp struct {
	Code       int64         `json:"code"`       // 状态码
	Message    string        `json:"message"`    // 提示信息
	Rankings   []RankingInfo `json:"rankings"`   // 排行榜列表
	Week_start string        `json:"week_start"` // 周开始日期
	Week_end   string        `json:"week_end"`   // 周结束日期
}

type ListDepartmentsReq struct {
}

type ListDepartmentsResp struct {
	Code        int64        `json:"code"`        // 状态码
	Message     string       `json:"message"`     // 提示信息
	Departments []Department `json:"departments"` // 科室列表
}

type ListDoctorsReq struct {
	Department_id int32 `json:"department_id"`
}

type ListDoctorsResp struct {
	Code    int64    `json:"code"`    // 状态码
	Message string   `json:"message"` // 提示信息
	Doctors []Doctor `json:"doctors"` // 医生列表
}

type ListTimeSlotsReq struct {
	Date string `json:"date"` // 查询日期
}

type ListTimeSlotsResp struct {
	Code      int64      `json:"code"`      // 状态码
	Message   string     `json:"message"`   // 提示信息
	Timeslots []TimeSlot `json:"timeslots"` // 时间段列表
}

type MakeAppointmentReq struct {
	Patient_id    int32 `json:"patient_id"`
	Doctor_id     int32 `json:"doctor_id"`
	Department_id int32 `json:"department_id"`
	Timeslot_id   int32 `json:"timeslot_id"`
}

type MakeAppointmentResp struct {
	Code        int64       `json:"code"`        // 状态码
	Message     string      `json:"message"`     // 提示信息
	Appointment Appointment `json:"appointment"` // 预约信息
}

type MedicinesInfo struct {
	Medicines_id               uint64  `json:"medicines_id"`               // 药品ID
	Medicines_number           string  `json:"medicines_number"`           // 药品编号
	Medicines_name             string  `json:"medicines_name"`             // 药品名称
	Medicines_type             string  `json:"medicines_type"`             // 药品类型
	Prescription_type          string  `json:"prescription_type"`          // 处方类型
	Prescription_price         float64 `json:"prescription_price"`         // 处方价格
	Unit                       string  `json:"unit"`                       // 单位
	Conversion                 int32   `json:"conversion"`                 // 换算
	Keywords                   string  `json:"keywords"`                   // 关键词
	Producter_id               string  `json:"producter_id"`               // 生产厂家ID
	Status                     string  `json:"status"`                     // 状态
	Medicines_stock_num        float64 `json:"medicines_stock_num"`        // 库存数量
	Medicines_stock_danger_num float64 `json:"medicines_stock_danger_num"` // 危险库存数量
	Create_time                string  `json:"create_time"`                // 创建时间
	Update_time                string  `json:"update_time"`                // 更新时间
	Create_by                  string  `json:"create_by"`                  // 创建人
	Update_by                  string  `json:"update_by"`                  // 更新人
	Del_flag                   string  `json:"del_flag"`                   // 删除标志
}

type Patient struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Id_card string `json:"id_card"`
}

type PerformanceDetail struct {
	Performance_id     string  `json:"performance_id"`     // 业绩ID
	Co_id              string  `json:"co_id"`              // 处方ID
	Payment_amount     float64 `json:"payment_amount"`     // 缴费金额
	Performance_amount float64 `json:"performance_amount"` // 业绩金额
	Performance_rate   float64 `json:"performance_rate"`   // 业绩比例
	Performance_date   string  `json:"performance_date"`   // 业绩日期
}

type PrescriptionInfo struct {
	Co_id        string             `json:"co_id"`        // 处方ID
	Co_type      string             `json:"co_type"`      // 处方类型
	User_id      int64              `json:"user_id"`      // 医生ID
	Patient_id   string             `json:"patient_id"`   // 患者ID
	Patient_name string             `json:"patient_name"` // 患者姓名
	Ch_id        string             `json:"ch_id"`        // 病例ID
	All_amount   float64            `json:"all_amount"`   // 总金额
	Create_by    string             `json:"create_by"`    // 创建人
	Create_time  string             `json:"create_time"`  // 创建时间
	Update_by    string             `json:"update_by"`    // 更新人
	Update_time  string             `json:"update_time"`  // 更新时间
	Items        []PrescriptionItem `json:"items"`        // 处方项目列表
}

type PrescriptionItem struct {
	Item_id     string  `json:"item_id"`     // 项目ID
	Item_ref_id string  `json:"item_ref_id"` // 引用ID
	Item_name   string  `json:"item_name"`   // 项目名称
	Item_type   string  `json:"item_type"`   // 项目类型
	Num         float64 `json:"num"`         // 数量
	Price       float64 `json:"price"`       // 单价
	Amount      float64 `json:"amount"`      // 金额
	Remark      string  `json:"remark"`      // 备注
	Status      string  `json:"status"`      // 状态
}

type RankingInfo struct {
	Rank               int64   `json:"rank"`               // 排名
	Doctor_id          int64   `json:"doctor_id"`          // 医生ID
	Doctor_name        string  `json:"doctor_name"`        // 医生姓名
	Dept_name          string  `json:"dept_name"`          // 科室名称
	Total_performance  float64 `json:"total_performance"`  // 总业绩
	Prescription_count int64   `json:"prescription_count"` // 处方数量
}

type ReviewPrescriptionReq struct {
	Co_id         string `json:"co_id"`         // 处方ID
	Review_status string `json:"review_status"` // 审核状态
	Review_by     string `json:"review_by"`     // 审核人
	Review_remark string `json:"review_remark"` // 审核备注
}

type ReviewPrescriptionResp struct {
	Code    int64  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Success bool   `json:"success"` // 是否成功
}

type SelectMedicinesReq struct {
	Medicines_id      string  `json:"medicines_id"`      // 药品ID
	Medicines_name    string  `json:"medicines_name"`    // 药品名称
	Medicines_type    string  `json:"medicines_type"`    // 药品类型
	Prescription_type string  `json:"prescription_type"` // 处方类型
	Num               float64 `json:"num"`               // 数量
	Unit              string  `json:"unit"`              // 单位
}

type SelectMedicinesResp struct {
	Code      int64         `json:"code"`      // 状态码
	Message   string        `json:"message"`   // 提示信息
	Medicines MedicinesInfo `json:"medicines"` // 药品信息
}

type TimeSlot struct {
	Id         int32  `json:"id"`
	Doctor_id  int32  `json:"doctor_id"`
	Date       string `json:"date"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	Available  int32  `json:"available"`
}
