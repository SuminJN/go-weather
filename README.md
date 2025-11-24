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

## 동작 방식

1. **환경 설정 로드**
    - .env 파일에서 `WEATHER_API_KEY` 읽기
    - API 키 없으면 에러 출력 후 종료
2. **도시 선택**
    - 명령줄 인자가 있으면 해당 도시 사용 (`go run . Seoul`)
    - 한글/영문 도시명을 영문으로 정규화 (서울 → Seoul)
3. **API 호출**
    - OpenWeatherMap API에 HTTP GET 요청
    - 응답 상태 코드 확인 (200 OK 여부)
4. **JSON 파싱**
    - API 응답을 `WeatherResponse` 구조체로 변환
    - 필요한 데이터 추출 (온도, 날씨, 습도 등)
5. **결과 출력**
    - 도시명, 좌표
    - 날씨 상태 및 설명
    - 현재/체감/최저/최고 온도
    - 습도
6. **에러 처리**
    - 각 단계에서 실패 시 에러 메시지 출력

## Go 학습 포인트

- 구조체
- 에러 처리 패턴
    
    ```go
    result, err := someFunction()
    if err != nil {
        return nil, fmt.Errorf("설명: %w", err)
    }
    ```
    
- HTTP 통신
    - GET 요청: `http.Get`
    - Response 처리: `defer resp.Body.Close()`
    - 상태 코드 확인: `resp.StatusCode`
- 외부 패키지 관리
    - `go.mod` 의존성 정의
- JSON 처리
- 문자열 처리

## 출력 예시

```go
--- 날씨 정보 ---
도시: Seoul (위도: 37.57, 경도: 126.98)
날씨: Clear (clear sky)
현재 온도: 15.3°C
체감 온도: 14.2°C
최저/최고 온도: 13.0°C / 17.5°C
습도: 65%
```

## 결과 화면

<img width="572" height="425" alt="image" src="https://github.com/user-attachments/assets/52c08bfe-b5a3-4c82-a0b5-93529eb65656" />
<img width="613" height="220" alt="image" src="https://github.com/user-attachments/assets/4e0ed1de-348b-4eb6-91b8-91e28f1f7cb0" />

