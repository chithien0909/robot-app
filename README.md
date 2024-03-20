# robot-app

## Getting Started

1. Install Docker and Docker Compose

2. Run app with docker
    ```base
    docker-compose up -d
    ```

## Source Tree

```base
├── Dockerfile
├── Makefile
├── cmd                   # Chứa các lệnh khởi tạo, command start app: vd serve
├── dto                   # Chứa các struct để transfer dữ liệu input từ người dùng xuống handler
├── internal              # Chứa các private application và lib code
│    ├── helper           # Các helper function: vd như format date, parse data
│    ├── http             # Tầng transport: trong source này thì chỉ có RESTful
│    ├── handler          # Api handler và các business logic
│    ├── middleware       # Pre processing các request từ người dùng
│    │      └── server    # Http server với các router       
│    ├── model            # Database model, gorm model
│    └── repository       # CRUD repository
├── pkg                   # Các thư viên bên thứ 3, vd như ở đây là gorm
│    ├── config           # Config with .env file
│    └── database         # Database connection with GORM
└── validation            # Validate request from client
```
- dto:
  - dto/base.go: Chứa dto sử dụng chung, vd Response, ResponseWithPagination
  - dto/find_device.go: Dành cho handler read data, lấy danh sách device, danh sách device theo id
- internal:
  - helper:
    - query.go: Các helper function hỗ trợ cho việc query data, build sort data từ request, build các filter
  - http:
    - handler:
      - device.go: Chứa api handler và các business logic liên quan đến device model
      - handler.go: Chứa api handler và tập hợp các handler hiện có
    - middleware:
      - authorization.go: Pre processing liên qua đến việc xác thực người dùng (Api Key,...)
    - server:
      - server.go: Định nghĩa http server, và các router (endpoint, method)
  - model:
    - config.go: model chứa các config, bộ filter 
    - device.go: device database model
  - repository:
    - repository.go: Tổng hợp các repository (Device, ...)
    - device.go: CRUD với device database model
- pkg
  - config
    - config.go: Load thông tin config từ file .env, chứa các giá trị config
  - database
    - db.go: Thông tin connect database và, chạy migrate và seed data
- validation
  - validation.go: Chứa các custom validate để chứa validate input từ request
- cmd:
  - cmd.go: Khởi tạo command cho người dùng.
  - serve.go: Chứa command để start được app

## Thư viện package được sử dụng
1. cobra: Xử lí command
2. gorm: là ORM, kết nối thao tác với Postgres
3. gin: http web framework
4. viper: load config file from .env, yaml
5. validator: Giúp validate struct


## Các bước viết api mới
1. Đọc requirement và làm rõ (có thể hỏi lại, không ngầm hiểu các vấn đề)
  - Yêu cầu input, output
  - Các tính năng của api
  - Các fields (thông số), 1:1 hay 1:n
2. Setup môi trường
  - Chọn ngôn ngữ, các công nghệ
3. Thiết kế database
   - Các field cần có
   - Ràng buộc giữa các field
   - Kiểu dữ liệu
4. Thiết kế api
  - Xác định endpoint, method
  - Xử lí logic
  - Validate
  - Bảo mật api với API_KEY
5. Testing
  - Dựa theo requirement để viết testcase
6. Viết tài liệu sử dụng API: postman, hoặc swagger
7. Hướng dẫn cách chạy ứng dụng