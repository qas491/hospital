#!/bin/bash

# 医院管理系统API测试脚本
BASE_URL="http://localhost:8888"

echo "开始API测试..."
echo "API地址: $BASE_URL"
echo "=================================="

# 测试函数
test_api() {
    local method=$1
    local endpoint=$2
    local data=$3
    local description=$4
    
    echo "测试: $description"
    echo "请求: $method $BASE_URL$endpoint"
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s -w "\n%{http_code}" "$BASE_URL$endpoint")
    else
        response=$(curl -s -w "\n%{http_code}" -H "Content-Type: application/json" -d "$data" "$BASE_URL$endpoint")
    fi
    
    # 分离响应体和状态码
    http_code=$(echo "$response" | tail -n1)
    response_body=$(echo "$response" | head -n -1)
    
    echo "状态码: $http_code"
    echo "响应: $response_body"
    
    if [ "$http_code" = "200" ]; then
        echo "✓ 成功"
    else
        echo "✗ 失败"
    fi
    echo "----------------------------------"
}

# 1. 测试获取科室列表
test_api "GET" "/departments" "" "获取科室列表"

# 2. 测试获取医生列表
test_api "GET" "/departments/1/doctors" "" "获取医生列表"

# 3. 测试获取时间段列表
test_api "GET" "/doctors/1/timeslots?date=$(date -d '+1 day' +%Y-%m-%d)" "" "获取时间段列表"

# 4. 测试预约挂号
appointment_data='{
    "patient_id": 1,
    "doctor_id": 1,
    "department_id": 1,
    "timeslot_id": 1
}'
test_api "POST" "/appointments" "$appointment_data" "预约挂号"

# 5. 测试获取预约详情
test_api "GET" "/appointments/1" "" "获取预约详情"

# 6. 测试取消预约
test_api "POST" "/appointments/1/cancel" "" "取消预约"

# 7. 测试创建处方
prescription_data='{
    "co_id": "PRES_1234567890",
    "co_type": "prescription",
    "user_id": 1,
    "patient_id": "P001",
    "patient_name": "张三",
    "ch_id": "CH001",
    "all_amount": 150.50,
    "create_by": "doctor001",
    "items": [
        {
            "item_id": "ITEM001",
            "item_ref_id": "MED001",
            "item_name": "阿莫西林胶囊",
            "item_type": "medicine",
            "num": 2.0,
            "price": 25.00,
            "amount": 50.00,
            "remark": "每日3次，每次1粒",
            "status": "active"
        }
    ]
}'
test_api "POST" "/prescriptions" "$prescription_data" "创建处方"

# 8. 测试获取处方列表
test_api "GET" "/prescriptions?page=1&page_size=10&patient_id=P001" "" "获取处方列表"

# 9. 测试获取药品列表
test_api "GET" "/medicines?page=1&page_size=10&medicines_name=阿莫西林" "" "获取药品列表"

# 10. 测试创建病例
care_history_data='{
    "ch_id": "CH_1234567890",
    "user_id": 1,
    "user_name": "李医生",
    "patient_id": "P001",
    "patient_name": "张三",
    "dept_id": 1,
    "dept_name": "内科",
    "receive_type": "门诊",
    "is_contagious": "否",
    "case_date": "'$(date +%Y-%m-%d)'",
    "reg_id": "REG001",
    "case_title": "感冒",
    "case_result": "上呼吸道感染",
    "doctor_tips": "多休息，多喝水",
    "remark": "患者症状轻微"
}'
test_api "POST" "/care-history" "$care_history_data" "创建病例"

# 11. 测试获取周排行榜
test_api "GET" "/rankings/weekly?limit=10" "" "获取周排行榜"

echo "测试完成！" 