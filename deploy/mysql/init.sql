-- 医院管理系统数据库初始化脚本

-- 创建数据库
CREATE DATABASE IF NOT EXISTS hospital DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE hospital;

-- 科室表
CREATE TABLE IF NOT EXISTS departments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '科室名称',
    description TEXT COMMENT '科室描述',
    status TINYINT DEFAULT 1 COMMENT '状态：1-启用，0-禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='科室信息表';

-- 医生表
CREATE TABLE IF NOT EXISTS doctors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '医生姓名',
    department_id INT NOT NULL COMMENT '所属科室ID',
    title VARCHAR(50) COMMENT '职称',
    profile TEXT COMMENT '医生简介',
    status TINYINT DEFAULT 1 COMMENT '状态：1-启用，0-禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES departments(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='医生信息表';

-- 患者表
CREATE TABLE IF NOT EXISTS patients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '患者姓名',
    gender VARCHAR(10) COMMENT '性别',
    phone VARCHAR(20) COMMENT '联系电话',
    id_card VARCHAR(18) COMMENT '身份证号',
    status TINYINT DEFAULT 1 COMMENT '状态：1-启用，0-禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='患者信息表';

-- 时间段表
CREATE TABLE IF NOT EXISTS time_slots (
    id INT PRIMARY KEY AUTO_INCREMENT,
    doctor_id INT NOT NULL COMMENT '医生ID',
    date DATE NOT NULL COMMENT '日期',
    start_time TIME NOT NULL COMMENT '开始时间',
    end_time TIME NOT NULL COMMENT '结束时间',
    available INT DEFAULT 10 COMMENT '可预约数量',
    status TINYINT DEFAULT 1 COMMENT '状态：1-可用，0-不可用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (doctor_id) REFERENCES doctors(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='时间段表';

-- 预约表
CREATE TABLE IF NOT EXISTS appointments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    patient_id INT NOT NULL COMMENT '患者ID',
    doctor_id INT NOT NULL COMMENT '医生ID',
    department_id INT NOT NULL COMMENT '科室ID',
    timeslot_id INT NOT NULL COMMENT '时间段ID',
    status VARCHAR(20) DEFAULT 'pending' COMMENT '预约状态：pending-待确认，confirmed-已确认，cancelled-已取消',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (doctor_id) REFERENCES doctors(id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (timeslot_id) REFERENCES time_slots(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预约表';

-- 药品表
CREATE TABLE IF NOT EXISTS medicines (
    medicines_id INT PRIMARY KEY AUTO_INCREMENT,
    medicines_number VARCHAR(50) UNIQUE NOT NULL COMMENT '药品编号',
    medicines_name VARCHAR(200) NOT NULL COMMENT '药品名称',
    medicines_type VARCHAR(100) COMMENT '药品类型',
    prescription_type VARCHAR(50) COMMENT '处方类型',
    prescription_price DECIMAL(10,2) DEFAULT 0.00 COMMENT '处方价格',
    unit VARCHAR(20) COMMENT '单位',
    conversion INT DEFAULT 1 COMMENT '换算',
    keywords VARCHAR(500) COMMENT '关键词',
    producter_id VARCHAR(50) COMMENT '生产厂家ID',
    status VARCHAR(20) DEFAULT 'active' COMMENT '状态',
    medicines_stock_num DECIMAL(10,2) DEFAULT 0 COMMENT '库存数量',
    medicines_stock_danger_num DECIMAL(10,2) DEFAULT 0 COMMENT '危险库存数量',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_by VARCHAR(50) COMMENT '创建人',
    update_by VARCHAR(50) COMMENT '更新人',
    del_flag VARCHAR(1) DEFAULT '0' COMMENT '删除标志'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='药品表';

-- 处方表
CREATE TABLE IF NOT EXISTS prescriptions (
    co_id VARCHAR(50) PRIMARY KEY COMMENT '处方ID',
    co_type VARCHAR(20) NOT NULL COMMENT '处方类型',
    user_id INT NOT NULL COMMENT '医生ID',
    patient_id VARCHAR(50) NOT NULL COMMENT '患者ID',
    patient_name VARCHAR(100) NOT NULL COMMENT '患者姓名',
    ch_id VARCHAR(50) COMMENT '病例ID',
    all_amount DECIMAL(10,2) DEFAULT 0.00 COMMENT '总金额',
    create_by VARCHAR(50) COMMENT '创建人',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(50) COMMENT '更新人',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'active' COMMENT '状态'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方表';

-- 处方项目表
CREATE TABLE IF NOT EXISTS prescription_items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    co_id VARCHAR(50) NOT NULL COMMENT '处方ID',
    item_id VARCHAR(50) NOT NULL COMMENT '项目ID',
    item_ref_id VARCHAR(50) COMMENT '引用ID',
    item_name VARCHAR(200) NOT NULL COMMENT '项目名称',
    item_type VARCHAR(50) COMMENT '项目类型',
    num DECIMAL(10,2) DEFAULT 1 COMMENT '数量',
    price DECIMAL(10,2) DEFAULT 0.00 COMMENT '单价',
    amount DECIMAL(10,2) DEFAULT 0.00 COMMENT '金额',
    remark TEXT COMMENT '备注',
    status VARCHAR(20) DEFAULT 'active' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (co_id) REFERENCES prescriptions(co_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方项目表';

-- 病例表
CREATE TABLE IF NOT EXISTS care_history (
    ch_id VARCHAR(50) PRIMARY KEY COMMENT '病例ID',
    user_id INT NOT NULL COMMENT '医生ID',
    user_name VARCHAR(100) NOT NULL COMMENT '医生姓名',
    patient_id VARCHAR(50) NOT NULL COMMENT '患者ID',
    patient_name VARCHAR(100) NOT NULL COMMENT '患者姓名',
    dept_id INT NOT NULL COMMENT '科室ID',
    dept_name VARCHAR(100) NOT NULL COMMENT '科室名称',
    receive_type VARCHAR(50) COMMENT '接诊类型',
    is_contagious VARCHAR(10) COMMENT '是否传染病',
    care_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '就诊时间',
    case_date DATE COMMENT '病例日期',
    reg_id VARCHAR(50) COMMENT '挂号ID',
    case_title VARCHAR(200) COMMENT '病例标题',
    case_result TEXT COMMENT '病例结果',
    doctor_tips TEXT COMMENT '医生建议',
    remark TEXT COMMENT '备注',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='病例表';

-- 插入初始数据
INSERT INTO departments (name, description) VALUES 
('内科', '内科疾病诊断与治疗'),
('外科', '外科手术与治疗'),
('儿科', '儿童疾病诊断与治疗'),
('妇产科', '妇科疾病与产科服务'),
('眼科', '眼部疾病诊断与治疗'),
('口腔科', '口腔疾病诊断与治疗');

INSERT INTO doctors (name, department_id, title, profile) VALUES 
('张医生', 1, '主任医师', '从事内科临床工作20年，擅长心血管疾病诊治'),
('李医生', 1, '副主任医师', '从事内科临床工作15年，擅长呼吸系统疾病诊治'),
('王医生', 2, '主任医师', '从事外科临床工作25年，擅长普外科手术'),
('赵医生', 3, '副主任医师', '从事儿科临床工作18年，擅长儿童常见病诊治');

INSERT INTO patients (name, gender, phone, id_card) VALUES 
('张三', '男', '13800138001', '110101199001011234'),
('李四', '女', '13800138002', '110101199002022345'),
('王五', '男', '13800138003', '110101199003033456');

INSERT INTO medicines (medicines_number, medicines_name, medicines_type, prescription_type, prescription_price, unit) VALUES 
('MED001', '阿莫西林胶囊', '抗生素', '处方药', 25.00, '盒'),
('MED002', '布洛芬片', '解热镇痛', '处方药', 15.50, '盒'),
('MED003', '感冒灵颗粒', '感冒药', '非处方药', 12.00, '盒'),
('MED004', '维生素C片', '维生素', '非处方药', 8.50, '瓶');

-- 创建索引
CREATE INDEX idx_doctors_department ON doctors(department_id);
CREATE INDEX idx_appointments_patient ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor ON appointments(doctor_id);
CREATE INDEX idx_time_slots_doctor_date ON time_slots(doctor_id, date);
CREATE INDEX idx_prescriptions_patient ON prescriptions(patient_id);
CREATE INDEX idx_care_history_patient ON care_history(patient_id);
CREATE INDEX idx_medicines_name ON medicines(medicines_name); 