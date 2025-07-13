#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
医院管理系统API测试脚本
"""

import requests
import json
import time
from datetime import datetime, timedelta

# API基础配置
BASE_URL = "http://localhost:8888"
HEADERS = {
    "Content-Type": "application/json"
}

class HospitalAPITester:
    def __init__(self):
        self.session = requests.Session()
        self.session.headers.update(HEADERS)
    
    def test_list_departments(self):
        """测试获取科室列表"""
        print("\n=== 测试获取科室列表 ===")
        try:
            response = self.session.get(f"{BASE_URL}/departments")
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_list_doctors(self, department_id=1):
        """测试获取医生列表"""
        print(f"\n=== 测试获取医生列表 (科室ID: {department_id}) ===")
        try:
            response = self.session.get(f"{BASE_URL}/departments/{department_id}/doctors")
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_list_timeslots(self, doctor_id=1):
        """测试获取时间段列表"""
        print(f"\n=== 测试获取时间段列表 (医生ID: {doctor_id}) ===")
        try:
            # 获取明天的日期
            tomorrow = (datetime.now() + timedelta(days=1)).strftime("%Y-%m-%d")
            params = {"date": tomorrow}
            response = self.session.get(f"{BASE_URL}/doctors/{doctor_id}/timeslots", params=params)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_make_appointment(self):
        """测试预约挂号"""
        print("\n=== 测试预约挂号 ===")
        try:
            data = {
                "patient_id": 1,
                "doctor_id": 1,
                "department_id": 1,
                "timeslot_id": 1
            }
            response = self.session.post(f"{BASE_URL}/appointments", json=data)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_get_appointment(self, appointment_id=1):
        """测试获取预约详情"""
        print(f"\n=== 测试获取预约详情 (预约ID: {appointment_id}) ===")
        try:
            response = self.session.get(f"{BASE_URL}/appointments/{appointment_id}")
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_cancel_appointment(self, appointment_id=1):
        """测试取消预约"""
        print(f"\n=== 测试取消预约 (预约ID: {appointment_id}) ===")
        try:
            response = self.session.post(f"{BASE_URL}/appointments/{appointment_id}/cancel")
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_create_prescription(self):
        """测试创建处方"""
        print("\n=== 测试创建处方 ===")
        try:
            data = {
                "co_id": f"PRES_{int(time.time())}",
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
            }
            response = self.session.post(f"{BASE_URL}/prescriptions", json=data)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_get_prescription_list(self):
        """测试获取处方列表"""
        print("\n=== 测试获取处方列表 ===")
        try:
            params = {
                "page": 1,
                "page_size": 10,
                "patient_id": "P001"
            }
            response = self.session.get(f"{BASE_URL}/prescriptions", params=params)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_get_medicines_list(self):
        """测试获取药品列表"""
        print("\n=== 测试获取药品列表 ===")
        try:
            params = {
                "page": 1,
                "page_size": 10,
                "medicines_name": "阿莫西林"
            }
            response = self.session.get(f"{BASE_URL}/medicines", params=params)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_create_care_history(self):
        """测试创建病例"""
        print("\n=== 测试创建病例 ===")
        try:
            data = {
                "ch_id": f"CH_{int(time.time())}",
                "user_id": 1,
                "user_name": "李医生",
                "patient_id": "P001",
                "patient_name": "张三",
                "dept_id": 1,
                "dept_name": "内科",
                "receive_type": "门诊",
                "is_contagious": "否",
                "case_date": datetime.now().strftime("%Y-%m-%d"),
                "reg_id": "REG001",
                "case_title": "感冒",
                "case_result": "上呼吸道感染",
                "doctor_tips": "多休息，多喝水",
                "remark": "患者症状轻微"
            }
            response = self.session.post(f"{BASE_URL}/care-history", json=data)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def test_get_weekly_ranking(self):
        """测试获取周排行榜"""
        print("\n=== 测试获取周排行榜 ===")
        try:
            params = {"limit": 10}
            response = self.session.get(f"{BASE_URL}/rankings/weekly", params=params)
            print(f"状态码: {response.status_code}")
            print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
            return response.status_code == 200
        except Exception as e:
            print(f"错误: {e}")
            return False
    
    def run_all_tests(self):
        """运行所有测试"""
        print("开始API测试...")
        print(f"API地址: {BASE_URL}")
        
        tests = [
            ("获取科室列表", self.test_list_departments),
            ("获取医生列表", lambda: self.test_list_doctors(1)),
            ("获取时间段列表", lambda: self.test_list_timeslots(1)),
            ("预约挂号", self.test_make_appointment),
            ("获取预约详情", lambda: self.test_get_appointment(1)),
            ("取消预约", lambda: self.test_cancel_appointment(1)),
            ("创建处方", self.test_create_prescription),
            ("获取处方列表", self.test_get_prescription_list),
            ("获取药品列表", self.test_get_medicines_list),
            ("创建病例", self.test_create_care_history),
            ("获取周排行榜", self.test_get_weekly_ranking),
        ]
        
        results = []
        for test_name, test_func in tests:
            print(f"\n{'='*50}")
            print(f"测试: {test_name}")
            print('='*50)
            
            try:
                success = test_func()
                results.append((test_name, success))
                status = "✓ 成功" if success else "✗ 失败"
                print(f"结果: {status}")
            except Exception as e:
                print(f"测试异常: {e}")
                results.append((test_name, False))
            
            time.sleep(1)  # 避免请求过快
        
        # 输出测试总结
        print(f"\n{'='*50}")
        print("测试总结")
        print('='*50)
        success_count = sum(1 for _, success in results if success)
        total_count = len(results)
        
        for test_name, success in results:
            status = "✓" if success else "✗"
            print(f"{status} {test_name}")
        
        print(f"\n总计: {success_count}/{total_count} 个测试通过")
        print(f"成功率: {success_count/total_count*100:.1f}%")

if __name__ == "__main__":
    tester = HospitalAPITester()
    tester.run_all_tests() 