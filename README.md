# Go Weather CLI

OpenWeatherMap API를 사용하여 한국 도시의 날씨 정보를 조회하는 CLI 애플리케이션입니다.

## 기능 요구사항

### 도시 선택 기능
- [x] 한국 주요 도시 목록을 정의한다
- [x] 도시 선택 프롬프트를 제공한다
  - [x] 번호로 도시를 선택한다
  - [x] Enter만 입력하면 기본값을 선택한다
  - [x] 잘못된 번호 입력 시 에러를 출력한다

### API 호출 기능
- [x] OpenWeatherMap API 엔드포인트에 연결한다
- [x] HTTP GET 요청을 구현한다
- [x] HTTP 응답 상태 코드를 검증한다
- [x] JSON 응답 데이터를 파싱한다
- [x] API 호출 실패 시 상세 에러 정보를 제공한다
- [x] 날씨 데이터 구조체를 정의한다

### 날씨 정보 출력
- [x] 도시 이름을 표시한다
- [x] 날씨 상태와 설명을 표시한다
- [x] 날씨 데이터 오류를 확인한다

## 데이터 흐름
1. 환경 설정 로드 (`config.go`)
2. 사용자로부터 도시 선택 (`city_selector.go`)
3. API 호출 및 데이터 수신 (`weather_response.go`)
4. 날씨 정보 출력 (`weather_display.go`)

