# Todo CLI 애플리케이션

Go 언어로 작성된 커맨드라인 Todo 애플리케이션입니다. 효율적인 작업 관리를 도와줍니다.

## 프로젝트 구조

```
todo-go-cli/
├── cmd/
│   └── todo/
│       └── main.go           # 애플리케이션 진입점
├── internal/
│   ├── config/
│   │   ├── config.go         # 설정 관리
│   │   ├── config_test.go    # 설정 테스트
│   │   └── config_suite_test.go
│   ├── domain/
│   │   ├── task.go          # 작업 도메인 모델
│   │   └── task_test.go     # 작업 도메인 테스트
│   ├── repository/
│   │   └── task_repository.go # 데이터 영속성 계층
│   └── service/
│       └── task_service.go   # 비즈니스 로직 계층
├── data/
│   └── tasks.json           # 작업 저장 파일
├── go.mod                   # Go 모듈 파일
├── go.sum                   # Go 의존성 체크섬
└── README.md               # 프로젝트 문서
```

## 컴포넌트

### 1. 도메인 계층 (`internal/domain/`)

- `task.go`: 핵심 Task 엔티티 정의
  - `Task` 구조체: ID, 설명, 완료 상태를 가진 할일 항목을 표현
  - `NewTask()`: 주어진 설명으로 새로운 작업 생성
  - `MarkAsDone()`: 작업을 완료로 표시
  - `String()`: 작업의 문자열 표현을 반환

### 2. 저장소 계층 (`internal/repository/`)

- `task_repository.go`: 데이터 영속성 처리
  - `TaskRepository` 인터페이스: 데이터 접근 메서드 정의
  - `FileTaskRepository` 구조체: 파일 기반 저장소 구현
  - 메서드:
    - `NewFileTaskRepository()`: 새로운 저장소 인스턴스 생성
    - `GetTasks()`: 모든 작업 조회
    - `SaveTasks()`: 작업을 파일에 저장
    - `LoadTasks()`: 파일에서 작업 로드
    - `AddTask()`: 새로운 작업 추가
    - `FindTaskByID()`: ID로 작업 찾기

### 3. 서비스 계층 (`internal/service/`)

- `task_service.go`: 비즈니스 로직 구현
  - `TaskService` 인터페이스: 비즈니스 작업 정의
  - `DefaultTaskService` 구조체: 작업 관련 기능 구현
  - 메서드:
    - `NewTaskService()`: 새로운 서비스 인스턴스 생성
    - `AddTask()`: 새로운 작업 생성 및 저장
    - `ListTasks()`: 모든 작업 나열
    - `MarkTaskAsDone()`: 작업을 완료로 표시

### 4. 설정 (`internal/config/`)

- `config.go`: 애플리케이션 설정 관리
  - `Config` 구조체: 애플리케이션 설정 보관
  - 메서드:
    - `NewConfig()`: 새로운 설정 생성
    - `GetProjectRootPath()`: 프로젝트 루트 디렉토리 결정

## 설치 방법

1. 저장소 복제:

   ```bash
   git clone <repository-url>
   cd todo-go-cli
   ```

2. Go 모듈 초기화 (이미 `go.mod`가 있다면 생략):

   ```bash
   go mod init todo-go-cli
   ```

3. 의존성 설치:

   ```bash
   go mod download
   go mod tidy  # 필요한 의존성을 추가하고 사용하지 않는 의존성을 제거
   ```

## 테스트 실행

1. 모든 테스트 실행:

   ```bash
   go test ./... -v
   ```

2. 특정 패키지 테스트 실행:

   ```bash
   # 도메인 테스트 실행
   go test ./internal/domain -v

   # 설정 테스트 실행
   go test ./internal/config -v
   ```

3. 커버리지 테스트 실행:
   ```bash
   go test ./... -coverprofile=coverage.out
   go tool cover -html=coverage.out  # 브라우저에서 커버리지 보기
   ```

## 사용 방법

애플리케이션 빌드 및 실행:

```bash
# 빌드
go build -o bin/todo cmd/todo/main.go

# 직접 실행
go run cmd/todo/main.go [명령어] [인자]
```

사용 가능한 명령어:

1. 새로운 작업 추가:

   ```bash
   go run cmd/todo/main.go add "장보기"
   ```

2. 모든 작업 나열:

   ```bash
   go run cmd/todo/main.go list
   ```

3. 작업 완료 표시:
   ```bash
   go run cmd/todo/main.go done 1  # ID가 1인 작업을 완료로 표시
   ```

## 개발

이 애플리케이션은 클린 아키텍처 원칙을 따릅니다:

1. **도메인 계층**: 비즈니스 엔티티와 핵심 비즈니스 규칙 포함
2. **저장소 계층**: 데이터 영속성 처리
3. **서비스 계층**: 비즈니스 로직 구현
4. **설정**: 애플리케이션 설정 관리

### 새로운 기능 추가

1. `internal/domain/`에 도메인 로직 추가
2. `internal/repository/`에 영속성 구현
3. `internal/service/`에 비즈니스 로직 추가
4. `cmd/todo/main.go`에 CLI 명령어 업데이트

### 테스팅

- Ginkgo와 Gomega를 사용한 BDD 스타일 테스트
- 각 패키지별 테스트 스위트
- 개발 중 자주 테스트 실행

## 의존성

- github.com/onsi/ginkgo/v2: BDD 테스트 프레임워크
- github.com/onsi/gomega: 매처/어설션 라이브러리

## 기여하기

1. 저장소 포크
2. 기능 브랜치 생성
3. 변경사항 커밋
4. 브랜치에 푸시
5. Pull Request 생성

## 라이선스

이 프로젝트는 MIT 라이선스를 따릅니다 - 자세한 내용은 LICENSE 파일을 참조하세요.
